package folder1

import ("fmt"
	"encoding/json"
	"os"
	"io/ioutil"
)

type shipObject struct {
	Name string
	Source string
}


func (p shipObject) toString() string {
	return toJson(p)

}


func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}


func Json_File_Read1() {
	file1, e := ioutil.ReadFile("/Users/shuvamoymondal/test.json")
	if e != nil {
		fmt.Println("File error %v\n", e)
		os.Exit(1)
	}


	var c []shipObject
	json.Unmarshal(file1, &c)
	var pages =c

	for _,p := range pages {
		fmt.Println(p)
		fmt.Println(p.toString())
	}
}