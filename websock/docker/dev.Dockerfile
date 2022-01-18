FROM golang:1.16

WORKDIR /app
COPY ./ /app
RUN go install

CMD ["modd"]
