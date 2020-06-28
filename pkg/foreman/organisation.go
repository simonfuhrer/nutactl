package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Organization Value=Name Path=organizations"

type Organization struct {
	// Inherits the base object's attributes
	ForemanObject
}
