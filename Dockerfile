FROM golang:1.22.0-alpine3.19 as build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

# stage imagem final
FROM scratch

WORKDIR /app

COPY --from=build /app/api ./

EXPOSE 8000

CMD [ "./api" ]