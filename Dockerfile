FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /home/geekbim/go/src/sagara

COPY . .

RUN go mod tidy

RUN cd cmd && go build -o ../bin

ENTRYPOINT ["/home/geekbim/go/src/sagara/bin"]