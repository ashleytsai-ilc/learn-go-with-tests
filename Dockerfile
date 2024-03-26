FROM golang:1.20-alpine 

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o learn-go-with-tests main.go

EXPOSE 8080

CMD [ "./learn-go-with-tests" ]
