package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jaswdr/faker"
)

var (
	port string
	add  bool
)

type news struct {
	Title       string `json:"title"`       //标题
	Category    string `json:"category"`    //所属分类
	Image       string `json:"image"`       //图片url
	Date        string `json:"date"`        //日期
	Description string `json:"description"` //摘要描述
	Content     string `json:"content"`     //新闻内容
	Url         string `json:"url"`         //新闻url
}

type solution struct {
	Image       string `json:"image"`       //图片url
	Title       string `json:"title"`       //标题
	Description string `json:"description"` //摘要
}

type security struct {
	Title       string `json:"title"`       //标题
	Description string `json:"description"` //摘要
	Rank        int    `json:"rank"`        //级别
	Date        string `json:"date"`        //时间
}

type banner struct {
	Name        string `json:"name"`        //标题
	Image       string `json:"image"`       //图片url
	Alt         string `json:"alt"`         //图片替换文本
	Description string `json:"description"` //描述
}

type faq struct {
	Question string `json:"question"` //问题
	Answer   string `json:"answer"`   //回答
}

type video struct {
	Name        string `json:"name"`        //标题
	Video       string `json:"video"`       //视频url
	Description string `json:"description"` //描述
}

func getNews(w http.ResponseWriter, r *http.Request) {
	num := 10
	faker := faker.New()
	var ns news
	var nes []news

	for i := 1; i < num; i++ {
		ns.Title = faker.Lorem().Sentence(6)
		ns.Category = faker.Lorem().Word()
		ns.Content = faker.Lorem().Paragraph(2)
		fpath := fmt.Sprintf("%s/assets/news/%d.html", apiurl(), i)
		os.WriteFile(fpath, []byte(ns.Content), 0755)
		ns.Date = time.Now().Format("2006-01-02")
		ns.Description = faker.Lorem().Sentence(30)
		imgPath := fmt.Sprintf("%s/assets/images/%d.jpg", apiurl(), i)
		ns.Image = imgPath
		ns.Url = fpath
		nes = append(nes, ns)
	}
	w.Header().Add("content-type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(nes)
}

func getFAQ(w http.ResponseWriter, r *http.Request) {
	num := 10
	faker := faker.New()
	var faqs []faq
	var fq faq
	for i := 1; i < num; i++ {
		fq.Question = faker.Lorem().Sentence(6)
		fq.Answer = faker.Lorem().Sentence(60)
		faqs = append(faqs, fq)
	}

	w.Header().Add("content-type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(faqs)
}

func getSolutions(w http.ResponseWriter, r *http.Request) {
	num := 10
	faker := faker.New()
	var solus []solution
	var solu solution
	for i := 1; i < num; i++ {
		solu.Title = faker.Lorem().Sentence(6)
		solu.Description = faker.Lorem().Sentence(30)
		solu.Image = fmt.Sprintf("%s/assets/images/%d.jpg", apiurl(), i)
		solus = append(solus, solu)
	}

	w.Header().Add("content-type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(solus)
}

func getSecurity(w http.ResponseWriter, r *http.Request) {
	num := 10
	faker := faker.New()
	var secs []security
	var sec security
	for i := 1; i < num; i++ {
		sec.Title = faker.Lorem().Sentence(6)
		sec.Description = faker.Lorem().Sentence(30)
		sec.Rank = rand.Intn(3) + 1
		sec.Date = time.Now().Format("2006-01-02")
		secs = append(secs, sec)
	}

	w.Header().Add("content-type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(secs)

}

func getBanners(w http.ResponseWriter, r *http.Request) {
	var banners []banner
	for i := 1; i <= 4; i++ {
		imageName := fmt.Sprintf("image%d", i)
		imageURL := fmt.Sprintf("%s/assets/images/%d.jpg", apiurl(), i)
		alt := fmt.Sprintf("image text %d", i)
		bs := banner{Name: imageName, Image: imageURL, Alt: alt, Description: "This is first image"}
		banners = append(banners, bs)
	}
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(banners)
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	var videos []video
	for i := 1; i <= 4; i++ {
		vidName := fmt.Sprintf("video%d", i)
		vidURL := fmt.Sprintf("%s/assets/videos/%d.mp4", apiurl(), i)
		des := fmt.Sprintf("video description text %d", i)
		vid := video{Name: vidName, Video: vidURL, Description: des}
		videos = append(videos, vid)
	}
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	json.NewEncoder(w).Encode(videos)
}

// make directory
func makeDirs() {
	os.MkdirAll("assets/images", 0755)
	os.MkdirAll("assets/videos", 0755)
	os.MkdirAll("assets/news", 0755)
}

// make news images
func makeData() {
	num := 10
	url := "https://picsum.photos/1200/600"
	faker := faker.New()
	var content string
	// make news
	for i := 1; i < num; i++ {
		content = faker.Lorem().Paragraph(2)
		newspath := fmt.Sprintf("%s/assets/news/%d.html", apiurl(), i)
		os.WriteFile(newspath, []byte(content), 0755)

	}
	// get and make images
	for i := 1; i <= num; i++ {
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
	flag.BoolVar(&add, "a", false, "add 10 news and 10 images")
	flag.StringVar(&port, "port", "8080", "default port is 8080")
	flag.Parse()
	if add {
		makeDirs()
		makeData()
		os.Exit(0)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	parseFlag()
	fs := http.FileServer(http.Dir("assets/"))
	http.HandleFunc("/wp-json/nautica/news", getNews)
	http.HandleFunc("/wp-json/nautica/faq", getFAQ)
	http.HandleFunc("/wp-json/nautica/solutions", getSolutions)
	http.HandleFunc("/wp-json/nautica/security", getSecurity)
	http.HandleFunc("/wp-json/nautica/banner", getBanners)
	http.HandleFunc("/wp-json/nautica/videos", getVideos)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Runnig at: ", port)
	fmt.Println(apiurl())
	fmt.Print(showApi())
	http.ListenAndServe(":"+port, nil)
}
