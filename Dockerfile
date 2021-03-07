FROM golang

WORKDIR /app

EXPOSE 40001 8080

RUN go get github.com/go-delve/delve/cmd/dlv

COPY ./go.mod ./go.sum main.go ./

CMD [ "dlv", "debug", "./main.go", "--listen=:38081", "--headless=true", "--api-version=2", "--log" ]