//go:build pact_go
// +build pact_go

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/utils"
)

func TestPactGoProvider(t *testing.T) {
	var port, _ = utils.GetFreePort()
	go startProvider(port)

	verifier := provider.NewVerifier()
	verifyRequest := provider.VerifyRequest{
		Provider:        "pactflow-example-provider-golang",
		ProviderBaseURL: fmt.Sprintf("http://127.0.0.1:%d", port),
	}
	pactFile := os.Getenv("PACT_FILE")
	if pactFile == "" {
		pactFile = "pact.json"
	}
	verifyRequest.PactFiles = []string{pactFile}

	err := verifier.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatalf("%v", err)
	}

}
