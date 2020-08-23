FROM golang:1.15 as builder

RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/bin/yacasc cmd/yacasc.go


FROM scratch
COPY --from=builder /app/bin/yacasc /yacasc
ENTRYPOINT ["/yacasc"]