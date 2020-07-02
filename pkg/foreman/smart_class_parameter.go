package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=SmartClassParameter Value=Name Path=smart_class_parameters"

type SmartClassParameter struct {
	// Inherits the base object's attributes
	ForemanObject
	Description         string      `json:"description,omitempty"`
	Override            bool        `json:"override,omitempty"`
	ParameterType       string      `json:"parameter_type,omitempty"`
	HiddenValue         bool        `json:"hidden_value,omitempty"`
	Omit                interface{} `json:"omit,omitempty"`
	Required            bool        `json:"required,omitempty"`
	ValidatorType       string      `json:"validator_type,omitempty"`
	ValidatorRule       string      `json:"validator_rule,omitempty"`
	MergeOverrides      bool        `json:"merge_overrides,omitempty"`
	MergeDefault        bool        `json:"merge_default,omitempty"`
	AvoidDuplicates     bool        `json:"avoid_duplicates,omitempty"`
	OverrideValueOrder  string      `json:"override_value_order,omitempty"`
	UsePuppetDefault    bool        `json:"use_puppet_default,omitempty"`
	Parameter           string      `json:"parameter,omitempty"`
	PuppetclassID       int         `json:"puppetclass_id,omitempty"`
	OverrideValuesCount int         `json:"override_values_count,omitempty"`
	DefaultValue        interface{} `json:"default_value,omitempty"`
	PuppetclassName     string      `json:"puppetclass_name,omitempty"`
}
