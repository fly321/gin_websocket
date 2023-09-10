# 使用 Golang 镜像作为基础镜像
FROM golang:latest
# 更新源
RUN apt-get update
# 安装 git
RUN apt-get install -y git

# 设置工作目录
WORKDIR /go/src/app

# 复制项目文件到工作目录中
COPY . .

# 安装依赖
RUN go mod download

# 编译 Go 代码
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]