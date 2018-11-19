FROM golang:1.11 as build

WORKDIR /src
COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
COPY --from=build /src/app /usr/local/bin/tribonacci

ENV TRIBONACCI_ADDR :8000
ENV TRIBONACCI_GRACE_TIMEOUT 10

EXPOSE 8000

CMD [ "tribonacci" ]
