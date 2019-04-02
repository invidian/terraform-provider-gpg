package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func resourceGPGEncryptedMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGPGEncryptedMessageCreate,
		// Those 2 functions below does nothing, but must be implemented.
		Read:   resourceGPGEncryptedMessageRead,
		Delete: resourceGPGEncryptedMessageDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				StateFunc: sha256sum,
			},
			"public_keys": &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					StateFunc: func(val interface{}) string {
						recipient, err := entityFromString(val.(string))
						if err != nil {
							// We only keep KeyId in state, as we want to keep it small and also
							// we always read public keys anyway. If public key is malformed,
							// creation of resource will fail anyway, so it's fine to set it here.
							return fmt.Sprintf("MALFORMED KEY")
						}

						// Instead of full ASCII-armored key, write only KeyId to state
						return recipient.PrimaryKey.KeyIdString()
					},
				},
				Required: true,
			},
			"result": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				StateFunc: sha256sum,
			},
		},
	}
}

func resourceGPGEncryptedMessageCreate(d *schema.ResourceData, m interface{}) error {
	// Store recipients for encryption
	recipients := []*openpgp.Entity{}
	// Store ID of each public key, to store them in state (StateFunc does not work for TypeList for some reason)
	pks_ids := []string{}

	// Iterate over public keys, decode, parse, collect their IDs and add to recipients list
	for i, pk := range d.Get("public_keys").([]interface{}) {
		recipient, err := entityFromString(pk.(string))
		if err != nil {
			return fmt.Errorf("Unable to decode public_keys[%d]: %v", i, err)
		}

		recipients = append(recipients, recipient)
		pks_ids = append(pks_ids, recipient.PrimaryKey.KeyIdString())
	}

	if err := d.Set("public_keys", pks_ids); err != nil {
		return fmt.Errorf("Failed to set public_keys property:  %v", err)
	}

	buf := new(bytes.Buffer)

	// We produce output in ASCII-armor format
	wc_encode, err := armor.Encode(buf, "PGP MESSAGE", nil)
	if err != nil {
		return fmt.Errorf("Encoding the message failed: %v", err)
	}
	wc_encrypt, err := openpgp.Encrypt(wc_encode, recipients, nil, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return fmt.Errorf("Encrypting the message failed: %v", err)
	}
	if _, err := io.Copy(wc_encrypt, strings.NewReader(d.Get("content").(string))); err != nil {
		return fmt.Errorf("Failed writing content to buffer: %v", err)
	}
	wc_encrypt.Close()
	wc_encode.Close()
	result := buf.String()
	if err := d.Set("result", result); err != nil {
		return fmt.Errorf("Failed to set result property: %v", err)
	}

	// Calculate SHA-256 checksum of message for ID
	d.SetId(fmt.Sprintf("%x", sha256sum(result)))

	return resourceGPGEncryptedMessageRead(d, m)
}

func resourceGPGEncryptedMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGPGEncryptedMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func entityFromString(key string) (*openpgp.Entity, error) {
	block, err := armor.Decode(strings.NewReader(key))
	if err != nil {
		return nil, fmt.Errorf("Unable to decode public_keys: %v", err)
	}

	recipient, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		return nil, fmt.Errorf("Unable to parse public_key: %v", err)
	}

	return recipient, nil
}

func sha256sum(data interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data.(string))))
}
