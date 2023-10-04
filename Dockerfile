FROM golang:1.21 as builder
WORKDIR /usr/src/cse138
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pa1 .

FROM debian:bookworm-slim
COPY --from=builder /usr/src/cse138/pa1 /usr/local/bin/pa1
EXPOSE 8090
CMD ["pa1"]