package yamigo

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type configuration struct {
	App struct {
		Url string `yaml:"url"`
		Port int `yaml:"port"`
	}
	Views struct {
		Template struct {
			BaseDir string `yaml:"baseDir"`
			BaseTemplate string `yaml:"baseTemplate"`
		}
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