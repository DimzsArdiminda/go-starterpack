package cmd

func MakeMiddleware() {
	name, ok := generatorName("make:middleware")
	if !ok {
		return
	}
	writeTemplate("template/middleware.tmpl", "Middleware/"+snakeCase(name)+".middleware.go", templateData{Name: name})
}
