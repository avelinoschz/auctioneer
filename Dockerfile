FROM golang:1.22.2-bullseye as builder
WORKDIR /app
COPY Makefile go.mod /app/
COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
RUN go mod download
RUN make build-go

FROM scratch
WORKDIR /app
COPY --from=builder /app/generated/bin/auctioneer .
CMD ["./auctioneer"]
