// Date: 2018/8/22 下午3:19
//
// Description:
package gconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

type jsonConfig struct {
}

func (jc *jsonConfig) parse(filePath string) (Configer, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return jc.parseData(content)
}

func (jc *jsonConfig) parseData(data []byte) (Configer, error) {
	x := &jsonConfigContainer{
		data: make(map[string]interface{}),
	}
	err := json.Unmarshal(data, &x.data)
	if err != nil {
		var wrappingArray []interface{}
		err2 := json.Unmarshal(data, &wrappingArray)
		if err2 != nil {
			return nil, err
		}
		x.data["rootArray"] = wrappingArray
	}

	x.data = expandValueEnvForMap(x.data)

	return x, nil
}

type jsonConfigContainer struct {
	data map[string]interface{}
	sync.RWMutex
}

func (jcr *jsonConfigContainer) String(key string) string {
	val := jcr.getData(key)
	if val != nil {
		if v, ok := val.(string); ok {
			return v
		}
	}
	return ""
}

func (jcr *jsonConfigContainer) Strings(key string) (res []string) {
	val := jcr.getData(key)
	if val != nil {
		if vv, ok := val.([]interface{}); ok {
			for _, v := range vv {
				if str, strOK := v.(string); strOK {
					res = append(res, str)
				}
			}
		}
	}
	return
}

func (jcr *jsonConfigContainer) Int(key string) (int, error) {
	val := jcr.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return int(v), nil
		}
		return 0, errors.New("not int value")
	}
	return 0, errors.New("not exist key:" + key)
}

func (jcr *jsonConfigContainer) Int64(key string) (int64, error) {
	val := jcr.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return int64(v), nil
		}
		return 0, errors.New("not int64 value")
	}
	return 0, errors.New("not exist key:" + key)
}

func (jcr *jsonConfigContainer) Bool(key string) (bool, error) {
	val := jcr.getData(key)
	if val != nil {
		return parseBool(val)
	}
	return false, fmt.Errorf("not exist key: %q", key)
}

func (jcr *jsonConfigContainer) Float(key string) (float64, error) {
	val := jcr.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return v, nil
		}
		return 0.0, errors.New("not float64 value")
	}
	return 0.0, errors.New("not exist key:" + key)
}

func (jcr *jsonConfigContainer) DefaultString(key string, defaultval string) string {
	if v := jcr.String(key); v != "" {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) DefaultStrings(key string, defaultval []string) []string {
	if v := jcr.Strings(key); v != nil {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) DefaultInt(key string, defaultval int) int {
	if v, err := jcr.Int(key); err == nil {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) DefaultInt64(key string, defaultval int64) int64 {
	if v, err := jcr.Int64(key); err == nil {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) DefaultBool(key string, defaultval bool) bool {
	if v, err := jcr.Bool(key); err == nil {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) DefaultFloat(key string, defaultval float64) float64 {
	if v, err := jcr.Float(key); err == nil {
		return v
	}
	return defaultval
}

func (jcr *jsonConfigContainer) Interface(key string) (interface{}, error) {
	val := jcr.getData(key)
	if val != nil {
		return val, nil
	}
	return nil, errors.New("not exist key")
}

// section.key or key
func (jcr *jsonConfigContainer) getData(key string) interface{} {
	if len(key) == 0 {
		return nil
	}

	jcr.RLock()
	defer jcr.RUnlock()

	sectionKeys := strings.Split(key, ".")
	if len(sectionKeys) >= 2 {
		curValue, ok := jcr.data[sectionKeys[0]]
		if !ok {
			return nil
		}
		for _, key := range sectionKeys[1:] {
			if v, ok := curValue.(map[string]interface{}); ok {
				if curValue, ok = v[key]; !ok {
					return nil
				}
			}
		}
		return curValue
	}
	if v, ok := jcr.data[key]; ok {
		return v
	}
	return nil
}

func newJsonConfigContainer(filePath string) (Configer, error) {
	j := &jsonConfig{}
	return j.parse(filePath)
}
