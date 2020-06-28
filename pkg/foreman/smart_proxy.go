package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=SmartProxy Value=Name Path=smart_proxies"

type SmartProxy struct {
	// Inherits the base object's attributes
	ForemanObject
	URL      string               `json:"url,omitempty"`
	Features []SmartProxyFeatures `json:"features,omitempty"`
}

type SmartProxyFeatures struct {
	Name         string      `json:"name,omitempty"`
	ID           int         `json:"id"`
	Capabilities interface{} `json:"capabilities"`
}
