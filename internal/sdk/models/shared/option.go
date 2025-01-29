// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Option struct {
	// Key matching a config_option from the component
	Key string `json:"key"`
	// The configured value for this option
	Value string `json:"value"`
}

func (o *Option) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *Option) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}
