package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"

	"github.com/axelrindle/nc-cfg-gen/cli"
	"github.com/axelrindle/nc-cfg-gen/nextcloud"
	"github.com/tzfqh/gmdtable"
)

const docUrl = "https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html"

func documentationUrl(key string) string {
	if key == "" {
		return ""
	}

	return fmt.Sprintf("[`%s`](%s#%s)", key, docUrl, key)
}

func makeConfigTable() string {
	instance := nextcloud.ConfigDynamic{}

	v := reflect.ValueOf(instance)
	t := v.Type()

	variables := []map[string]any{}

	for i := range t.NumField() {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}

		variables = append(variables, map[string]any{
			"Environment Variable": fmt.Sprintf("`%s`", field.Tag.Get("env")),
			"Description":          field.Tag.Get("envDesc"),
			"Default":              field.Tag.Get("envDefault"),
			"Documentation":        documentationUrl(field.Tag.Get("doc")),
		})
	}

	headers := []string{"Environment Variable", "Description", "Default", "Documentation"}
	table, err := gmdtable.Convert(headers, variables)
	if err != nil {
		log.Fatal(err)
	}

	return table
}

func main() {
	tpl, err := template.ParseFiles("README.tpl")
	if err != nil {
		log.Fatal(err)
	}

	target, err := os.OpenFile("README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(target, map[string]string{
		"Usage":         cli.Usage(),
		"Configuration": makeConfigTable(),
	})
	if err != nil {
		log.Fatal(err)
	}
}
