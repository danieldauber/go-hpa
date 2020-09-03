FROM golang:1.14-alpine as builder

WORKDIR /go/src/app
COPY ./src/app .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o main

FROM scratch
COPY --from=builder  /go/src/app . 

EXPOSE 8000

CMD [ "./main" ]