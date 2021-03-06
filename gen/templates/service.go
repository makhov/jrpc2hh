package templates

var Service string = `package {{.Package}}
//-------------------------------------------------------------------------------//
// This file was generated by andrskom/jrpc2hh package. Please don't remove it. //
//-------------------------------------------------------------------------------//

import (
	"fmt"
	{{range $index, $element := .Imports}}{{$element}} "{{$index}}"
	{{end}}
)

func (s *{{.Service}}) Call(reqBody *jModels.RequestBody) (interface{}, *jModels.Error) {
	switch reqBody.GetMethod() {
	{{range $element := .Methods}}{{$element}}
	{{end}}default:
		return nil, jModels.NewError(jModels.ErrorCodeMethodNotFound, fmt.Sprintf("Unknown method '%s' for service '%s'", reqBody.GetMethod(), "{{.Service}}"), nil)
	}
}`
