// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/shared"
	"net/http"
)

type ListConfigurationsRequest struct {
	// Page number for pagination
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of items per page
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=pageSize"`
}

func (l ListConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListConfigurationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListConfigurationsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListConfigurationsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type Pagination struct {
	Page     *int64 `json:"page,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	Total    *int64 `json:"total,omitempty"`
}

func (o *Pagination) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *Pagination) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *Pagination) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

// ListConfigurationsResponseBody - List of app configurations
type ListConfigurationsResponseBody struct {
	Configurations []shared.ConfigurationMetadata `json:"configurations,omitempty"`
	Pagination     *Pagination                    `json:"pagination,omitempty"`
}

func (o *ListConfigurationsResponseBody) GetConfigurations() []shared.ConfigurationMetadata {
	if o == nil {
		return nil
	}
	return o.Configurations
}

func (o *ListConfigurationsResponseBody) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type ListConfigurationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of app configurations
	Object *ListConfigurationsResponseBody
}

func (o *ListConfigurationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConfigurationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConfigurationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConfigurationsResponse) GetObject() *ListConfigurationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
