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
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/spacemonkeygo/openssl"
)

type Sha256 struct {
	*CryptoAlgorithm
	Args struct {
		Algorithm string `positional-arg-name:"algorithm" required:"true" description:"Pick either crypto or openssl"`
	} `positional-args:"true"`
}

func (s *Sha256) Name() string {
	return "Sha256"
}

func (s *Sha256) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(s.Name(), "Performs benchmark for Sha256", "Computes Sha256 of random data and reports the result along with benchmark", s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sha256) Compute(data []byte) ([]byte, error) {
	if s.Args.Algorithm == CRYPTO_ALGORITHM {
		result_bytes := sha256.Sum256(data)
		return result_bytes[:], nil
	} else if s.Args.Algorithm == OPENSSL_ALGORITHM {
		result_bytes, err := openssl.SHA256(data)
		return result_bytes[:], err
	} else {
		return []byte{}, errors.New(fmt.Sprintf("Unknown algorithm: %s", s.Args.Algorithm))
	}
}

func (s *Sha256) Run(child interface{}) error {
	return s.CryptoAlgorithm.Run(s)
}
