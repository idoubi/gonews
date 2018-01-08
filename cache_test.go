package main

import (
	"testing"
)

func TestSetNewsCache(t *testing.T) {
	cache := map[string]interface{}{
		"id":    "1",
		"title": "这是今天的新闻",
		"link":  "http://news/1122",
		"ctime": "20171211",
	}
	err := setNewsCache(cache)
	t.Error(err)
}

func TestGetNewsCache(t *testing.T) {
	key := "gonews:20171211:1"
	cache, _ := getNewsCache(key)
	t.Errorf("%+v", cache)
}

func TestGetAllNews(t *testing.T) {
	newsList := getAllNews(2, 3)
	t.Error(newsList)
}

func TestSearchNews(t *testing.T) {
	newsList := searchNews("学习", 2, 2)
	newsList2 := searchNews("201801", 1, 3)
	t.Error(newsList, newsList2)

}
