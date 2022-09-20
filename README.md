# **HertzBoot**

## Go Web Template

### Base on

- [Hertz](https://github.com/cloudwego/hertz)
- [GORM(MySql)](https://github.com/go-gorm/gorm)
- [Swag](https://github.com/swaggo/swag)
- [Redis](https://github.com/go-redis/redis)
- [Casbin](https://github.com/casbin/casbin)
- [Zap](https://github.com/casbin/casbin)
- [Viper](https://github.com/spf13/viper)

## Tree

```shell
├── Dockerfile
├── Makefile
├── README.md
├── app
│   ├── global
│   ├── jobs
│   ├── middleware
│   ├── provider
│   ├── request
│   └── response
├── config      配置目录
├── config.docker.yaml
├── config.yaml
├── config.yaml.example
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── latest_log  
├── log         日志目录
├── main.go
├── modules
│   ├── admin
│   ├── api
│   └── core
├── resource    资源文件
├── source      初始表数据
└── tools       工具类

