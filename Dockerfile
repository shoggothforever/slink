
FROM mysql:latest
FROM golang:1.19
MAINTAINER slGroup6
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ENV GO111MODULE=auto
ENV GOPROXY="https://goproxy.io"
RUN go build -o main.go

#docker build -t shortlink-1 .
#docker run --rm -it --name shortlink-demo --network slinknet -p 3000:3000 shortlink-1

#ARG PWD=E:/dockerMySql/root/mysql
#RUN mkdir $PWD/data
#RUN mkdir $PWD/conf
#RUN touch $PWD/conf/my.cnf
#RUN ./mysql.sh