/**
 * Copyright 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ------------------------------------------------------------------------------
 */

package main

import (
	"bufio"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

// All subcommands implement this interface
type Command interface {
	Register(*flags.Command) error
	Name() string
	Run(interface{}) error
	Compute([]byte) ([]byte, error)
}

type Opts struct {
	Version    bool   `short:"V" long:"version" description:"Display version information"`
	CPUProfile string `short:"p" long:"profile" description:"Generate pprof CPU profile"`
}

type CryptoAlgorithm struct{}

var DISTRIBUTION_NAME string
var DISTRIBUTION_VERSION string

func (c *CryptoAlgorithm) Register(*flags.Command) error { return nil }

func (c *CryptoAlgorithm) Name() string { return "" }

func (c *CryptoAlgorithm) Compute(data []byte) ([]byte, error) { return []byte{}, nil }

func (c *CryptoAlgorithm) Run(child interface{}) error {
	pid := os.Getpid()
	fmt.Println("Start performance measuring tool against the process id: ", strconv.Itoa(pid))
	fmt.Println("Then press [ENTER] key to continue!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Create data of uint8 to perform crypto algorithm
	data := make([]byte, NUMBER_OF_INPUT_BYTES)
	start := time.Now()
	var result []byte
	var err error
	for i := 0; i < LOOP_TIMES; i++ {
		switch child.(type) {
		case *Sha256:
			result, err = child.(*Sha256).Compute(data)
		case *Sha512:
			result, err = child.(*Sha512).Compute(data)
		case *Sha384:
			result, err = child.(*Sha384).Compute(data)
		case *Sha3_256:
			result, err = child.(*Sha3_256).Compute(data)
		case *Sha3_384:
			result, err = child.(*Sha3_384).Compute(data)
		case *Ecdsa_P256:
			// performs both sign() and verify() on the data
			result, err = child.(*Ecdsa_P256).Compute(data)
		default:
			break
		}
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Total time taken for crypto operation: ", elapsed)
	fmt.Println("Result: ", result)

	return err
}

func init() {
	if len(DISTRIBUTION_VERSION) == 0 {
		DISTRIBUTION_VERSION = "Unknown"
	}
}

func main() {
	return_status := 0
	defer func() {
		os.Exit(return_status)
	}()
	arguments := os.Args[1:]
	for _, arg := range arguments {
		if arg == "-V" || arg == "--version" {
			fmt.Println(DISTRIBUTION_NAME + " version " + DISTRIBUTION_VERSION)
			return_status = 0
			return
		}
	}

	var opts Opts
	parser := flags.NewParser(&opts, flags.Default)
	parser.Command.Name = "go-crypto-bmark"

	commands := []Command{
		&Sha256{},
		&Sha512{},
		&Sha384{},
		&Sha3_256{},
		&Sha3_384{},
		&Ecdsa_P256{},
	}

	for _, cmd := range commands {
		err := cmd.Register(parser.Command)
		if err != nil {
			fmt.Errorf("Couldn't register command %v: %v", cmd.Name(), err)
			return_status = 1
			return
		}
	}

	remaining, err := parser.Parse()
	if e, ok := err.(*flags.Error); ok {
		if e.Type != flags.ErrHelp {
			return_status = 1
		}
		return
	}

	if len(remaining) > 0 {
		fmt.Println("Error: Unrecognized arguments passed: ", remaining)
		return_status = 2
		return
	}

	if parser.Command.Active == nil {
		return_status = 2
		return
	}

	cpuprofile := opts.CPUProfile
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			fmt.Errorf("Error creating profile")
			return_status = 1
			return
		}
		defer f.Close()
		fmt.Println("CPU Profile using pprof enabled")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	name := parser.Command.Active.Name
	for _, cmd := range commands {
		if cmd.Name() == name {
			err := cmd.Run("")
			if err != nil {
				fmt.Println("Error: ", err)
				return_status = 1
			}
			return
		}
	}

	fmt.Println("Error: Command not found: ", name)
}
