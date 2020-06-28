package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Model Value=Name Path=models"

type Model struct {
	// Inherits the base object's attributes
	ForemanObject
}
