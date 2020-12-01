package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BodyJson struct {
	Deviceid string `json:"deviceid"`
	Data map[string]string `json:"data"`
}

type Switcher struct {
	Ip string
	Id string
}

func (sw *Switcher)On() {
	b := BodyJson{
		Deviceid: sw.Id,
		Data: map[string]string{"switch":"on"},
	}

	js, _ := json.Marshal(b)
	r := bytes.NewReader(js)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post("http://"+sw.Ip+":8081/zeroconf/switch", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}

func (sw *Switcher) Off() {
	b := BodyJson{
		Deviceid: sw.Id,
		Data: map[string]string{"switch":"off"},
	}

	js, _ := json.Marshal(b)
	r := bytes.NewReader(js)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post("http://"+sw.Ip+":8081/zeroconf/switch", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}