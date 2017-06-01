package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

// func main() {
//     response, _ := http.Get("http://www.baidu.com")
//     defer response.Body.Close()
//     body, _ := ioutil.ReadAll(response.Body)
//     fmt.Println(string(body))
// }

func main() {
    client := &http.Client{}
    reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

    reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
    reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
    reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
    reqest.Header.Set("Cache-Control","max-age=0")
    reqest.Header.Set("Connection","keep-alive")

    response, _ := client.Do(reqest)
    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
        bodystr := string(body);
        fmt.Println(bodystr)
    }
}