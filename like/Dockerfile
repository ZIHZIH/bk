# FROM ... AS builder : 表示依赖的镜像只是使用在编译阶段
FROM golang:1.20 AS builder

# 编译阶段的工作目录，也可以作为全局工作目录
WORKDIR /app

# 把当前目录的所有内容copy到 WORKDIR指定的目录中
COPY . .

# 设置路径
ENV APP_PATH=./cmd/like

# 设置代理环境变量
ARG GOPROXY="https://goproxy.cn,direct"

# 下载所有依赖
RUN go mod download

# 构建应用程序。禁用 CGO 并设置目标操作系统为 linux
RUN CGO_ENABLED=0 GOOS=linux go build -o like $APP_PATH

# 使用 alpine 镜像作为基础，它是一个小型的 Linux 发行版
FROM alpine:latest

# copy建造层中的执行文件到新的一层
COPY --from=builder /app/like .

# 运行应用程序
CMD ["./like"]