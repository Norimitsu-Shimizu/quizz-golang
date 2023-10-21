FROM golang:1.21.1-alpine as build

ENV GO111MODULE="on"

WORKDIR /go/src/github.com/Norimitsu-Shimizu/quizz-golang

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git

COPY ./ .

RUN go build -o app /go/src/github.com/Norimitsu-Shimizu/quizz-golang/main.go

# 本番用のイメージ(仮)
FROM alpine:3.16 as production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/Norimitsu-Shimizu/quizz-golang/api/app .
EXPOSE 8080
CMD ["./app"]

# ローカル用のイメージ
FROM golang:1.21.1-alpine as development

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
