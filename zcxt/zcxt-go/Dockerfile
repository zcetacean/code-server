FROM golang:latest
#定义环境变量

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV GOROOT=/usr/local/go
ENV GOPATH=/www/go/
ENV PATH=$GOPATH/bin:$PATH
#首先安装beego bee mysql扩展
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
## 复制Dockerfile 同目录的所有文件到docker 中的指定位置
COPY . /www/go/src/zcxt_go/
## 开放 9000 端口
EXPOSE 9000
## 启动执行 bee run  命令 
CMD ["bee", "run"]