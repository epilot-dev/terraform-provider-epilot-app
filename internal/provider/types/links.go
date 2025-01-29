// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Links struct {
	Auth        *PortalAuth       `tfsdk:"auth"`
	Condition   types.String      `tfsdk:"condition"`
	Description *TranslatedString `tfsdk:"description"`
	ID          types.String      `tfsdk:"id"`
	Name        *TranslatedString `tfsdk:"name"`
	Redirect    *Redirect         `tfsdk:"redirect"`
	Type        types.String      `tfsdk:"type"`
}
