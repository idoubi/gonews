package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// 自动获取数据, interval单位为分钟
func initDataPuller(dir string, interval int64) {
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	go func() {
		for {
			// 拉取数据
			cmd := exec.Command("git", "pull", "origin", "master")
			cmd.Dir = dir
			out, err := cmd.Output()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", "Pull failed")
				fmt.Fprintf(os.Stderr, "%s\n", err)
				fmt.Fprintf(os.Stderr, "%s\n", out)
			} else {
				if !strings.Contains(string(out), "Already up to date") {
					fmt.Fprintf(os.Stdout, "%s\n", "Pull success")

					// 缓存数据操作
					files := getFileList(dir)
					for _, file := range files {
						wg.Add(1)
						go cacheNews(file)
					}
					wg.Wait()
					fmt.Fprintf(os.Stdout, "%s\n", "Success to cache news")
				}
			}
			// 定时器
			<-ticker.C
		}
	}()
}
