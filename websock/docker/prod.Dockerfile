FROM golang:1.16 as build

WORKDIR /app
COPY ./ /app
RUN env GOOS=linux GOARCH=386 go build -o bin/websock .

FROM alpine as runtime

COPY --from=build /app/bin/websock /

CMD ["/websock"]