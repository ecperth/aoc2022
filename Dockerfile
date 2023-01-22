FROM golang:1.19

WORKDIR /app

COPY days ./days
COPY inputs ./inputs
COPY utils ./utils
COPY main.go ./
COPY go.mod ./

RUN go build main.go

ENTRYPOINT ["./main"] 