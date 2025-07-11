// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// OverrideDevMode - Override URL when app is in dev mode
type OverrideDevMode struct {
	// URL of the web component object in dev mode
	OverrideURL *string `json:"override_url,omitempty"`
}

func (o *OverrideDevMode) GetOverrideURL() *string {
	if o == nil {
		return nil
	}
	return o.OverrideURL
}
