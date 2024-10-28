# 天气查询作品

## 说明
使用 go + 原生 http 构建后端，接口采用[新知天气](https://www.seniverse.com/)，前端采用 ts + vite + react + unocss 构建。

## 项目结构
```
├── api
│   └── api.go
├── web
|   |── public
|   |   └── icon             // 天气图标
|   |── src
|   |   |── hooks
|   |   |   └── weather.ts   // 天气请求封装
|   |   |── App.tsx          // 主界面
|   |   └── main.tsx         // 入口文件
|   |── uno.config.ts         // uno 配置文件
|   |── vite.config.ts        // vite 配置文件
│   └── index.html           // 单入口页面
├── main.mod                 // 后端主程序
└── build                    // 交叉编译脚本



```


## 项目运行

项目已编译，可直接运行，也可自行编译。

- windows
```shell
# 后端运行
cd builds && ./weather-win-amd64.exe
```

- linux
```shell
# 后端运行
cd builds && ./weather-linux-amd64
```

- mac intel
```shell
# 后端运行
cd builds && ./weather-darwin-amd64
```

- mac arm
```shell
# 后端运行
cd builds && ./weather-darwin-arm64
```



## 项目调试

```shell
# 后端运行
go run main.go

# 前端运行
cd web && bun install && bun dev

```

## 项目打包

1、先打包前端

因先后端嵌入前端文件，所以打包前需要先打包前端

```shell
cd web && bun run build
```
2、打包后端

需要在 macos、linux 上运行，windows 使用 win 命令

```shell
# 后端打包
./build
```


