package main

import (
	"os"
	"text/template"
)

type secretObject struct {
	NameSpace string
	SecretName string
	SecretData map[interface{}]interface {}
}

func SecretMap() error {
	var secretAddr map[interface{}]interface{}
	var secret secretObject

	mapYamlFile := "secret.yaml"

	err := GetMap(mapYamlFile, &secretAddr)
	if err != nil {
		return err
	}

	secret.NameSpace = secretAddr["name_space"].(string)
	secret.SecretName = secretAddr["secret_name"].(string)
	secret.SecretData = secretAddr["secret_data"].(map[interface{}]interface {})

	secretTemplate := "secret-template.yaml"
	t := template.New("secret")
	t, err = t.ParseFiles(secretTemplate)
	if(err != nil) {
		return err
	}

	t.ExecuteTemplate(os.Stdout, secretTemplate, secret)

	return nil
}
