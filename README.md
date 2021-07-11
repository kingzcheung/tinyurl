# Tiny URL

短网址生成服务。

### 使用

1. 构建

```shell
make # 构建
```

2. 启动

```shell
./tinyurl -dbpath=tinyurl.db #-dbpath 可选
```

3. 打开地址

```shell
http://127.0.0.1:3311
```



### 开发

环境：

- go 1.16 以上
- nodejs 12 以上

后端依赖

```
go mod tidy
```

前端依赖

```shell
cd web && yarn
```

启动服务:

```shell
go run cmd/tinyurl/main.go
yarn run dev
```



