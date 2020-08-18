package gpg_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/invidian/terraform-provider-gpg/gpg"
)

func TestProvider(t *testing.T) {
	if err := gpg.Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("validating provider internally: %v", err)
	}
}
