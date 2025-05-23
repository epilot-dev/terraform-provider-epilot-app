// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-app/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppResourceModel) ToSharedInstallRequest() *shared.InstallRequest {
	var optionValues []shared.OptionsRef = []shared.OptionsRef{}
	for _, optionValuesItem := range r.OptionValues {
		var componentID string
		componentID = optionValuesItem.ComponentID.ValueString()

		var optionsVar []shared.Option = []shared.Option{}
		for _, optionsItem := range optionValuesItem.Options {
			var key string
			key = optionsItem.Key.ValueString()

			var value string
			value = optionsItem.Value.ValueString()

			optionsVar = append(optionsVar, shared.Option{
				Key:   key,
				Value: value,
			})
		}
		optionValues = append(optionValues, shared.OptionsRef{
			ComponentID: componentID,
			Options:     optionsVar,
		})
	}
	version := new(string)
	if !r.Version.IsUnknown() && !r.Version.IsNull() {
		*version = r.Version.ValueString()
	} else {
		version = nil
	}
	out := shared.InstallRequest{
		OptionValues: optionValues,
		Version:      version,
	}
	return &out
}

func (r *AppResourceModel) RefreshFromSharedInstallation(resp *shared.Installation) {
	if resp != nil {
		r.Manifest = []types.String{}
		for _, v := range resp.Manifest {
			r.Manifest = append(r.Manifest, types.StringValue(v))
		}
		r.AppID = types.StringValue(resp.AppID)
		componentsResult, _ := json.Marshal(resp.Components)
		r.Components = types.StringValue(string(componentsResult))
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		if resp.InstallationAudit == nil {
			r.InstallationAudit = nil
		} else {
			r.InstallationAudit = &tfTypes.InstallationAudit{}
			r.InstallationAudit.CreatedAt = types.StringPointerValue(resp.InstallationAudit.CreatedAt)
			r.InstallationAudit.CreatedBy = types.StringPointerValue(resp.InstallationAudit.CreatedBy)
			r.InstallationAudit.UpdatedAt = types.StringPointerValue(resp.InstallationAudit.UpdatedAt)
			r.InstallationAudit.UpdatedBy = types.StringPointerValue(resp.InstallationAudit.UpdatedBy)
		}
		r.InstalledVersion = types.StringValue(resp.InstalledVersion)
		r.InstallerOrgID = types.StringValue(resp.InstallerOrgID)
		r.Name = types.StringValue(resp.Name)
		r.OptionValues = []tfTypes.OptionsRef{}
		if len(r.OptionValues) > len(resp.OptionValues) {
			r.OptionValues = r.OptionValues[:len(resp.OptionValues)]
		}
		for optionValuesCount, optionValuesItem := range resp.OptionValues {
			var optionValues1 tfTypes.OptionsRef
			optionValues1.ComponentID = types.StringValue(optionValuesItem.ComponentID)
			optionValues1.Options = []tfTypes.Option{}
			for optionsVarCount, optionsVarItem := range optionValuesItem.Options {
				var optionsVar1 tfTypes.Option
				optionsVar1.Key = types.StringValue(optionsVarItem.Key)
				optionsVar1.Value = types.StringValue(optionsVarItem.Value)
				if optionsVarCount+1 > len(optionValues1.Options) {
					optionValues1.Options = append(optionValues1.Options, optionsVar1)
				} else {
					optionValues1.Options[optionsVarCount].Key = optionsVar1.Key
					optionValues1.Options[optionsVarCount].Value = optionsVar1.Value
				}
			}
			if optionValuesCount+1 > len(r.OptionValues) {
				r.OptionValues = append(r.OptionValues, optionValues1)
			} else {
				r.OptionValues[optionValuesCount].ComponentID = optionValues1.ComponentID
				r.OptionValues[optionValuesCount].Options = optionValues1.Options
			}
		}
		r.Role = types.StringPointerValue(resp.Role)
	}
}
