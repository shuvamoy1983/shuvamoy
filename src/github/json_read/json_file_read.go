package json_read

import ("fmt"
	"encoding/json"
         "os"
         "io/ioutil"
)

type jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	Databases []DatabaseType
}

type DatabaseType struct {
	Host string
	User string
	Pass string
	Type string
	Name string
	Tables []TablesType

}
type TablesType struct {
	Name string
	Statement string
	Regex string
	Types []TypesType
}

type TypesType struct {
	Id int
	Value string
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func (p jsonobject) toString() string {
	return toJson(p)
}


func Json_File_Read() {
	file, e := ioutil.ReadFile("/Users/shuvamoymondal/config.json")
	if e !=nil {
		fmt.Println("File error %v\n",e)
		os.Exit(1)
	}

var c []jsonobject
	json.Unmarshal(file, &c)
	var pages =c

	for _,p := range pages {
		fmt.Println(p.toString())
	}
}