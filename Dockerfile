FROM golang:1.19.9 as builder

# first (build) stage

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o curriculum-api

# final (target) stage

FROM alpine:3.18.0
COPY --from=builder /app/curriculum-api /
CMD ["/curriculum-api"]
