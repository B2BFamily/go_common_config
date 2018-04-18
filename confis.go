package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var configArg = "config:"

func getConfigName() (path string, err error) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	dirname := ""
	dirname, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	var mode = ""
	for _, arg := range os.Args {
		if strings.Contains(arg, configArg) {
			mode = strings.Replace(arg, configArg, "", -1) + "."
		}
	}

	path = filepath.Dir(dirname) + "\\config\\config." + mode + "json"
	return
}

func GetConfig(configuration interface{}) (err error) {
	filename, _ := getConfigName()
	if len(filename) == 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return
	}

	return
}

func GetConfigPath(pathName string, configuration interface{}) (err error) {
	filename, _ := getConfigName()
	if len(filename) == 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	b := buf.Bytes()
	//читаем файл
	file.Read(b)

	// создаем срез для параметров json объекта верхнего уровня
	c := make(map[string]interface{})

	// unmarschal JSON
	json.Unmarshal(b, &c)
	for s, value := range c {
		if s == pathName {
			jsonString, _ := json.Marshal(value)
			err = json.Unmarshal(jsonString, &configuration)
			return
		}
	}
	return
}
