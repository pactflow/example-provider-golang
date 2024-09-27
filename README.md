# Example Provider

![Build](https://github.com/pactflow/example-provider-golang/workflows/Build/badge.svg)

[![Can I deploy Status](https://test.pactflow.io/pacticipants/pactflow-example-provider-golang/branches/master/latest-version/can-i-deploy/to-environment/production/badge.svg)](https://test.pactflow.io/overview/provider/pactflow-example-consumer-golang/consumer/pactflow-example-consumer-golang)

This is an example of a (Gin-based) Golang provider that uses Pact, [PactFlow](https://pactflow.io) and GitHub Actions to ensure that it is compatible with the expectations its consumers have of it.

The project uses a Makefile to simulate a very simple build pipeline with two stages - test and deploy.

It is using a public tenant on PactFlow, which you can access [here](https://test.pactflow.io) using the credentials `dXfltyFMgNOFZAxr8io9wJ37iUpY42M`/`O5AIZWxelWbLvqMd8PkAVycBJh2Psyg1`. The latest version of the Example Consumer/Example Provider pact is published [here](https://test.pactflow.io/pacts/provider/pactflow-example-provider-golang/consumer/pactflow-example-consumer/latest).

## Pre-Requisites

In order to use pact-go v2, you must install pact ffi system libraries to your machine.

```
make install
```

In order to use the pact cli tools, which interact with a pact broker, the Ruby standalone binaries must be installed, the following script will download them and add them to your system PATH

```
make install_cli
```

## Usage

See the [PactFlow CI/CD Workshop](https://github.com/pactflow/ci-cd-workshop).

```
make fake_ci
```
