// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BlueprintRef struct {
	// ID of the job that created the blueprint
	JobID *string `json:"job_id,omitempty"`
	// ID of the blueprint
	ManifestID *string `json:"manifest_id,omitempty"`
}

func (o *BlueprintRef) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *BlueprintRef) GetManifestID() *string {
	if o == nil {
		return nil
	}
	return o.ManifestID
}
