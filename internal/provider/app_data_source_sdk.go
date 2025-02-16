// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-app/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppDataSourceModel) RefreshFromSharedApp(resp *shared.App) {
	if resp != nil {
		if resp.AccessLevel != nil {
			r.AccessLevel = types.StringValue(string(*resp.AccessLevel))
		} else {
			r.AccessLevel = types.StringNull()
		}
		r.AppID = types.StringPointerValue(resp.AppID)
		if resp.Author == nil {
			r.Author = nil
		} else {
			r.Author = &tfTypes.Author{}
			r.Author.Company = types.StringValue(resp.Author.Company)
			r.Author.Email = types.StringPointerValue(resp.Author.Email)
			r.Author.Name = types.StringPointerValue(resp.Author.Name)
		}
		r.Components = []tfTypes.BaseComponent{}
		if len(r.Components) > len(resp.Components) {
			r.Components = r.Components[:len(resp.Components)]
		}
		for componentsCount, componentsItem := range resp.Components {
			var components1 tfTypes.BaseComponent
			if componentsItem.Schemas != nil {
				components1.CustomJourneyBlock = &tfTypes.Schemas{}
				components1.CustomJourneyBlock.ComponentType = types.StringValue(string(componentsItem.Schemas.ComponentType))
				components1.CustomJourneyBlock.Configuration.ComponentTag = types.StringValue(componentsItem.Schemas.Configuration.ComponentTag)
				components1.CustomJourneyBlock.Configuration.ComponentURL = types.StringValue(componentsItem.Schemas.Configuration.ComponentURL)
				components1.CustomJourneyBlock.ID = types.StringValue(componentsItem.Schemas.ID)
				if componentsItem.Schemas.Name == nil {
					components1.CustomJourneyBlock.Name = nil
				} else {
					components1.CustomJourneyBlock.Name = &tfTypes.TranslatedString{}
					components1.CustomJourneyBlock.Name.De = types.StringPointerValue(componentsItem.Schemas.Name.De)
					components1.CustomJourneyBlock.Name.En = types.StringPointerValue(componentsItem.Schemas.Name.En)
				}
				components1.CustomJourneyBlock.Options = []tfTypes.Options{}
				for optionsVarCount, optionsVarItem := range componentsItem.Schemas.Options {
					var optionsVar1 tfTypes.Options
					optionsVar1.Description = types.StringPointerValue(optionsVarItem.Description)
					optionsVar1.Key = types.StringValue(optionsVarItem.Key)
					optionsVar1.Label = types.StringPointerValue(optionsVarItem.Label)
					optionsVar1.Required = types.BoolPointerValue(optionsVarItem.Required)
					optionsVar1.Type = types.StringValue(string(optionsVarItem.Type))
					optionsVar1.Value = types.StringPointerValue(optionsVarItem.Value)
					if optionsVarCount+1 > len(components1.CustomJourneyBlock.Options) {
						components1.CustomJourneyBlock.Options = append(components1.CustomJourneyBlock.Options, optionsVar1)
					} else {
						components1.CustomJourneyBlock.Options[optionsVarCount].Description = optionsVar1.Description
						components1.CustomJourneyBlock.Options[optionsVarCount].Key = optionsVar1.Key
						components1.CustomJourneyBlock.Options[optionsVarCount].Label = optionsVar1.Label
						components1.CustomJourneyBlock.Options[optionsVarCount].Required = optionsVar1.Required
						components1.CustomJourneyBlock.Options[optionsVarCount].Type = optionsVar1.Type
						components1.CustomJourneyBlock.Options[optionsVarCount].Value = optionsVar1.Value
					}
				}
			}
			if componentsItem.PortalExtensionComponentSchemas != nil {
				components1.PortalExtension = &tfTypes.PortalExtensionComponentSchemas{}
				components1.PortalExtension.ComponentType = types.StringValue(string(componentsItem.PortalExtensionComponentSchemas.ComponentType))
				components1.PortalExtension.Configuration.Hooks = []tfTypes.Hooks{}
				for hooksCount, hooksItem := range componentsItem.PortalExtensionComponentSchemas.Configuration.Hooks {
					var hooks1 tfTypes.Hooks
					if hooksItem.Auth == nil {
						hooks1.Auth = nil
					} else {
						hooks1.Auth = &tfTypes.PortalAuth{}
						if len(hooksItem.Auth.Headers) > 0 {
							hooks1.Auth.Headers = make(map[string]types.String)
							for key1, value1 := range hooksItem.Auth.Headers {
								hooks1.Auth.Headers[key1] = types.StringValue(value1)
							}
						}
						hooks1.Auth.Method = types.StringPointerValue(hooksItem.Auth.Method)
						hooks1.Auth.Type = types.StringPointerValue(hooksItem.Auth.Type)
						hooks1.Auth.URL = types.StringPointerValue(hooksItem.Auth.URL)
					}
					if hooksItem.Call == nil {
						hooks1.Call = nil
					} else {
						hooks1.Call = &tfTypes.Call{}
						if len(hooksItem.Call.Headers) > 0 {
							hooks1.Call.Headers = make(map[string]types.String)
							for key2, value2 := range hooksItem.Call.Headers {
								hooks1.Call.Headers[key2] = types.StringValue(value2)
							}
						}
						if len(hooksItem.Call.Params) > 0 {
							hooks1.Call.Params = make(map[string]types.String)
							for key3, value3 := range hooksItem.Call.Params {
								hooks1.Call.Params[key3] = types.StringValue(value3)
							}
						}
						hooks1.Call.URL = types.StringPointerValue(hooksItem.Call.URL)
					}
					hooks1.ID = types.StringPointerValue(hooksItem.ID)
					hooks1.Interval = []types.String{}
					for _, v := range hooksItem.Interval {
						hooks1.Interval = append(hooks1.Interval, types.StringValue(v))
					}
					if hooksItem.Name == nil {
						hooks1.Name = nil
					} else {
						hooks1.Name = &tfTypes.TranslatedString{}
						hooks1.Name.De = types.StringPointerValue(hooksItem.Name.De)
						hooks1.Name.En = types.StringPointerValue(hooksItem.Name.En)
					}
					hooks1.Type = types.StringPointerValue(hooksItem.Type)
					if hooksCount+1 > len(components1.PortalExtension.Configuration.Hooks) {
						components1.PortalExtension.Configuration.Hooks = append(components1.PortalExtension.Configuration.Hooks, hooks1)
					} else {
						components1.PortalExtension.Configuration.Hooks[hooksCount].Auth = hooks1.Auth
						components1.PortalExtension.Configuration.Hooks[hooksCount].Call = hooks1.Call
						components1.PortalExtension.Configuration.Hooks[hooksCount].ID = hooks1.ID
						components1.PortalExtension.Configuration.Hooks[hooksCount].Interval = hooks1.Interval
						components1.PortalExtension.Configuration.Hooks[hooksCount].Name = hooks1.Name
						components1.PortalExtension.Configuration.Hooks[hooksCount].Type = hooks1.Type
					}
				}
				components1.PortalExtension.Configuration.ID = types.StringPointerValue(componentsItem.PortalExtensionComponentSchemas.Configuration.ID)
				components1.PortalExtension.Configuration.Links = []tfTypes.Links{}
				for linksCount, linksItem := range componentsItem.PortalExtensionComponentSchemas.Configuration.Links {
					var links1 tfTypes.Links
					if linksItem.Auth == nil {
						links1.Auth = nil
					} else {
						links1.Auth = &tfTypes.PortalAuth{}
						if len(linksItem.Auth.Headers) > 0 {
							links1.Auth.Headers = make(map[string]types.String)
							for key4, value4 := range linksItem.Auth.Headers {
								links1.Auth.Headers[key4] = types.StringValue(value4)
							}
						}
						links1.Auth.Method = types.StringPointerValue(linksItem.Auth.Method)
						links1.Auth.Type = types.StringPointerValue(linksItem.Auth.Type)
						links1.Auth.URL = types.StringPointerValue(linksItem.Auth.URL)
					}
					links1.Condition = types.StringPointerValue(linksItem.Condition)
					if linksItem.Description == nil {
						links1.Description = nil
					} else {
						links1.Description = &tfTypes.TranslatedString{}
						links1.Description.De = types.StringPointerValue(linksItem.Description.De)
						links1.Description.En = types.StringPointerValue(linksItem.Description.En)
					}
					links1.ID = types.StringPointerValue(linksItem.ID)
					if linksItem.Name == nil {
						links1.Name = nil
					} else {
						links1.Name = &tfTypes.TranslatedString{}
						links1.Name.De = types.StringPointerValue(linksItem.Name.De)
						links1.Name.En = types.StringPointerValue(linksItem.Name.En)
					}
					if linksItem.Redirect == nil {
						links1.Redirect = nil
					} else {
						links1.Redirect = &tfTypes.Redirect{}
						if len(linksItem.Redirect.Params) > 0 {
							links1.Redirect.Params = make(map[string]types.String)
							for key5, value5 := range linksItem.Redirect.Params {
								links1.Redirect.Params[key5] = types.StringValue(value5)
							}
						}
						links1.Redirect.URL = types.StringPointerValue(linksItem.Redirect.URL)
					}
					links1.Type = types.StringPointerValue(linksItem.Type)
					if linksCount+1 > len(components1.PortalExtension.Configuration.Links) {
						components1.PortalExtension.Configuration.Links = append(components1.PortalExtension.Configuration.Links, links1)
					} else {
						components1.PortalExtension.Configuration.Links[linksCount].Auth = links1.Auth
						components1.PortalExtension.Configuration.Links[linksCount].Condition = links1.Condition
						components1.PortalExtension.Configuration.Links[linksCount].Description = links1.Description
						components1.PortalExtension.Configuration.Links[linksCount].ID = links1.ID
						components1.PortalExtension.Configuration.Links[linksCount].Name = links1.Name
						components1.PortalExtension.Configuration.Links[linksCount].Redirect = links1.Redirect
						components1.PortalExtension.Configuration.Links[linksCount].Type = links1.Type
					}
				}
				components1.PortalExtension.ID = types.StringValue(componentsItem.PortalExtensionComponentSchemas.ID)
				if componentsItem.PortalExtensionComponentSchemas.Name == nil {
					components1.PortalExtension.Name = nil
				} else {
					components1.PortalExtension.Name = &tfTypes.TranslatedString{}
					components1.PortalExtension.Name.De = types.StringPointerValue(componentsItem.PortalExtensionComponentSchemas.Name.De)
					components1.PortalExtension.Name.En = types.StringPointerValue(componentsItem.PortalExtensionComponentSchemas.Name.En)
				}
				components1.PortalExtension.Options = []tfTypes.Options{}
				for optionsVarCount1, optionsVarItem1 := range componentsItem.PortalExtensionComponentSchemas.Options {
					var optionsVar2 tfTypes.Options
					optionsVar2.Description = types.StringPointerValue(optionsVarItem1.Description)
					optionsVar2.Key = types.StringValue(optionsVarItem1.Key)
					optionsVar2.Label = types.StringPointerValue(optionsVarItem1.Label)
					optionsVar2.Required = types.BoolPointerValue(optionsVarItem1.Required)
					optionsVar2.Type = types.StringValue(string(optionsVarItem1.Type))
					optionsVar2.Value = types.StringPointerValue(optionsVarItem1.Value)
					if optionsVarCount1+1 > len(components1.PortalExtension.Options) {
						components1.PortalExtension.Options = append(components1.PortalExtension.Options, optionsVar2)
					} else {
						components1.PortalExtension.Options[optionsVarCount1].Description = optionsVar2.Description
						components1.PortalExtension.Options[optionsVarCount1].Key = optionsVar2.Key
						components1.PortalExtension.Options[optionsVarCount1].Label = optionsVar2.Label
						components1.PortalExtension.Options[optionsVarCount1].Required = optionsVar2.Required
						components1.PortalExtension.Options[optionsVarCount1].Type = optionsVar2.Type
						components1.PortalExtension.Options[optionsVarCount1].Value = optionsVar2.Value
					}
				}
				if componentsItem.PortalExtensionComponentSchemas.Origin != nil {
					components1.PortalExtension.Origin = types.StringValue(string(*componentsItem.PortalExtensionComponentSchemas.Origin))
				} else {
					components1.PortalExtension.Origin = types.StringNull()
				}
			}
			if componentsCount+1 > len(r.Components) {
				r.Components = append(r.Components, components1)
			} else {
				r.Components[componentsCount].CustomJourneyBlock = components1.CustomJourneyBlock
				r.Components[componentsCount].PortalExtension = components1.PortalExtension
			}
		}
		r.CreatedAt = types.StringPointerValue(resp.CreatedAt)
		r.CreatedBy = types.StringPointerValue(resp.CreatedBy)
		if resp.Description == nil {
			r.Description = nil
		} else {
			r.Description = &tfTypes.TranslatedString{}
			r.Description.De = types.StringPointerValue(resp.Description.De)
			r.Description.En = types.StringPointerValue(resp.Description.En)
		}
		r.DocumentationURL = types.StringPointerValue(resp.DocumentationURL)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.IconURL = types.StringPointerValue(resp.IconURL)
		r.InstallationID = types.StringPointerValue(resp.InstallationID)
		r.InstalledAt = types.StringPointerValue(resp.InstalledAt)
		r.InstalledBy = types.StringPointerValue(resp.InstalledBy)
		r.Internal = types.BoolPointerValue(resp.Internal)
		if resp.Name == nil {
			r.Name = nil
		} else {
			r.Name = &tfTypes.TranslatedString{}
			r.Name.De = types.StringPointerValue(resp.Name.De)
			r.Name.En = types.StringPointerValue(resp.Name.En)
		}
		r.OptionValues = []tfTypes.OptionsRef{}
		if len(r.OptionValues) > len(resp.OptionValues) {
			r.OptionValues = r.OptionValues[:len(resp.OptionValues)]
		}
		for optionValuesCount, optionValuesItem := range resp.OptionValues {
			var optionValues1 tfTypes.OptionsRef
			optionValues1.ComponentID = types.StringValue(optionValuesItem.ComponentID)
			optionValues1.Options = []tfTypes.Option{}
			for optionsVarCount2, optionsVarItem2 := range optionValuesItem.Options {
				var optionsVar3 tfTypes.Option
				optionsVar3.Key = types.StringValue(optionsVarItem2.Key)
				optionsVar3.Value = types.StringValue(optionsVarItem2.Value)
				if optionsVarCount2+1 > len(optionValues1.Options) {
					optionValues1.Options = append(optionValues1.Options, optionsVar3)
				} else {
					optionValues1.Options[optionsVarCount2].Key = optionsVar3.Key
					optionValues1.Options[optionsVarCount2].Value = optionsVar3.Value
				}
			}
			if optionValuesCount+1 > len(r.OptionValues) {
				r.OptionValues = append(r.OptionValues, optionValues1)
			} else {
				r.OptionValues[optionValuesCount].ComponentID = optionValues1.ComponentID
				r.OptionValues[optionValuesCount].Options = optionValues1.Options
			}
		}
		r.OrganizationID = types.StringPointerValue(resp.OrganizationID)
		r.OwnerOrgID = types.StringPointerValue(resp.OwnerOrgID)
		if resp.Status != nil {
			r.Status = types.StringValue(string(*resp.Status))
		} else {
			r.Status = types.StringNull()
		}
		r.UpdatedAt = types.StringPointerValue(resp.UpdatedAt)
		r.UpdatedBy = types.StringPointerValue(resp.UpdatedBy)
		r.Version = types.StringPointerValue(resp.Version)
	}
}
