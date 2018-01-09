# *GoCN每日新闻*  可视化浏览与检索平台


## 介绍

- 线上Demo：[GoNews](http://gonews.idoubi.cc)
- 数据来源：[GoCN每日新闻](https://github.com/gocn/news)
  
## 部署


- 获取源码

```
go get -u github.com/mikemintang/gonews
```

- 获取新闻数据

```
git clone https://github.com/gocn/news /data/news
```

- 解析数据

```
cd /data/go/src/github.com/mikemintang/gonews
go install
nohup gonews -d /data/news > ./gonews.log 2>&1 
```

- 启动Api

```
nohup gonews -a api -p 8017 > ./gonews.log 2>&1 &
```

- 前端部署

```
cd /data/go/src/github.com/mikemintang/gonews/web
npm install
npm run build
```

- Nginx配置

```
server {
    listen       80;
    server_name gonews.idoubi.cc;
    index index.html index.htm index.php;
    root  /data/go/src/mikemintang/gonews/web;

    location /api {
        proxy_pass http://127.0.0.1:8017;
    }
}
```

- Shell脚本

```
#!/bin/sh

cd /data/news
git pull origin master
cd /data/go/src/gitbub.com/mikemintang/gonews
./gonews -d /data/news
```

- 定时任务

```
crontab -e
*/10 * * * * /bin/sh /data/shell/cache_news.sh
```


## 用到的技术

### golang包

- github.com/go-redis/redis
- encoding/json
- flag
- net/http
- net/url
- strconv
- sync
- crypto/md5
- fmt
- io
- io/ioutil
- net/url
- os
- path/filepath
- regexp
- strconv
- strings
- time

### npm包

- vue
- vuex
- vue-router
- axios
- moment
- mockjs


欢迎提交Pull Request


