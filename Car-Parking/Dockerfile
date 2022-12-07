# syntax=docker/dockerfile:1

FROM golang:1.19.1

WORKDIR /app

COPY ./ ./

RUN go mod download

COPY . ./

RUN go build -o /Car-Parking/

RUN ls /Car-Parking/

EXPOSE 8080

CMD [ "/Car-Parking/Parking.com" ]
