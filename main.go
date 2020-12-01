package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pinger/action"
	"reflect"
)

func ConvertInterface(data interface{}) []map[string]string  {
	object := reflect.ValueOf(data)
	var resData []map[string]string
	for i := 0; i < object.Len(); i++ {
		interMap := object.Index(i).Elem()
		key := interMap.MapKeys()
		var mapRsd = map[string]string{}
		for _,v := range key{
			k := v.String()
			vv := interMap.MapIndex(v)
			mapRsd[k] = vv.Elem().String()
		}
		resData = append(resData, mapRsd)
	}
	return resData
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

func main() {
	data,err := ReadIp()
	if err != nil{
		return
	}
	clean := ConvertInterface(data)
	action.InPoint(clean)
}