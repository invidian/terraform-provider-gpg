package gpg

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider exports terraform-provider-gpg, which can be used in tests
// for other providers.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gpg_encrypted_message": resourceGPGEncryptedMessage(),
		},
	}
}
