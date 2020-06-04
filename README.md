<p align="center">
    <img src="https://gitee.com/bullteam/lime/raw/master/docs/image/logo.png" height="145">
</p>


# Lime Soft
青橙软件，专注于正版小说、漫画数字类产品网络分销开源免费软件
该软件居于 `golang + vue + zeus(宙斯权限系统)`


## 项目介绍
> - `Lime Soft` 定位为数字阅读营销平台，为推动文化产业更加繁荣而提供网络软件支持。
> - 居于 `Zeus 宙斯`权限系统开发。
> - 该软件定位为开源免费，任何个人和团队都可以在此基础之上开发使用。
> - `Lime` 支持多国语言，多终端展示。
> - 计划支持平台包括PC 桌面软件、WEB 站、Android、IOS、小程序、H5移动端等。
> - 该软件支持分销功能，让数字数字类从业者和作者得到更大的收益。
> - 严禁使用该软件用于非法盗版版权用途，如果采用该软件用于非法用途一切后果自负。
## 功能
- [x] H5阅读客户端
- [x] 后台小说管理
- [x] 后台漫画管理

## 安装部署文档
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
# H5手机阅读页面
cd pkg/h5
npm install
npm run build

cd ~/lime

export LIME_MYSQL_USERNAME=root
export LIME_MYSQL_PASSWORD=123456
export LIME_MYSQL_HOST=127.0.0.1
export LIME_MYSQL_DB=zeus
export LIME_MYSQL_PORT=3306
export LIME_REDIS_HOST=127.0.0.1
export LIME_REDIS_PORT=6379
export LIME_REDIS_PASSWORD=""
./lime server -c ./config/in-local.yaml --cors=true
```

## 相关截图
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/books.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/chapters.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/category.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/upload.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/1.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/2.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/3.png"></img>
<img src="https://gitee.com/bullteam/lime/raw/master/docs/image/4.png"></img>

## 贡献代码

非常欢迎优秀的开发者来贡献`Lime`。在提Pull Request之前，请首先阅读源码，了解原理和架构。如果不懂的可以加他的微信 `cosohuang` 注明 `Lime`。

## 社区

如果您觉得 `Lime` 对您有帮助或者想使用该软件，请扫描下方群二维码，请加微信 `cosohuang` 并注明`lime 开源交流`，他会将你拉入群一起学习讨论。

<p align="center">
    <img src="https://gitee.com/bullteam/lime/raw/master/docs/image/wx.jpg" height="360">
</p>
