package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var dir string
var act string

var wg sync.WaitGroup

func init() {
	flag.StringVar(&dir, "d", "", "the path of news to parse")
	flag.StringVar(&act, "a", "cache", "cache news or show news in the web")
}

func main() {
	flag.Parse()
	if act == "api" { // web服务
		http.HandleFunc("/news", getNewsApi) // 获取新闻数据api
		http.ListenAndServe(":8017", nil)
	} else { // 缓存数据操作
		files := getFileList(dir)
		for _, file := range files {
			wg.Add(1)
			go cacheNews(file)
		}
		wg.Wait()
	}
}

// 并发缓存数据
func cacheNews(path string) {
	defer wg.Done()
	news := getNews(path)
	for _, item := range news {
		cache := map[string]interface{}{
			"id":    strconv.Itoa(int(item.ID)),
			"title": item.Title,
			"link":  item.Link.String(),
			"ctime": item.Ctime.Format("20060102"),
		}
		setNewsCache(cache) // 缓存数据
	}
}

// 获取新闻的api
func getNewsApi(w http.ResponseWriter, req *http.Request) {
	page := int64(1)
	per := int64(50)
	keyword := ""
	query, err := url.ParseQuery(req.URL.RawQuery)
	if err == nil {
		if len(query["page"]) > 0 && query["page"][0] != "" {
			iPage, _ := strconv.Atoi(query["page"][0])
			page = int64(iPage)
		}
		if len(query["per"]) > 0 && query["per"][0] != "" {
			iPer, _ := strconv.Atoi(query["per"][0])
			per = int64(iPer)
		}
		if len(query["keyword"]) > 0 && query["keyword"][0] != "" {
			keyword = query["keyword"][0]
		}
	}
	var news []map[string]string
	if keyword != "" {
		news = searchNews(keyword, page, per)
	} else {
		news = getAllNews(page, per)
	}
	jData, _ := json.Marshal(news)
	w.Write(jData)
}
