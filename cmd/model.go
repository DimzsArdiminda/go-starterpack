package cmd

func MakeModel() {
	name, ok := generatorName("make:model")
	if !ok {
		return
	}

	writeTemplate("template/model.tmpl", "Model/"+snakeCase(name)+".model.go", templateData{
		Name:      name,
		TableName: pluralize(snakeCase(name)),
	})
}
