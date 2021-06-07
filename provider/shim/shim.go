package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-docker/internal/provider"
)

func NewProvider() func() *schema.Provider {
	return provider.New("")
}