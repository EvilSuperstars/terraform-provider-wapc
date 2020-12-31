package main

import (
	"context"
	"log"
	"os"

	"github.com/EvilSuperstars/terraform-provider-wapc/provider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"
)

func main() {
	ctx := context.Background()

	factory, err := tfmux.NewSchemaServerFactory(ctx, provider.Server)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = tf5server.Serve("wapc", func() tfprotov5.ProviderServer {
		return factory.Server()
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
