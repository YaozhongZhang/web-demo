FROM golang:alpine as build

WORKDIR /go/src/demo
COPY . .

RUN go build -o /go/bin/app

FROM alpine

COPY --from=build /go/bin/app /usr/bin/app

CMD ["app"]