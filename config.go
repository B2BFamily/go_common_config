//Библиотека для работы с конфигами
//Конфиги должны хранится в json файлах в папке config рядом с запускаемым файлом и называться config.{mode.}json
//{mode} - атрибут для выбора конфиг файла, указывается при запуске программы при помощи флага config:{mode}
//к примеру
//	./main.exe config:dev
//для конфига будет браться файл /config/config.dev.json
//если флаг отстутствует, то берется config.json
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	//Путь к исполняемой программе, необходим для отладки
	CurrentPath = ""
	configArg   = "config:"
)

func getConfigName() (path string, err error) {
	path = ""
	dirname := ""
	if len(CurrentPath) != 0 {
		dirname = CurrentPath
	} else {
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

	path = filepath.Join(dirname, "config", "config."+mode+"json")
	fmt.Println(dirname)
	fmt.Println(path)
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
	// создаем срез для параметров json объекта верхнего уровня
	obj := make(map[string]interface{})

	// unmarschal JSON
	err = json.Unmarshal(configStr, &obj)
	if err != nil {
		return err
	}

	if len(pathName) != 0 {
		paths := strings.Split(pathName, ".")
		for _, path := range paths {
			if obj[path] != nil {
				obj = obj[path].(map[string]interface{})
			} else {
				return errors.New("path not found")
			}

		}
	}

	jsonString, _ := json.Marshal(obj)
	err = json.Unmarshal(jsonString, &configuration)
	return err
}

//получение текущей конфигурации для всего проекта
func GetConfig(configuration interface{}) (err error) {
	configStr, _ := readStringFromConfigFile()
	return getConfigFromPath(configStr, "", &configuration)
}

//получение конфигурации, находящейся по указанному в pathName пути
func GetConfigPath(pathName string, configuration interface{}) (err error) {
	configStr, _ := readStringFromConfigFile()
	return getConfigFromPath(configStr, pathName, &configuration)
}
