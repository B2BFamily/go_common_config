package config

type fakeFullStruct struct {
	Path1 struct {
		Path2 fakeInnerStruct `json:"path2"`
	} `json:"path1"`
}

type fakeInnerStruct struct {
	Int    int    `json:"int"`
	String string `json:"string"`
	Bool   bool   `json:"bool"`
}

var referenceValue = fakeInnerStruct{
	1,
	"test",
	true,
}
var zeroValue = fakeInnerStruct{}

const (
	Success     = iota
	ErrorPath   = iota
	ErrorInt    = iota
	ErrorString = iota
	ErrorBool   = iota
)

type fakeConfigurator struct {
	ErrorType int
}

func (base *fakeConfigurator) readStringFromConfigFile() (jsonStr []byte, err error) {
	switch base.ErrorType {
	case Success:
		return []byte(`{"path1":{"path2":{"int":1,"string":"test","bool":true}}}`), nil
	case ErrorPath:
		return []byte(`{"path2":{"path2":{"int":1,"string":"test","bool":true}}}`), nil
	case ErrorInt:
		return []byte(`{"path1":{"path2":{"int":"NO","string":"test","bool":true}}}`), nil
	case ErrorString:
		return []byte(`{"path1":{"path2":{"int":1,"string":1,"bool":true}}}`), nil
	case ErrorBool:
		return []byte(`{"path1":{"path2":{"int":1,"string":"test","bool":1}}}`), nil
	}
	return []byte(""), nil
}
