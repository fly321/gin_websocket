# 使用 Golang 镜像作为基础镜像
FROM golang:1.19.13

# 设置工作目录
WORKDIR /go/src/app

# 复制项目文件到工作目录中
COPY . .

# 安装依赖
RUN go mod tidy && \n    go get -u github.com/gin-gonic/gin && \n

# 编译 Go 代码
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 修改文件权限
RUN chmod +x main

# 运行应用
CMD ["./main"]