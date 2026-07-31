package cmd

func MakeRequest() {
	name, ok := generatorName("make:request")
	if !ok {
		return
	}
	writeTemplate("template/request.tmpl", "Request/"+snakeCase(name)+".request.go", templateData{Name: name})
}
