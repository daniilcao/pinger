package action

import (
	"fmt"
	"os/exec"
	"regexp"
	"sync"
	"time"
)

type Ping struct {
	ip []map[string]string
	ipAction string
}

func (p Ping) pingAddr(ip string ) string  {
	out, _ := exec.Command("ping", ip, "-c 2", "-i 3", "-w 5").Output()
	stat := p.regTest(`\b0%`,string(out))

	var status = "non"
	switch stat{
	case "ok":
		status = "Ok"
	default:
		status = "reboot"
	}
	return status
}

func (p Ping) rebootRele(ch <- chan map[string]string,quit <- chan int)  {
	for {
		select {
		case k := <-ch:
			if k["status"] == "reboot"{
				fmt.Println("reboot")
				sw := Switcher{Ip:k["sanOffIp"],Id:k["id"]}
				sw.Off()
				time.Sleep(3 * time.Second)
				sw.On()
				//fmt.Println("reboot Rele")
			}
		case q := <- quit:
			fmt.Printf("exit%v\t",q)
			break
		}
	}
}

func (p *Ping) regTest(pattern,text string) string{
	var res = "non"
	matched, _ := regexp.Match(pattern, []byte(text))
	if matched {res = "ok"}
	return res
}

func (p *Ping) InPoint(){

	ch := make(chan map[string]string)
	quit :=  make(chan int)

	go p.rebootRele(ch,quit)
	wg := &sync.WaitGroup{}
	for _,v := range p.ip{
		wg.Add(1)
		go func(i map[string]string) {
			defer wg.Done()
			dataRe := map[string]string{
				"status":p.pingAddr(i["fermaIp"]),
				"sanOffIp":i["sanOffIp"],
				"id":i["id"],
			}
			ch <- dataRe
		}(v)
	}
	wg.Wait()
	quit <- 0

}

func InPoint(data []map[string]string) {
	p:= Ping{ip:data}
	p.InPoint()
}



