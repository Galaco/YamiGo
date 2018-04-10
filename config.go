package golaco

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type configuration struct {
	App struct {
		Url string
		Port int
	}
	Views struct {
		TemplateDir string
	}
}

var Configuration configuration

func (this *configuration) parse(environment string) {
	filename := "config."
	if environment != "" {
		filename += environment + "."
	}
	filename += "yml"

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(data, this)
}