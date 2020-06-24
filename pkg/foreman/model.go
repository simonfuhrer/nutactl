package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Model"

type Model struct {
	// Inherits the base object's attributes
	ForemanObject
}
