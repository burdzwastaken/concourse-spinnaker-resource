package concourse

type Source struct {
	SpinnakerAPI string `json:"spinnaker_api"`
	X509Cert     string `json:"spinnaker_x509_cert"`
	X509Key      string `json:"spinnaker_x509_key"`
}

type Version struct {
	ExecutionID string `json:"execution_id,omitempty"` // optional
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OutParams struct {
	SpinnakerApplication string            `json:"spinnaker_application"`
	SpinnakerPipeline    string            `json:"spinnaker_pipeline"`
	TriggerParams        map[string]string `json:"trigger_params,omitempty"` // optional
}

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutResponse struct {
	Version  Version        `json:"version"`
	Metadata []MetadataPair `json:"metadata"`
}
