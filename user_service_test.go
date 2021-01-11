package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

func TestPactProvider(t *testing.T) {
	go startProvider()

	pact := createPact()

	selectors := make([]types.ConsumerVersionSelector, 0)
	if os.Getenv("SELECTORS") != "" {
		selectors = []types.ConsumerVersionSelector{
			{
				Tag: "master",
			},
			{
				Tag: "prod",
			},
		}
	}

	// Verify the Provider - Tag-based Published Pacts for any known consumers
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		BrokerURL:                  fmt.Sprintf(os.Getenv("PACT_BROKER_BASE_URL")),
		ConsumerVersionSelectors:   selectors,
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		BrokerUsername:             os.Getenv("PACT_BROKER_USERNAME"),
		BrokerPassword:             os.Getenv("PACT_BROKER_PASSWORD"),
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		StateHandlers:              stateHandlers,
		EnablePending:              isPending(),
	})

	if err != nil {
		t.Fatalf("%v", err)
	}
}

var token = "" // token will be dynamic based on state etc.

// Provider state handlers
var stateHandlers = types.StateHandlers{
	"a product with ID 10 exists": func() error {
		productRepository = productExists
		return nil
	},
	"no products exist": func() error {
		productRepository = noProductsExist
		return nil
	},
}

// Starts the provider API with hooks for provider states.
// This essentially mirrors the main.go file, with extra routes added.
func startProvider() {
	router := gin.Default()
	router.GET("/product/:id", GetProduct)

	router.Run(fmt.Sprintf(":%d", port))
}

// Configuration / Test Data
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

// Provider States data sets
var productExists = &ProductRepository{
	Products: map[string]*Product{
		"10": {
			Name: "Pizza",
			ID:   "10",
			Type: "food",
		},
	},
}

var noProductsExist = &ProductRepository{}

// Setup the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider:                 "pactflow-example-provider-golang",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
	}
}

func isPending() bool {
	if os.Getenv("PENDING") != "" {
		return true
	}
	return false
}
