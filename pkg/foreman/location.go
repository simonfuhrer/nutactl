package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Location Value=Name Path=locations"

type Location struct {
	// Inherits the base object's attributes
	ForemanObject
}
