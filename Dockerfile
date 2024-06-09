FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /my-go-project

EXPOSE 8080

CMD [ "/my-go-project" ]
