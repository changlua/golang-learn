package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

//对应请求的响应体
type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description struct {
			Source string `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

func Test_Project2_cidian(t *testing.T) {
	for {
		fmt.Println("请输入要查询的词：")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') //读取一行
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		word := strings.TrimSpace(input)
		if word == "q" {
			fmt.Println("成功退出！")
			break
		}
		//查询
		query (word)
	}
}

//动态查询
func query(word string)  {
	client := &http.Client{}
	//构造请求体
	request := DictRequest{
		TransType: "en2zh",
		Source: word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	//var data = strings.NewReader(`{"trans_type":"en2zh","source": word }`)
	var data = bytes.NewReader(buf)
	//1、创建请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	//2、设置请求头
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//3、发起响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//4、读取数据
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	//将对应的json字符串转为结构体对象
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)   //json => 对象
	if err != nil {
		log.Fatal(err)   //日志打印
	}
	//fmt.Printf("%#v\n", dictResponse)
	//遍历其中的结果集
	for index , item := range dictResponse.Dictionary.Explanations {  //下标：结果集
		fmt.Println(index , " ", item)
	}
}