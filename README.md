# YamiGo
YamiGo is (by its own name) Yet Another Micro-framework In GO. Everything your simple web application requires, YamiGo can be used for.

### Usage

Copy `config.example.yml` to your application route directory as `config.development.yml`.

`config.development.yml`:
```yml
app:
  url: http://localhost
  port: 8080

views:
  template:
    baseDir: 'tmpl/'
    baseTemplate: 'base.html.tmpl'
```


Create your view templates:

`tmpl/base.html.yml`:
```html
{{ define "base" }}

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ Param "title" }}</title>
</head>
<body>
    {{ template "content" .}}
</body>
</html>
{{ end }}
```

`tmpl/index.html.yml`:
```html
{{ define "content" }}
<h1>Hello world</h1>
{{ end }}
```

Create the application:

`main.go`:
```golang
package main

import (
	"github.com/galaco/yamigo"
)

type Page struct {
    yamigo.Route
}

func (page Page) Execute(request yamigo.Request, response yamigo.Response) {
	template := yamigo.NewTemplate("index.html.tmpl")
	response.SetTemplate(template)

	response.Execute()
}

func main() {
	app := yamigo.New("development")
	app.Router().Register("GET", "/", Page{})

	app.Run()
}
```
