package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/utils"
)

func parseDate(dateStr string) *time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}
	return &date
}

func TestPactProvider(t *testing.T) {
	go startProvider()

	verifier := provider.NewVerifier()

	// Verify the Provider - fetch pacts from PactFlow
	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		Provider:                   "pactflow-example-provider-golang",
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		BrokerURL:                  fmt.Sprint(os.Getenv("PACT_BROKER_BASE_URL")),
		ConsumerVersionSelectors:   getSelectors(),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		BrokerUsername:             os.Getenv("PACT_BROKER_USERNAME"),
		BrokerPassword:             os.Getenv("PACT_BROKER_PASSWORD"),
		PublishVerificationResults: envBool("PACT_BROKER_PUBLISH_VERIFICATION_RESULTS"),
		ProviderVersion:            os.Getenv("GIT_COMMIT"),
		StateHandlers:              stateHandlers,
		EnablePending:              true,
		IncludeWIPPactsSince:       parseDate("2024-01-01"),
		ProviderBranch:             os.Getenv("GIT_BRANCH"),
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

}

// Provider state handlers
var stateHandlers = models.StateHandlers{
	"a product with ID 10 exists": func(setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
		productRepository = productExists
		return models.ProviderStateResponse{}, nil
	},
	"no products exist": func(setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
		productRepository = noProductsExist
		return models.ProviderStateResponse{}, nil
	},
}

// Starts the provider API with hooks for provider states.
func startProvider() {
	router := gin.Default()
	router.GET("/product/:id", GetProduct)

	router.Run(fmt.Sprintf(":%d", port))
}

// Provider States data sets
var productExists = &ProductRepository{
	Products: map[string]*Product{
		"10": {
			Name:    "Pizza",
			ID:      "10",
			Type:    "food",
			Version: "1.0.0",
		},
	},
}

var noProductsExist = &ProductRepository{}

// Configuration / Test Data
var port, _ = utils.GetFreePort()

func getSelectors() []provider.Selector {
	selectors := make([]provider.Selector, 0)
	if os.Getenv("SELECTORS") != "" {
		selectors = []provider.Selector{
			&provider.ConsumerVersionSelector{
				DeployedOrReleased: true,
			},
			&provider.ConsumerVersionSelector{
				MainBranch: true,
			},
		}
	}

	return selectors
}

func envBool(k string) bool {
	return os.Getenv(k) != ""
}
