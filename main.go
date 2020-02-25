package main

import (
	"Harmes/Curl"
	"Harmes/Email"
	"encoding/json"
	"time"
)

type Asset struct {
	Url string `json:"url"`
}

type Good struct {
	Sku    string  `json:sku`
	Title  string  `json:"title"`
	Assets []Asset `json:"assets"`
	Url    string  `json:"url"`
}

type Goods struct {
	Total    int    `json:total`
	Products []Good `json:products`
}

var lastSearch = make(map[string]Good)
var newGoods = make(map[string]Good)
var url = "https://bck.hermes.cn/product?locale=cn_zh&category=WOMENBAGSBAGSCLUTCHES&sort=relevance"

var EmailParam = Email.EmailParam{
	"smtp.163.com",
	25,
	"suyuhao1994@163.com",
	"Suyuhao940225",
	"suyuhao1994@163.com",
	"644208937@qq.com",
}

func main() {

	list := Curl.GetGoodsList(url)
	Email.InitEmail(&EmailParam)
	var goodsList Goods
	json.Unmarshal([]byte(list), &goodsList)

	//欢迎订阅
	body := "欢迎订阅HERMES SUPER UPDATER V1,官网推新我们会及时推送的哟"
	for _, v := range goodsList.Products {
		lastSearch[v.Sku] = v
	}
	Email.SendEmail("欢迎订阅HERMES SUPER UPDATER V1", body)
	for {
		checkIsNew()
		time.Sleep(30 * time.Second)
	}
}

func getLatest() Goods {
	list := Curl.GetGoodsList(url)
	var goodsList Goods
	json.Unmarshal([]byte(list), &goodsList)

	return goodsList
}

func checkIsNew() {
	goodsList := getLatest()
	for _, v := range goodsList.Products {
		_, ok := lastSearch[v.Sku]

		if !ok {
			Email.SendEmail("HERMES SUPER上新", "")
			break
		}
	}
	lastSearch = make(map[string]Good)
	for _, v := range goodsList.Products {
		lastSearch[v.Sku] = v
	}
}
