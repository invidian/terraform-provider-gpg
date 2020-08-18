package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/invidian/terraform-provider-gpg/gpg"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gpg.Provider,
	})
}
