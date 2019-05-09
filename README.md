# Go Crypto Benchmark
Go Crypto Benchmark is a tool developed to perform benchmark activities on
different popular algorithms. Tool supports Go Crypto and OpenSSL wrapper
libraries.

## Introduction
The tool will run popular cryto algorithms on the machine, reporting the
benchmark numbers on the hardware.

**Note:** The tool uses forked version of [OpenSSL wrapper](https://github.com/spacemonkeygo/openssl)
which can be found [here](https://github.com/arsulegai/openssl).
All credits to the original authors.

## Prerequisites
The project has dependency on following libraries
* OpenSSL version 1.1.1, please refer to docker/Dockerfile for information
on installing it on Ubuntu if there's no debian distribution available.

## Build

### Docker build dependencies
The build has been tested on these versions and same is recommended.
* Docker version 18.06
* Docker compose version 1.22.0

### Procedure
Run the following to generate `bin/gocryptobenchmark`
(no Go compiler required):
```
docker-compose -f docker/compose/docker-compose-build.yaml up
```
**Note:**
1. Generated binary is targeted for Ubuntu 18.04. Binary would be
stored in `bin` directory.
2. If the build is successful, docker-compose up should exit with
status code 0.

## Run
To run the benchmark using the OpenSSL and Go crypto, run the following:
```
./bin/gocryptobenchmark <CRYPTO ALGORITHM TO BENCHMARK> <LIBRARY>

where
  <CRYPTO ALGORITHM TO BENCHMARK> is one of
    Sha256
    Sha384
    Sha512
    Sha3_256
    Sha3_384
    Ecdsa_P256
  <LIBRARY> is one of
    openssl
    crypto
```
For example, to benchmark Ecdsa_P256, the command is
```
./bin/gocryptobenchmark Ecdsa_P256 openssl
./bin/gocryptobenchmark Ecdsa_P256 crypto
```
**Note:** For ECDSA P256, the data is first hashed using SHA 256.

## Contributing
This software is in development phase and is Apache 2.0 licensed. We accept
contributions via [GitHub](https://github.com/arsulegai/GoCryptoBenchmark) pull
requests.
Each commit must include a `Signed-off-by:` in the commit message
(`git commit -s`). This sign-off means you agree the commit satisfies the
[Developer Certificate of Origin (DCO)](https://developercertificate.org/).

## License
This software is licensed under the [Apache License Version 2.0](LICENSE)
software license.

&copy; Copyright 2019, Intel Corporation.
