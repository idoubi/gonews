// 数据缓存
package main

import (
	"sort"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	if client == nil {
		client = newClient()
	}
}

func newClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       cacheDB,
	})
	return client
}

// 设置新闻缓存
func setNewsCache(cache map[string]interface{}) error {
	key0 := cachePrefix
	key1 := ""
	key2 := ""
	if value, ok := cache["ctime"].(string); ok {
		key1 = key0 + ":" + value
		err := client.SAdd(key0, value).Err()
		if err != nil {
			return err
		}
	}
	if value, ok := cache["title"].(string); ok {
		key2 = key1 + ":" + value
		err := client.SAdd(key1, value).Err()
		if err != nil {
			return err
		}
		return client.HMSet(key2, cache).Err()
	}
	return nil
}

// 获取新闻缓存
func getNewsCache(key string) (map[string]string, error) {
	return client.HGetAll(key).Result()
}

// 获取全部新闻
func getAllNews(page int64, per int64) ([]map[string]string, int64) {
	key0 := cachePrefix
	keys1, _ := client.SMembers(key0).Result()
	newsList := map[string]map[string]string{}
	for _, key1 := range keys1 {
		keys2, _ := client.SMembers(key0 + ":" + key1).Result()
		for _, key2 := range keys2 {
			news, err := getNewsCache(key0 + ":" + key1 + ":" + key2)
			if err == nil {
				newsList[news["ctime"]+news["id"]] = news
			}
		}
	}
	allNews := sortNews(newsList)
	pageNews := []map[string]string{}
	var i int64 = 0
	for _, item := range allNews {
		if i >= (page-int64(1))*per && i < page*per {
			pageNews = append(pageNews, item)
		}
		i++
	}
	return pageNews, i
}

// 条件查询
func searchNews(keyword string, page int64, per int64) ([]map[string]string, int64) {
	keys, _ := client.Keys(cachePrefix + "*" + keyword + "*").Result()
	newsList := map[string]map[string]string{}
	for _, key := range keys {
		news, err := getNewsCache(key)
		if err == nil {
			newsList[news["ctime"]+news["id"]] = news
		}
	}
	allNews := sortNews(newsList)
	pageNews := []map[string]string{}
	var i int64 = 0
	for _, item := range allNews {
		if i >= (page-int64(1))*per && i < page*per {
			pageNews = append(pageNews, item)
		}
		i++
	}
	return pageNews, i
}

// 对新闻进行排序
func sortNews(raw map[string]map[string]string) []map[string]string {
	keys := []string{}
	data := []map[string]string{}
	for key := range raw {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, key := range keys {
		data = append(data, raw[key])
	}
	return data
}
