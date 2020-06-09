FROM golang:1.13

WORKDIR /app/graph-study
COPY . .

RUN go get ./...
RUN go build

ENV PORT=8080

CMD ["./graph-study"]