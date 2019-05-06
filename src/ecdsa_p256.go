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
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/elliptic"
	"errors"
	"crypto"
	"fmt"
	"github.com/arsulegai/openssl"
	"github.com/jessevdk/go-flags"
)

type Ecdsa_P256 struct {
	*CryptoAlgorithm
	Args struct {
		Algorithm string `positional-arg-name:"algorithm" required:"true" description:"Pick either crypto or openssl"`
	} `positional-args:"true"`
}

func (s *Ecdsa_P256) Name() string {
	return "Ecdsa_P256"
}

func (s *Ecdsa_P256) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(s.Name(), "Performs benchmark for Ecdsa_P256", "Signs and Verifies using Ecdsa_P256 of random data and reports the result along with benchmark", s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Ecdsa_P256) Compute(data []byte) ([]byte, error) {
	if s.Args.Algorithm == CRYPTO_ALGORITHM {
		result_bytes, err := compute_crypto_ecdsa_p256(data)
		return result_bytes[:], err
	} else if s.Args.Algorithm == OPENSSL_ALGORITHM {
		result_bytes, err := compute_openssl_ecdsa_p256(data)
		return result_bytes[:], err
	} else {
		return []byte{}, errors.New(fmt.Sprintf("Unknown algorithm: %s", s.Args.Algorithm))
	}
}

func (s *Ecdsa_P256) Run(child interface{}) error {
	return s.CryptoAlgorithm.Run(s)
}

func compute_crypto_ecdsa_p256(data []byte) ([]byte, error) {
	digest_bytes := sha256.Sum256(data)
	digest := digest_bytes[:]
	privatekey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return []byte{}, err
	}
	signature, err := privatekey.Sign(rand.Reader, digest, crypto.SHA256)
	if err != nil {
		return []byte{}, err
	}
	pubkey := privatekey.PublicKey
	r, s, err := ecdsa.Sign(rand.Reader, privatekey, digest)
	if err != nil {
		return []byte{}, err
	}
	if ecdsa.Verify(&pubkey, digest, r, s) {
		return signature[:], nil
	} else {
		return []byte{}, errors.New("Signature verification failed!")
	}
}

func compute_openssl_ecdsa_p256(data []byte) ([]byte, error) {
	digest_bytes, err := openssl.SHA256(data)
	if err != nil {
		return []byte{}, err
	}
	digest := digest_bytes[:]
	privatekey, err := openssl.GenerateECKey(openssl.Prime256v1)
	if err != nil {
		return []byte{}, err
	}
	// pubkey := privatekey.Public()
	signature, err := privatekey.SignPKCS1v15(openssl.SHA256_Method, digest)
	if err != nil {
		return []byte{}, err
	}
	err = privatekey.VerifyPKCS1v15(openssl.SHA256_Method, digest, signature)
	if err != nil {
		return []byte{}, err
	}
	return signature[:], nil
}
