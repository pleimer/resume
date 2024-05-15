package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"sigs.k8s.io/yaml"
)

type ResumeVars struct {
	Address string
	Phone   string
	Zipcode string
}

func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

func main() {

	args := os.Args[1:]

	configFile := "vars.public.json"

	if len(args) >= 2 {

		if args[0] != "--config" && args[0] != "-c" {
			fmt.Println("Usage: program_name --config <config_file>")
			fmt.Println("       program_name -c <config_file>")

			return
		} else {
			configFile = args[1]
		}
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		panic(Wrap(err, "reading vars file"))
	}

	var vars ResumeVars

	err = json.Unmarshal(data, &vars)
	if err != nil {
		panic(Wrap(err, "parsing vars file"))
	}

	resumeTemplate, err := os.ReadFile("resume.template.yaml")
	if err != nil {
		panic(Wrap(err, "reading template file"))
	}

	tmpl, err := template.New("resume").Parse(string(resumeTemplate))
	if err != nil {
		panic(Wrap(err, "parsing resume template file"))
	}

	// template populated
	var final bytes.Buffer
	if err := tmpl.Execute(&final, vars); err != nil {
		panic(Wrap(err, "applying vars to template"))
	}

	// serialize to JSON file

	resJson, err := yaml.YAMLToJSON(final.Bytes())
	if err != nil {
		panic(Wrap(err, "converting YAML to JSON"))
	}

	file, err := os.Create("resume.json")
	if err != nil {
		panic(Wrap(err, "opening resume file"))
	}
	defer file.Close()

	if _, err := file.Write(resJson); err != nil {
		panic(Wrap(err, "writing JSON to file"))
	}

	fmt.Println("rendered resume template successfullly")

}
