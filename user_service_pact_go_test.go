package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/utils"
)

func TestPactGoProvider(t *testing.T) {
	go startProvider()

	verifier := provider.NewVerifier()
	verifyRequest := provider.VerifyRequest{
		Provider:        "pactflow-example-provider-golang",
		ProviderBaseURL: fmt.Sprintf("http://127.0.0.1:%d", port),
	}
	verifyRequest.PactFiles = []string{os.Getenv("PACT_FILE")}

	err := verifier.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatalf("%v", err)
	}

}

func startProvider() {
	router := gin.Default()
	router.GET("/product/:id", GetProduct)
	router.Run(fmt.Sprintf(":%d", port))
}

// Configuration / Test Data
var port, _ = utils.GetFreePort()
