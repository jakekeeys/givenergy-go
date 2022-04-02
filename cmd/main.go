package main

//java -jar openapi-generator-cli-6.0.0-20211025.061654-22.jar generate -i v1.openapi.yaml -g go --additional-properties=packageName=givenergy

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/jakekeeys/givenergy-go"
)

func main() {
	configuration := givenergy.NewConfiguration()
	configuration.AddDefaultHeader("Accept", "application/json")
	configuration.AddDefaultHeader("Content-Type", "application/json")
	configuration.AddDefaultHeader("Authorization", "Bearer ")

	api_client := givenergy.NewAPIClient(configuration)

	execute, response, err := api_client.AccountApi.AccountGet(context.Background()).Execute()
	spew.Dump(execute, response, err)
}
