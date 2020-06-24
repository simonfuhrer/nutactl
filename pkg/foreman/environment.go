package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Environment"

type Environment struct {
	// Inherits the base object's attributes
	ForemanObject
}
