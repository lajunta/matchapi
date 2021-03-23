package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/jaswdr/faker"
)

type news struct {
	Title       string `json:"title"`       //标题
	Category    string `json:"category"`    //所属分类
	Date        string `json:"date"`        //日期
	Description string `json:"description"` //摘要描述
	Content     string `json:"content"`     //新闻内容
	Url         string `json:"url"`         //新闻url
}

type solution struct {
	image       string //图片url
	title       string //标题
	description string //摘要
}

type security struct {
	title    string    //标题
	descript string    //摘要
	rank     int       //级别
	date     time.Time //时间
}

type banner struct {
	Name        string `json:"name"`        //标题
	Image       string `json:"image"`       //图片url
	Alt         string `json:"alt"`         //图片替换文本
	Description string `json:"description"` //描述
}

var nes []news
var solus []solution
var secs []security
var banners []banner

func getNews(w http.ResponseWriter, r *http.Request) {
	num := 10
	faker := faker.New()
	var ns news
	for i := 1; i < num; i++ {
		ns.Title = faker.Lorem().Sentence(6)
		ns.Category = faker.Lorem().Text(1)
		ns.Content = faker.Lorem().Paragraph(2)
		fpath := fmt.Sprintf("assets/news/%d.html", i)
		ioutil.WriteFile(fpath, []byte(ns.Content), 0755)
		ns.Date = time.Now().Format("2006-01-02")
		ns.Description = faker.Lorem().Sentence(30)
		ns.Url = fpath
		nes = append(nes, ns)
	}
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(nes)
}

func getSolutions(w http.ResponseWriter, r *http.Request) {
	// ns := news{title: "dfads", category: "unknown", date: time.Now(), description: "dsfasd", content: "dsfdsaf", url: "dsfasdf"}
	// nes = append(nes, ns)
	w.Header().Add("content-type", "application/json")
	fmt.Fprint(w, nes)
}

func getSecurity(w http.ResponseWriter, r *http.Request) {

}

func getBanners(w http.ResponseWriter, r *http.Request) {
	var banners []banner
	for i := 1; i <= 3; i++ {
		imageName := fmt.Sprintf("image%d", i)
		imageURL := fmt.Sprintf("/assets/images/%d.jpg", i)
		alt := fmt.Sprintf("image text %d", i)
		bs := banner{Name: imageName, Image: imageURL, Alt: alt, Description: "This is first image"}
		banners = append(banners, bs)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banners)
}

// make news images
func makeData() {
	num := 10
	url := "https://picsum.photos/1200/600"
	faker := faker.New()
	var content string
	for i := 1; i < num; i++ {
		content = faker.Lorem().Paragraph(2)
		newspath := fmt.Sprintf("assets/news/%d.html", i)
		ioutil.WriteFile(newspath, []byte(content), 0755)

	}
	for i := 1; i <= 3; i++ {
		response, err := http.Get(url)
		if err != nil {
			println("get photo error")
		}
		defer response.Body.Close()
		imgpath := fmt.Sprintf("assets/images/%d.jpg", i)
		destFile, _ := os.Create(imgpath)
		defer destFile.Close()
		io.Copy(destFile, response.Body)
	}

}

func parseFlag() {
	add := flag.Bool("a", false, "add 10 news")
	flag.Parse()
	if *add {
		makeData()
	}
}

func main() {
	parseFlag()
	fs := http.FileServer(http.Dir("assets/"))
	http.HandleFunc("/wp_json/cyberbuf/news", getNews)
	http.HandleFunc("/wp_json/cyberbuf/get_solutions", getNews)
	http.HandleFunc("/wp_json/cyberbuf/get_security", getSecurity)
	http.HandleFunc("/wp_json/cyberbuf/get_banner", getBanners)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8080", nil)
}
