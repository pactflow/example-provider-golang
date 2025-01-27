package main

import (
	"testing"
	// "github.com/pact-foundation/pact-go/v2/provider"
)

func TestLibPactFfiProvider(t *testing.T) {
	go startProvider()

	var res = test_provider(port)
	if res != 0 {
		t.Fatalf("%v", res)
	}

}

// func startProvider() {
// 	router := gin.Default()
// 	router.GET("/product/:id", GetProduct)
// 	router.Run(fmt.Sprintf(":%d", port))
// }

// Configuration / Test Data
// var port, _ = utils.GetFreePort()
