FROM golang:latest as build

LABEL maintainer "jxsl13@gmail.com"
WORKDIR /build
COPY *.go ./
COPY go.mod .
COPY go.sum .
RUN go get -d && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o main .


FROM alpine
WORKDIR /app
COPY --from=build /build/main .
EXPOSE 3000/tcp
ENTRYPOINT ["/app/main"]
#VOLUME ["/in", "/out"]
