# 多阶段构建：提升构建速度，减少镜像大小

# 从官方仓库中获取 1.17 的 Go 基础镜像
FROM golang:1.17.2-alpine3.14 AS builder
# 设置工作目录
WORKDIR /go/src/sport
# 复制项目文件
ADD . /go/src/sport
# 下载依赖
RUN export GOPROXY="https://goproxy.cn,direct" && go get -d -v ./...
# 构建名为"app"的二进制文件
RUN go build -o app .

# 获取轻型 Linux 发行版，大小仅有 5M 左右
FROM alpine:latest
# 将上一阶段构建好的二进制文件复制到本阶段中
COPY --from=builder /go/src/sport/app .
# 将配置文件复制到本阶段
COPY --from=builder /go/src/sport/conf ./conf
# 设置监听端口
EXPOSE 8099
# 配置启动命令
ENTRYPOINT ["./app"]
