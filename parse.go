// 解析数据
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 单条新闻数据结构
type news struct {
	ID    int64
	Title string
	Link  url.URL
	Ctime time.Time
}

// 遍历文件夹
func getFileList(dir string) []string {
	files := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if strings.Contains(path, ".git") {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		baseName := filepath.Base(path)
		if !strings.Contains(baseName, "-") {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".md" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

// 计算文件哈希
func getFileHash(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return ""
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 获取新闻
func getNews(path string) []news {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	newsList := []news{}
	reg := regexp.MustCompile(`(?i)#{0,3}\s*GoCN每日新闻\(([\d-]*)\)\n+\D*((.*\n)+?\n)`)
	all := reg.FindAllSubmatch(data, -1)
	for _, item := range all {
		singleNews := news{}
		loc, _ := time.LoadLocation("Local")
		ctime, err := time.ParseInLocation("2006-01-02", string(item[1]), loc)
		if err == nil {
			singleNews.Ctime = ctime
		}
		sreg := regexp.MustCompile(`(\d)\.\s*(.*)\s+(http.*)\n?`)
		sall := sreg.FindAllSubmatch(item[2], -1)
		for _, sitem := range sall {
			id, err := strconv.Atoi(string(sitem[1]))
			if err == nil {
				singleNews.ID = int64(id)
			}
			singleNews.Title = strings.TrimSpace(string(sitem[2]))
			surl, err := url.Parse(string(sitem[3]))
			if err == nil {
				singleNews.Link = *surl
			}
			newsList = append(newsList, singleNews)
		}
	}

	return newsList
}
