package intermediate

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.New("example")

	tmpl, err := tmpl.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "John",
	}

	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
