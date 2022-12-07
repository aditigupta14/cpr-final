# syntax=docker/dockerfile:1

FROM golang:1.19.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /cpr/

RUN ls /cpr/

EXPOSE 8080

CMD [ "/cpr/Parking.com" ]