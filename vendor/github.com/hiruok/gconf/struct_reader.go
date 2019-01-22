// Author : dmc
//
// Date: 2018/9/4 上午10:27
//
// Description:
package gconf

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

const (
	str          = "string"
	strs         = "[]string"
	it           = "int"
	it64         = "int64"
	flt          = "float"
	bol          = "bool"
	defaultValue = "default"
)

type read2Struct struct {
	rv reflect.Value
}

// Read2Struct 将配置信息读取到结构体,不使用默认值
func Read2Struct(path string, out interface{}) error {
	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invaild received param , need ptr")
	}
	return read(path, out)
}

// Read2StructByDefault 将配置信息读取到结构体,使用默认值
func Read2StructByDefault(path string, out interface{}) error {
	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invaild received param , need ptr")
	}
	err := read(path, out)
	if err != nil {
		return err
	}
	rs := read2Struct{
		rv: rv,
	}
	return rs.value(out)
}

func read(path string, out interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	f := strings.Split(path, ".")
	switch f[len(f)-1] {
	case "json":
		err = json.Unmarshal(b, out)
	case "yml", "yaml":
		err = yaml.Unmarshal(b, out)
	default:
		return fmt.Errorf("file type error , only support json/yaml/yml")
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *read2Struct) value(out interface{}) error {
	rv := r.rv
	te := rv.Type().Elem()
	for i := 0; i < te.NumField(); i++ {
		var err error
		field := te.Field(i)
		dv := field.Tag.Get(defaultValue)
		if dv == "" {
			continue
		}
		rf := rv.Elem().Field(i)
		err = r.assign(field, dv, rf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *read2Struct) assign(key reflect.StructField, dv string, rv reflect.Value) error {
	switch rv.Type().String() {
	case str:
		if rv.String() == "" {
			rv.SetString(dv)
		}
	case it, it64:
		if rv.Int() == 0 {
			i, err := strconv.ParseInt(dv, 10, 64)
			if err != nil {
				return nil
			}
			rv.SetInt(i)
		}
	case flt:
		if rv.Float() == 0 {
			f, err := strconv.ParseFloat(dv, 64)
			if err != nil {
				return nil
			}
			rv.SetFloat(f)
		}
	case bol:
		if dv == "true" {
			rv.SetBool(true)
		}
	default:
		return fmt.Errorf("complex types cannot use initial values")
	}
	return nil
}
