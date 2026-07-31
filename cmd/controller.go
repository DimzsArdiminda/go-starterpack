package cmd

func MakeController() {
	name, ok := generatorName("make:controller")
	if !ok {
		return
	}

	writeTemplate("template/controller.tmpl", "Controller/"+snakeCase(name)+".controller.go", templateData{Name: name})
}
