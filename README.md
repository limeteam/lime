# lime
青橙软件平台，专注于正版小说、漫画数字类产品网络分销开源软件

## 部署文档
```
git clone git@github.com:limeteam/lime.git
export GOPROXY=https://goproxy.cn
export GO111MODULE=on
#后端编译
go build -o lime
#前端编译
cd pkg/ui
npm install
#正常情况下，会生成dist目录，可自己部署web服务器(如nginx)，提供前端服务
npm run build:prod
cd ~/lime

export LIME_MYSQL_USERNAME=root
export LIME_MYSQL_PASSWORD=123456
export LIME_MYSQL_HOST=127.0.0.1
export LIME_MYSQL_DB=zeus
export LIME_MYSQL_PORT=3306
export LIME_REDIS_HOST=127.0.0.1
export LIME_REDIS_PORT=6379
export LIME_REDIS_PASSWORD=""

#修改in-local.yamln内部的project.merge为false,然后再启动
./lime server -c ./config/in-local.yaml --cors=true
```