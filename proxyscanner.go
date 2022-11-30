package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

type resMsg struct {
	IP      string `json:"IP"`
	Message string `json:"Message"`
}

var proxyList = []string{}
var WaitDelList = []string{}

type AuthInfo struct {
	Username string
	Password string
}

func init() {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[Proxy Url]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	var w = &sync.WaitGroup{}
	var enterProxy string
	fmt.Print("Enter Proxy: ")
	_, e := fmt.Scan(&enterProxy)
	if e != nil {
		fmt.Println("Enter Err")
		return

	} else {
		fmt.Println(enterProxy)
		username := strings.Split(enterProxy, ":")[2]
		password := strings.Split(enterProxy, ":")[3]
		authInfo := AuthInfo{}
		authInfo.Username = username
		authInfo.Password = password
		host := strings.Split(enterProxy, ":")[0]
		newHost := strings.Join(strings.Split(host, ".")[:3], ".")
		fmt.Println("Generate Done! Start Test!")
		fmt.Println(proxyList)

		//var cc = []string{"nihao","nibuhao"}
		//dd := "nihao"
		//fmt.Println(cc)
		//cc = DelFromSlice(cc,dd)
		//fmt.Println(cc)
		for hhi := 1; hhi < 256; hhi++ {
			for c := 63350; c < 65532; c++ {
				//DelFromSlice(&proxyList, proxyList[c])
				//fmt.Println(proxyList, newHost+"."+strconv.Itoa(hhi)+"."+strconv.Itoa(c))
				//fmt.Println(ppo)
				w.Add(1)
				go func(newHost string, hhi int, c int, authinfo AuthInfo) {
					var resMsg = resMsg{}
					fmt.Println("YES")
					var userInfo = url.UserPassword(authInfo.Username, authInfo.Password)
					dd, _ := grequests.Get("http://api.logalerts.services/v1/test/addr", &grequests.RequestOptions{Proxies: map[string]*url.URL{"http": {Host: newHost + "." + strconv.Itoa(hhi) + ":" + strconv.Itoa(c), User: userInfo}}, RequestTimeout: 350000000})
					json.Unmarshal([]byte(dd.String()), &resMsg)
					fmt.Println(dd.String())
					fmt.Println(newHost + "." + strconv.Itoa(hhi) + ":" + strconv.Itoa(c))
					if resMsg.Message == "Success" {
						log.Println(newHost + "." + strconv.Itoa(hhi) + ":" + strconv.Itoa(c)+":"+authinfo.Username+":"+authinfo.Password)
					}
					w.Done()
				}(newHost, hhi, c, authInfo)
				w.Wait()

			}
		}

	}

}

func TestProxies() {

}
func DelFromSlice(slices *[]string, k string) {
	ss := *slices
	for s := 0; s < len(*slices); s++ {
		if ss[s] == k {
			*slices = append(ss[:s], ss[s+1:]...)
		}
	}
}