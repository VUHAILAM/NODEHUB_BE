FROM golang:1.16

LABEL base.name="nodeHub_server"

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV MYSQL_DSN="admin:13091999@tcp(nodehubdb.coghck9xckk7.ap-southeast-1.rds.amazonaws.com:3306)/nodehub?parseTime=true"

RUN go build

EXPOSE 8080

CMD [ "./job4e_be" ]