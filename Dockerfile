# sotaneum/go-git-push-pull:latest
FROM golang:alpine as build

# Alpine에 Git 설치(Go 라이브러리가 Git를 사용)
RUN apk add --no-cache bash git openssh
RUN apk --no-cache add tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /go/src/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM scratch as final

COPY --from=build /go/src/main .
# TimeZone Error
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
# x509: certificate signed by unknown authority
# 해결 출처 : https://titanwolf.org/Network/Articles/Article?AID=674fa38a-dfd5-4f67-bfd9-3b3197ac3ad4#gsc.tab=0
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080/tcp
EXPOSE 8080/udp

ENV TZ=Asia/Seoul

CMD ["/main"]
