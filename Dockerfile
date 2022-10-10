FROM golang:1.19-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -v -o /bin/program ./main.go

RUN cp .env /.env

EXPOSE 3000

CMD ["/bin/program"]