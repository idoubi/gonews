package main

import (
	"testing"
)

func TestGetFiles(t *testing.T) {
	dir := "/data/others/news"
	files := getFileList(dir)
	for _, file := range files {
		t.Error(file)
	}
}

func TestGetFileHash(t *testing.T) {
	path := "/data/others/news/201711/20171113-20171119.md"
	hash := getFileHash(path)
	t.Error(hash)
}

func TestGetNews(t *testing.T) {
	path := "/data/others/news/201711/20171113-20171119.md"
	news := getNews(path)
	for _, item := range news {
		t.Error(item.ID, item.Title, item.Link.String(), item.Ctime.Format("20060102"))
	}
}
