FROM golang:1.21 as builder

WORKDIR /build
COPY . .
RUN go build -ldflags="-s -w" -a -tags netgo -o app ./main.go

FROM alpine:3

ENV ACCENT_COLOR='#f39200'

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /build/app .
COPY index.html .

ENTRYPOINT ["/app/app"]