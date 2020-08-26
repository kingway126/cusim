#源镜像
FROM golang:alpine

#作者
MAINTAINER kingway "zhengjinwei681@outlook.com"

#设置工作目录
WORKDIR /www/cusim

#将服务器的go工程代码加入到docker容器中
ADD . /www/cusim

# ENV 设置环境变量
ENV GOPROXY="https://goproxy.cn"
ENV GO111MODULE=on

#go构建可执行文件
RUN go build .

#暴露端口
EXPOSE 2222

#最终运行docker的命令
ENTRYPOINT  ["./cusim"]