# bubble
Gin示例 代码

## 项目编译

- Windows

```
# 编译代码
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

- Linux

```
# 编译代码
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

- MAC

```
# 编译代码
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go  # linux
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go  # windows
```

## 项目启动

```
cd bubble/
./main
```

## 实现功能
- 代办清单的增删改查

## 有待完善
- 增加配置文件
- docker 容器化
