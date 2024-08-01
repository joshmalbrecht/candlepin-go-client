package main

import (
	"context"

	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	serverConfig := openapiclient.ServerConfiguration{
		URL:         "https://localhost:8443/candlepin",
		Description: "testing",
	}

	config := openapiclient.Configuration{
		Servers: openapiclient.ServerConfigurations{serverConfig},
	}

	apiClient := openapiclient.NewAPIClient(&config)

	// TODO: I think we need to add BasicAuth to the context and then pass the context into OwnerAPI.CreateOwner

	createOwnerRequest := apiClient.OwnerAPI.CreateOwner(context.Background())

	var displayName string = "client-test-owner"
	owner := openapiclient.OwnerDTO{
		DisplayName: &displayName,
	}

	createOwnerRequest = createOwnerRequest.OwnerDTO(owner)

	createdOwner, response, err := createOwnerRequest.Execute()
	if err != nil {
		println("ERROR: " + err.Error())
	} else {
		println(response)
		println(createdOwner)
	}
}
