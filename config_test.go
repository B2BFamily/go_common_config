package config

import "testing"

func TestGetConfig_Success(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: Success,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "", &configuration)
	if err != nil {
		t.Error("GetConfig error on exec", err)
	}
	if configuration.Path1.Path2.Bool != referenceValue.Bool || configuration.Path1.Path2.Int != referenceValue.Int || configuration.Path1.Path2.String != referenceValue.String {
		t.Error("GetConfig error on reading data")
	}
}

func TestGetConfig_ErrorStruct(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: ErrorPath,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "", &configuration)
	if err != nil {
		t.Error("GetConfig error on exec", err)
	}
	if configuration.Path1.Path2.Bool != zeroValue.Bool || configuration.Path1.Path2.Int != zeroValue.Int || configuration.Path1.Path2.String != zeroValue.String {
		t.Error("GetConfig error on reading data")
	}
}

func TestGetConfig_ErrorPath(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: ErrorPath,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "path1.path2", &configuration)
	if err == nil {
		t.Error("GetConfig don't say about error 'not found path'")
	}
}

func TestGetConfig_ErrorInt(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: ErrorInt,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "", &configuration)
	if err == nil {
		t.Error("cannot unmarshal {type} into Go struct field {field} of type {type}")
	}
}
func TestGetConfig_ErrorString(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: ErrorString,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "", &configuration)
	if err == nil {
		t.Error("cannot unmarshal {type} into Go struct field {field} of type {type}")
	}
}
func TestGetConfig_ErrorBool(t *testing.T) {
	base := fakeConfigurator{
		ErrorType: ErrorBool,
	}
	configuration := fakeFullStruct{}
	jsonStr, _ := base.readStringFromConfigFile()
	err := getConfigFromPath(jsonStr, "", &configuration)
	if err == nil {
		t.Error("cannot unmarshal {type} into Go struct field {field} of type {type}")
	}
}
