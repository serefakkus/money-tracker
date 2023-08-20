# syntax=docker/dockerfile:1

FROM golang


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /sign

EXPOSE 80

CMD [ "/sign" ]