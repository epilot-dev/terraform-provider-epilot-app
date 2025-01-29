// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/internal/utils"
)

// AppConfigurationAccessLevel - Access level of the app.
type AppConfigurationAccessLevel string

const (
	AppConfigurationAccessLevelPublic  AppConfigurationAccessLevel = "public"
	AppConfigurationAccessLevelPrivate AppConfigurationAccessLevel = "private"
)

func (e AppConfigurationAccessLevel) ToPointer() *AppConfigurationAccessLevel {
	return &e
}
func (e *AppConfigurationAccessLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		*e = AppConfigurationAccessLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppConfigurationAccessLevel: %v", v)
	}
}

type AppConfigurationStatus string

const (
	AppConfigurationStatusPublished AppConfigurationStatus = "published"
	AppConfigurationStatusPending   AppConfigurationStatus = "pending"
)

func (e AppConfigurationStatus) ToPointer() *AppConfigurationStatus {
	return &e
}
func (e *AppConfigurationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "published":
		fallthrough
	case "pending":
		*e = AppConfigurationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppConfigurationStatus: %v", v)
	}
}

// AppConfiguration - Configuration of the published app
type AppConfiguration struct {
	// Access level of the app.
	AccessLevel *AppConfigurationAccessLevel `default:"public" json:"access_level"`
	AppID       *string                      `json:"app_id,omitempty"`
	Author      *Author                      `json:"author,omitempty"`
	Components  []BaseComponent              `json:"components,omitempty"`
	CreatedAt   *string                      `json:"created_at,omitempty"`
	CreatedBy   *string                      `json:"created_by,omitempty"`
	Description *TranslatedString            `json:"description,omitempty"`
	// URL of the app documentation.
	DocumentationURL *string `json:"documentation_url,omitempty"`
	// URL of the app icon.
	IconURL *string `json:"icon_url,omitempty"`
	// Flag to indicate if the app is built by epilot.
	Internal *bool             `default:"false" json:"internal"`
	Name     *TranslatedString `json:"name,omitempty"`
	// Organization ID of the app owner, required for private apps
	OwnerOrgID *string                 `json:"owner_org_id,omitempty"`
	Status     *AppConfigurationStatus `json:"status,omitempty"`
	UpdatedAt  *string                 `json:"updated_at,omitempty"`
	UpdatedBy  *string                 `json:"updated_by,omitempty"`
	Version    *string                 `json:"version,omitempty"`
}

func (a AppConfiguration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppConfiguration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppConfiguration) GetAccessLevel() *AppConfigurationAccessLevel {
	if o == nil {
		return nil
	}
	return o.AccessLevel
}

func (o *AppConfiguration) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppConfiguration) GetAuthor() *Author {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *AppConfiguration) GetComponents() []BaseComponent {
	if o == nil {
		return nil
	}
	return o.Components
}

func (o *AppConfiguration) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppConfiguration) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *AppConfiguration) GetDescription() *TranslatedString {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppConfiguration) GetDocumentationURL() *string {
	if o == nil {
		return nil
	}
	return o.DocumentationURL
}

func (o *AppConfiguration) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *AppConfiguration) GetInternal() *bool {
	if o == nil {
		return nil
	}
	return o.Internal
}

func (o *AppConfiguration) GetName() *TranslatedString {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AppConfiguration) GetOwnerOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerOrgID
}

func (o *AppConfiguration) GetStatus() *AppConfigurationStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AppConfiguration) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AppConfiguration) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *AppConfiguration) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
