package cmd

import "fmt"

func MakeCrud() {
	name, ok := generatorName("make:crud")
	if !ok {
		return
	}

	writeTemplate("template/model.tmpl", "Model/"+snakeCase(name)+".model.go", templateData{Name: name, TableName: pluralize(snakeCase(name))})
	writeTemplate("template/controller.tmpl", "Controller/"+snakeCase(name)+".controller.go", templateData{Name: name})
	fmt.Printf("Register %sController in Routes/routes.go when you are ready.\n", name)
}
