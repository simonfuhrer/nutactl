package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=User"

type User struct {
	// Inherits the base object's attributes
	ForemanObject
}
