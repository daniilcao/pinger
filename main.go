package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func ConvertInterface(data interface{}) []map[string]string  {
	object := reflect.ValueOf(data)
	
	for i := 0; i < object.Len(); i++ {
		fmt.Println(object.Index(i))
	}
	return []map[string]string{}
}

func ReadIp()  (interface{},error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plan, err := ioutil.ReadFile(pwd+"/ip.json")
	if err != nil{
		fmt.Println("Reade File",err)
		return nil,err
	}

	var ip interface{}
	err = json.Unmarshal(plan, &ip)
	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
		return nil,err
	}
	return ip,nil

}

func main()  {
	data,err := ReadIp()
	if err != nil{
		return
	}
	clean := ConvertInterface(data)
	fmt.Println(clean)

}