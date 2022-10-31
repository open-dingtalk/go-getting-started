[![Build Status](https://github.com/open-dingtalk/go-getting-started/workflows/Go/badge.svg?branch=master)](https://github.com/open-dingtalk/go-getting-started/actions)
![GitHub issues](https://img.shields.io/github/issues/open-dingtalk/go-getting-started)
![GitHub](https://img.shields.io/github/license/open-dingtalk/go-getting-started)

# go-getting-started
钉钉酷应用脚手架go语言版本。

## 功能介绍
- 获取钉钉用户身份
- 通过机器人向群里发送普通消息；
- 通过机器人向群里发送互动卡片；
- 通过机器人向群里发送吊顶卡片；

## 启动配置
参考config/config.yaml.example

## 启动方式

### 源码启动
```bash
$ go mod tidy
$ go run cmd/main.go -c config/config.yaml -p 7001 -s static/
```

