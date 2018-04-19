package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	configArg   = "config:"
	CurrentPath = ""
)

func getConfigName() (path string, err error) {
	path = ""
	fmt.Println(CurrentPath)
	dirname := ""
	if len(CurrentPath) != 0 {
		dirname = CurrentPath
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		dirname, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return path, err
		}
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

func readStringFromConfigFile() (jsonStr []byte, err error) {

	filename, e := getConfigName()
	if len(filename) == 0 {
		err = e
		return
	}

	file, e := os.Open(filename)
	if e != nil {
		err = e
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	b := buf.Bytes()
	//читаем файл
	file.Read(b)
	jsonStr = b
	return
}

func getConfigFromPath(configStr []byte, pathName string, configuration interface{}) (err error) {
	paths := strings.Split(pathName, ".")
	// создаем срез для параметров json объекта верхнего уровня
	c := make(map[string]interface{})

	// unmarschal JSON
	json.Unmarshal(configStr, &c)
	for s, value := range c {
		if s == paths[0] {
			jsonString, _ := json.Marshal(value)
			if len(paths) == 1 {
				err = json.Unmarshal(jsonString, &configuration)
			} else {
				getConfigFromPath(jsonString, strings.Join(paths[1:], "."), &configuration)
			}
			return
		}
	}
	return
}

//получение текущей конфигурации для всего проекта
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

//получение конфигурации, находящейся по указанному в pathName пути
func GetConfigPath(pathName string, configuration interface{}) (err error) {
	configStr, _ := readStringFromConfigFile()
	return getConfigFromPath(configStr, pathName, &configuration)
}
