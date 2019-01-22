# gconf

c.json
```json
{
    "ABC":"a",
    "CCC":12,
    "B":["abc","asc","as"],
    "DD":{
        "DDD":23
    }
}
```

```go
package testjson

import "github.com/hiruok/gconf"

func init(){
    err := gconf.Register("conffile","/abc/c.json||./c.json||abc/c.json") //如果/abc/c.json不存在就查找./c.json，./c.json不存在就查找abc/c.json
    if err != nil {
        panic(err)
    }
}


func getConf(){
    conffile,err := gconf.GetConfiger("conffile")
    if err != nil {
        panic(err)
    }
    conffile.String("ABC") // => "a"
    conffile.Int("CCC")  // => 12
    conffile.Strings("B") // => []string{"abc","asc","as"}
    conffile.DefaultString("a","bc") // => "bc"
    conffile.Int("DD.DDD") // => 23
}
```


c.yaml

```yaml
user : domgoer
password : qq123456
sex: 1
married: false
ts:
  tss : dd
tlist :
  - a
  - b
  - c
```

```go
package testyml
import "github.com/hiruok/gconf"

func init(){
    err := gconf.Register("cyaml","c.yaml")
    if err != nil {
            panic(err)
        }
}

func getConfig(){
    cyaml, err := gconf.GetConfiger("cyaml")
    if err != nil {
        panic(err)
    }
    cyaml.String("user") // => domgoer
    cyaml.Int("sex")  // => 1
    cyaml.Bool("married") // => false
    cyaml.String("ts.tss") //=>dd
    cyaml.Strings("tlist") //=>[]string{"a","b","c"}
}
```

```go
package testyml
import "github.com/hiruok/gconf"

type sRead struct{
    User     string                 `yaml:"user"`
    Password string                 `yaml:"password"`
    Sex      int                    `yaml:"sex"`
    Married  bool                   `yaml:"married"`
    Ts       map[string]interface{} `yaml:"ts"`
    Tlist    []string               `yaml:"tlist"`
    Default  int                    `yaml:"default"  default:"12"`
} 

func getConfig(){
    sr := &sRead{}
    // 不需要init
    gconf.Read2Struct("filepath",sr)
    /*
        &sRead{
            User:     "domgoer",
            Password: "qq123456",
            Sex:      1,
            Married:  false,
            Ts:       map[string]interface{}{"tss": "dd"},
            Tlist:    []string{"a", "b", "c"},
            Default:  0,
        }
    */
    gconf.Read2StructByDefault("filepath",sr)
    /*
            &sRead{
                User:     "domgoer",
                Password: "qq123456",
                Sex:      1,
                Married:  false,
                Ts:       map[string]interface{}{"tss": "dd"},
                Tlist:    []string{"a", "b", "c"},
                Default:  1,
            }
        */
}
```
