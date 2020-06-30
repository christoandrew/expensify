FROM golang:latest

WORKDIR /usr/app

COPY . .

RUN go get github.com/pilu/fresh

RUN go build

EXPOSE 9087

CMD ["fresh"]