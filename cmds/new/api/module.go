package api

const moduleTmpl = `
package {{.moduleName|mpkgName|lName}}

import "github.com/micro-plat/hydra/component"

type I{{.moduleName|lName|humpName}} interface {
}

type {{.moduleName|lName|humpName}} struct {
	c component.IContainer
}


func New{{.moduleName|lName|humpName}}(c component.IContainer) *{{.moduleName|lName|humpName}} {
	return &{{.moduleName|lName|humpName}}{
		c: c,
	}
}



`
const sqlTmpl = `
package sql
`
