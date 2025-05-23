// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Grants struct {
	// The action the app can perform
	Action string `json:"action"`
	// The resource the app can access
	Resource *string `json:"resource,omitempty"`
}

func (o *Grants) GetAction() string {
	if o == nil {
		return ""
	}
	return o.Action
}

func (o *Grants) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}
