// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Author struct {
	// Company of the author
	Company string `json:"company"`
	// Email of the author
	Email *string `json:"email,omitempty"`
	// Name of the author
	Name *string `json:"name,omitempty"`
}

func (o *Author) GetCompany() string {
	if o == nil {
		return ""
	}
	return o.Company
}

func (o *Author) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Author) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
