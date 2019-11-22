FROM golang:1.12

# set env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /app

COPY . .

RUN go build -o work-manager .

EXPOSE 9000

ENTRYPOINT ["./work-manager"]