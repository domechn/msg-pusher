//Date: 2018/8/22 下午3:19
//
//Description:
package gconf

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strings"
	"time"
)

type Configer interface {
	String(key string) string
	Strings(key string) []string
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	DefaultString(key string, defaultVal string) string
	DefaultStrings(key string, defaultVal []string) []string
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64
	Interface(key string) (interface{}, error)
}

type Config interface {
	parse(key string) (Configer, error)
	parseData(data []byte) (Configer, error)
}

var configers = make(map[string]Configer)

func Register(alias, filePath string) error {
	var err error
	filePath, err = getOnePath(filePath)
	if err != nil {
		return err
	}
	if _, ok := configers[alias]; ok {
		return errors.New("already has config named " + alias)
	}
	f := strings.Split(filePath, ".")
	switch f[len(f)-1] {
	case "json":
		jcr, err := newJsonConfigContainer(filePath)
		if err != nil {
			return err
		}
		configers[alias] = jcr
	case "yaml", "yml":
		ycr, err := newYamlConfigContainer(filePath)
		if err != nil {
			return err
		}
		configers[alias] = ycr
	default:
		return errors.New("file type error , only support json/yaml/yml")
	}
	return nil
}

func GetConfiger(alias string) (Configer, error) {
	if _, ok := configers[alias]; !ok {
		return nil, errors.New("no config named " + alias)
	}
	return configers[alias], nil
}

func getOnePath(filePath string) (string, error) {
	filePaths := strings.Split(filePath, "||")
	for _, v := range filePaths {
		_, err := os.Stat(v)
		if err == nil {
			return v, nil
		}
	}
	return "", errors.New("file not exist")
}

func expandValueEnvForMap(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		switch value := v.(type) {
		case string:
			m[k] = expandValueEnv(value)
		case map[string]interface{}:
			m[k] = expandValueEnvForMap(value)
		case map[string]string:
			for k2, v2 := range value {
				value[k2] = expandValueEnv(v2)
			}
			m[k] = value
		}
	}
	return m
}

func expandValueEnv(value string) (realValue string) {
	realValue = value

	vLen := len(value)
	// 3 = ${}
	if vLen < 3 {
		return
	}
	// Need start with "${" and end with "}", then return.
	if value[0] != '$' || value[1] != '{' || value[vLen-1] != '}' {
		return
	}

	key := ""
	defaultV := ""
	// value start with "${"
	for i := 2; i < vLen; i++ {
		if value[i] == '|' && (i+1 < vLen && value[i+1] == '|') {
			key = value[2:i]
			defaultV = value[i+2 : vLen-1] // other string is default value.
			break
		} else if value[i] == '}' {
			key = value[2:i]
			break
		}
	}

	realValue = os.Getenv(key)
	if realValue == "" {
		realValue = defaultV
	}

	return
}

// ParseBool returns the boolean value represented by the string.
//
// It accepts 1, 1.0, t, T, TRUE, true, True, YES, yes, Yes,Y, y, ON, on, On,
// 0, 0.0, f, F, FALSE, false, False, NO, no, No, N,n, OFF, off, Off.
// Any other value returns an error.
func parseBool(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
				return true, nil
			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
				return false, nil
			}
		case int8, int32, int64:
			strV := fmt.Sprintf("%d", v)
			if strV == "1" {
				return true, nil
			} else if strV == "0" {
				return false, nil
			}
		case float64:
			if v == 1.0 {
				return true, nil
			} else if v == 0.0 {
				return false, nil
			}
		}
		return false, fmt.Errorf("parsing %q: invalid syntax", val)
	}
	return false, fmt.Errorf("parsing <nil>: invalid syntax")
}

// ToString converts values of any type to string.
func toString(x interface{}) string {
	switch y := x.(type) {

	// Handle dates with special logic
	// This needs to come above the fmt.Stringer
	// test since time.Time's have a .String()
	// method
	case time.Time:
		return y.Format("A Monday")

		// Handle type string
	case string:
		return y

		// Handle type with .String() method
	case fmt.Stringer:
		return y.String()

		// Handle type with .Error() method
	case error:
		return y.Error()

	}

	// Handle named string type
	if v := reflect.ValueOf(x); v.Kind() == reflect.String {
		return v.String()
	}

	// Fallback to fmt package for anything else like numeric types
	return fmt.Sprint(x)
}
