package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Environment Value=Name Path=environments"

type Environment struct {
	// Inherits the base object's attributes
	ForemanObject

	TemplateCombinations interface{} `json:"template_combinations,omit_empty"`
	Puppetclasses        interface{} `json:"puppetclasses,omit_empty"`
}
