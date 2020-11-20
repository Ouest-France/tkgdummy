FROM golang AS builder
COPY . /go/src/github.com/Ouest-France/tkgdummy
WORKDIR /go/src/github.com/Ouest-France/tkgdummy
RUN CGO_ENABLED=0 go build -ldflags="-w -s"

FROM scratch
COPY --from=builder /go/src/github.com/Ouest-France/tkgdummy/tkgdummy /tkgdummy
CMD ["/tkgdummy"]