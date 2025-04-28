
<div align=center>
<img src="https://img.shields.io/badge/envoy-1.34.0--dev-blue"/>
<img src="https://img.shields.io/badge/rabbitmq-3.8.9-blue"/>
<img src="https://img.shields.io/badge/meilisearch-1.13.0-blue"/>
<img src="https://img.shields.io/badge/docker-28.0.1-blue"/>
<img src="https://img.shields.io/badge/golang-1.22.5windows%2Famd64-blue"/>
</div>



# 一个用Go开发的仿B站项目后端

> **前言**：本项目是作者按当时的观察想象分析所写，代码中不免有一些令人疑惑的方法。本人有些腻了，还不想开发剩余功能了，有能力的小伙伴可以自行优化升级。总的来说，本项目每个文件夹都是手动创建的，应该都好理解。最近作者打算搞别的感觉有趣的东西去了。     ---- 2025/4/28 Yux


## 项目：随便取了platform_backend

- 使用 Go 语言仿b站开发后端，实现了一个视频网站所具备的主要功能。
- 采用前后端分离模式开发，后端主要用gRPC协议与RabbitMQ作为 通信手段。

### 项目结构
- **log 日志文件**
- **deploy 容器配置文件(kubernetes由于作者笔记本配置跟不上，所以放弃了，请看docker内介绍)**
- **generated protobuf生成文件的目录**
- **microservices 各个微服务所在文件夹，每个服务都有自己的入口**
- **pkg 一些共用的工具函数或者中间件的客户端方法**
- **protobuf gRPC的扩展语言定义所在**
- **script 一些go写的脚本，主要是用于将内容插入数据库**

### microservices内 结构
- **client(仅aggregator存在) gRPC客户端文件夹** 
- **cache 缓存方法所用的文件夹**
- **cmd 程序入口**
- **config 依赖管理所在文件夹**
- **internal 接口函数的内部逻辑**
- **log 日志文件**

- **messaging 消息队列所在文件夹**
- **messaging/dispatch 收集分散请求的设计，有缺陷**
- **messaging/receiver 消息队列订阅者**

- **recommend 推荐系统所在文件夹**
- **repository 数据库方法**
- **service gRPC服务器启动以及接口所在**
- **tools 工具函数**

### **声明：本项目只用作学习参考，无任何商业用途，若他人使用本项目造成的侵权问题，本人概不负责**

## 项目功能

- **首页视频随机推荐**
- **用户注册登录**
- **个人中心信息修改**
- **视频投稿**
- **视频审核**
- **内容搜索（视频）**
- **视频详情页（观看 + 点赞 + 收藏 + 评论）**
- **个人空间（用户作品 + 收藏夹等）**

其他由于时间问题，暂停开发的功能：

- 视频分区功能
- 动态服务
- 消息服务
- 收藏夹分类功能
- 视频合集功能
- 数据统计服务
- 弹幕服务

## 部署

- 请先根据deploy/docker内readme提示进行Docker部署,然后到microservice内各个文件夹下启动cmd的main.go

### 后端

1. 项目使用 `阿里云OSS` 存储视频，图片，请自行准备。
2. 无框架开发，全手搓，最初抱着锻炼原生开发的心态做的。