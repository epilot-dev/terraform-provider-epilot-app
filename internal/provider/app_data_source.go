// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-app/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &AppDataSource{}
var _ datasource.DataSourceWithConfigure = &AppDataSource{}

func NewAppDataSource() datasource.DataSource {
	return &AppDataSource{}
}

// AppDataSource is the data source implementation.
type AppDataSource struct {
	client *sdk.SDK
}

// AppDataSourceModel describes the data model.
type AppDataSourceModel struct {
	AccessLevel      types.String              `tfsdk:"access_level"`
	AppID            types.String              `tfsdk:"app_id"`
	Author           *tfTypes.Author           `tfsdk:"author"`
	Components       []tfTypes.BaseComponent   `tfsdk:"components"`
	CreatedAt        types.String              `tfsdk:"created_at"`
	CreatedBy        types.String              `tfsdk:"created_by"`
	Description      *tfTypes.TranslatedString `tfsdk:"description"`
	DocumentationURL types.String              `tfsdk:"documentation_url"`
	Enabled          types.Bool                `tfsdk:"enabled"`
	IconURL          types.String              `tfsdk:"icon_url"`
	InstallationID   types.String              `tfsdk:"installation_id"`
	InstalledAt      types.String              `tfsdk:"installed_at"`
	InstalledBy      types.String              `tfsdk:"installed_by"`
	Internal         types.Bool                `tfsdk:"internal"`
	Name             *tfTypes.TranslatedString `tfsdk:"name"`
	OptionValues     []tfTypes.OptionsRef      `tfsdk:"option_values"`
	OrganizationID   types.String              `tfsdk:"organization_id"`
	OwnerOrgID       types.String              `tfsdk:"owner_org_id"`
	Status           types.String              `tfsdk:"status"`
	UpdatedAt        types.String              `tfsdk:"updated_at"`
	UpdatedBy        types.String              `tfsdk:"updated_by"`
	Version          types.String              `tfsdk:"version"`
}

// Metadata returns the data source type name.
func (r *AppDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

// Schema defines the schema for the data source.
func (r *AppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App DataSource",

		Attributes: map[string]schema.Attribute{
			"access_level": schema.StringAttribute{
				Computed:    true,
				Description: `Access level of the app.`,
			},
			"app_id": schema.StringAttribute{
				Required: true,
			},
			"author": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"company": schema.StringAttribute{
						Computed:    true,
						Description: `Company of the author`,
					},
					"email": schema.StringAttribute{
						Computed:    true,
						Description: `Email of the author`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the author`,
					},
				},
			},
			"components": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"custom_journey_block": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"component_type": schema.StringAttribute{
									Computed: true,
								},
								"configuration": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"component_tag": schema.StringAttribute{
											Computed:    true,
											Description: `Custom element tag for the component`,
										},
										"component_url": schema.StringAttribute{
											Computed:    true,
											Description: `URL of the web component object`,
										},
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `Unique identifier for the component`,
								},
								"name": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"de": schema.StringAttribute{
											Computed:    true,
											Description: `German translation`,
										},
										"en": schema.StringAttribute{
											Computed:    true,
											Description: `English translation`,
										},
									},
								},
								"options": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed:    true,
												Description: `Detailed description of what this configuration option does`,
											},
											"key": schema.StringAttribute{
												Computed:    true,
												Description: `Unique identifier for this configuration option`,
											},
											"label": schema.StringAttribute{
												Computed:    true,
												Description: `Human-readable label for the configuration option`,
											},
											"required": schema.BoolAttribute{
												Computed:    true,
												Description: `Flag to indicate if this option is required`,
											},
											"type": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed:    true,
												Description: `The configured value for this option. Is only present when the component is installed.`,
											},
										},
									},
									Description: `List of options for the app component`,
								},
							},
						},
						"portal_extension": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"component_type": schema.StringAttribute{
									Computed: true,
								},
								"configuration": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"hooks": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"auth": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"headers": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"method": schema.StringAttribute{
																Computed: true,
															},
															"type": schema.StringAttribute{
																Computed: true,
															},
															"url": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"call": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"headers": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"params": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"url": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"id": schema.StringAttribute{
														Computed: true,
													},
													"interval": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
													},
													"name": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"de": schema.StringAttribute{
																Computed:    true,
																Description: `German translation`,
															},
															"en": schema.StringAttribute{
																Computed:    true,
																Description: `English translation`,
															},
														},
													},
													"type": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"links": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"auth": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"headers": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"method": schema.StringAttribute{
																Computed: true,
															},
															"type": schema.StringAttribute{
																Computed: true,
															},
															"url": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"condition": schema.StringAttribute{
														Computed: true,
													},
													"description": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"de": schema.StringAttribute{
																Computed:    true,
																Description: `German translation`,
															},
															"en": schema.StringAttribute{
																Computed:    true,
																Description: `English translation`,
															},
														},
													},
													"id": schema.StringAttribute{
														Computed: true,
													},
													"name": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"de": schema.StringAttribute{
																Computed:    true,
																Description: `German translation`,
															},
															"en": schema.StringAttribute{
																Computed:    true,
																Description: `English translation`,
															},
														},
													},
													"redirect": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"params": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"url": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"type": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `Unique identifier for the component`,
								},
								"name": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"de": schema.StringAttribute{
											Computed:    true,
											Description: `German translation`,
										},
										"en": schema.StringAttribute{
											Computed:    true,
											Description: `English translation`,
										},
									},
								},
								"options": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed:    true,
												Description: `Detailed description of what this configuration option does`,
											},
											"key": schema.StringAttribute{
												Computed:    true,
												Description: `Unique identifier for this configuration option`,
											},
											"label": schema.StringAttribute{
												Computed:    true,
												Description: `Human-readable label for the configuration option`,
											},
											"required": schema.BoolAttribute{
												Computed:    true,
												Description: `Flag to indicate if this option is required`,
											},
											"type": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed:    true,
												Description: `The configured value for this option. Is only present when the component is installed.`,
											},
										},
									},
									Description: `List of options for the app component`,
								},
								"origin": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"de": schema.StringAttribute{
						Computed:    true,
						Description: `German translation`,
					},
					"en": schema.StringAttribute{
						Computed:    true,
						Description: `English translation`,
					},
				},
			},
			"documentation_url": schema.StringAttribute{
				Computed:    true,
				Description: `URL of the app documentation.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Flag to indicate if the app is enabled.`,
			},
			"icon_url": schema.StringAttribute{
				Computed:    true,
				Description: `URL of the app icon.`,
			},
			"installation_id": schema.StringAttribute{
				Computed:    true,
				Description: `Unique identifier for the app installation`,
			},
			"installed_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp of app creation`,
			},
			"installed_by": schema.StringAttribute{
				Computed:    true,
				Description: `User ID of the user who installed the app`,
			},
			"internal": schema.BoolAttribute{
				Computed:    true,
				Description: `Flag to indicate if the app is built by epilot.`,
			},
			"name": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"de": schema.StringAttribute{
						Computed:    true,
						Description: `German translation`,
					},
					"en": schema.StringAttribute{
						Computed:    true,
						Description: `English translation`,
					},
				},
			},
			"option_values": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"component_id": schema.StringAttribute{
							Computed:    true,
							Description: `ID of the component these values are for`,
						},
						"options": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed:    true,
										Description: `Key matching a config_option from the component`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The configured value for this option`,
									},
								},
							},
						},
					},
				},
				Description: `Configuration values for the app components options`,
			},
			"organization_id": schema.StringAttribute{
				Computed:    true,
				Description: `Unique identifier for the organization the app is installed in`,
			},
			"owner_org_id": schema.StringAttribute{
				Computed:    true,
				Description: `Organization ID of the app owner, required for private apps`,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp of the last update`,
			},
			"updated_by": schema.StringAttribute{
				Computed:    true,
				Description: `User ID of the user who last updated the app`,
			},
			"version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AppDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var appID string
	appID = data.AppID.ValueString()

	request := operations.GetInstalledAppRequest{
		AppID: appID,
	}
	res, err := r.client.GetInstalledApp(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.App != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
