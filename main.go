package main

import (
	"log"
	"os"

	"github.com/EvilSuperstars/terraform-provider-wapc/provider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
)

func main() {
	if err := tf5server.Serve("wapc", provider.Server); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
