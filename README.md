# El-mall
基于小米商城的商业模式开发的电商平台

此项目为前后端分离项目，前端没有完善，接口已经部署到服务器 http://81.68.187.30/#/ ，用 golang 实现接口函数
开发测试文档:https://www.showdoc.cc/1801188651790976 密码：0000


本项目采用了一系列 golang 中比较流行的组件来进行开发

Gin
Gorm
mysql
redis
godotenv
jwt-go

项目结构：

api：用于定义接口函数

cache：redis 相关操作

conf：用于存储配置文件

middleware：应用中间件

model：应用数据库模型

pkg / e：封装错误码

pkg / util：工具函数


serializer：将数据序列化为 json 的函数

server 路由逻辑处理

service：接口函数的实现


本项目使用Go Mod管理依赖。

git clone https://github.com/jiayoukun/El-mall.git

go run main.go

项目运行后启动在 3000 端口
