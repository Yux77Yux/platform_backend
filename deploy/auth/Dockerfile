# 使用 Golang 基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将项目的 go.mod 和 go.sum 文件复制到工作目录中
COPY ../../../go.mod ../../../go.sum ./

# 下载依赖
RUN go mod download

# 复制项目源代码到容器中
COPY ./ ./microservices/auth/
COPY ../shared/ ./microservices/shared/

# 编译程序
RUN go build -o auth-service .

# 暴露应用端口
EXPOSE 50011

# 启动应用
CMD ["./auth-service"]
