FROM golang:1.16

LABEL base.name="nodeHub_server"

RUN mkdir /app

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080

CMD [ "./job4e_be" ]