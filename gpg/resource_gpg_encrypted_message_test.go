package gpg_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/invidian/terraform-provider-gpg/gpg"
)

const basicConfig = `
resource "gpg_encrypted_message" "example" {
  content     = "This is example of GPG encrypted message."
  public_keys = [
    var.gpg_public_key,
  ]
}

variable "gpg_public_key" {
  default = <<EOF
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
=ZC5j
-----END PGP PUBLIC KEY BLOCK-----
EOF
}
`

const badPublicKey = `
resource "gpg_encrypted_message" "example" {
  content     = "This is example of GPG encrypted message."
  public_keys = [
		"not valid message",
  ]
}
`

const badPublicKeyPEMEncoded = `
resource "gpg_encrypted_message" "example" {
  content     = "This is example of GPG encrypted message."
	public_keys = [
		<<EOF
-----BEGIN PGP PUBLIC KEY BLOCK-----

bm9wZQo=
-----END PGP PUBLIC KEY BLOCK-----
EOF
		,
	]
}

`

const noPublicKeys = `
resource "gpg_encrypted_message" "example" {
  content     = "This is example of GPG encrypted message."
  public_keys = []
}
`

func TestGPGEncryptedMessage(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"gpg": gpg.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicConfig,
			},
			{
				Config:             basicConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config:      badPublicKey,
				ExpectError: regexp.MustCompile(`decoding public key: EOF`),
			},
			{
				Config:      badPublicKeyPEMEncoded,
				ExpectError: regexp.MustCompile(`parsing public key`),
			},
		},
	})
}

func TestGPGEncryptedMessageBadArguments(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"gpg": gpg.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config:      noPublicKeys,
				ExpectError: regexp.MustCompile(`attribute supports 1 item as a minimum`),
				Destroy:     false,
			},
		},
	})
}
