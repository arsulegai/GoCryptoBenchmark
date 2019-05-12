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
	"errors"
	"fmt"
	"github.com/arsulegai/openssl"
	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/sha3"
)

type Sha3_384 struct {
	*CryptoAlgorithm
	Args struct {
		Algorithm string `positional-arg-name:"algorithm" required:"true" description:"Pick either crypto or openssl"`
	} `positional-args:"true"`
}

func (s *Sha3_384) Name() string {
	return "Sha3_384"
}

func (s *Sha3_384) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(s.Name(), "Performs benchmark for Sha3_384", "Computes Sha3_384 of random data and reports the result along with benchmark", s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sha3_384) Compute(data []byte) ([]byte, error) {
	if s.Args.Algorithm == CRYPTO_ALGORITHM {
		result_bytes := sha3.Sum384(data)
		return result_bytes[:], nil
	} else if s.Args.Algorithm == OPENSSL_ALGORITHM {
		result_bytes, err := openssl.SHA3_384(data)
		return result_bytes[:], err
	} else {
		return []byte{}, errors.New(fmt.Sprintf("Unknown algorithm: %s", s.Args.Algorithm))
	}
}

func (s *Sha3_384) Run(child interface{}) error {
	return s.CryptoAlgorithm.Run(s)
}
