ARG IMAGE=scratch
ARG OS=linux
ARG ARCH=amd64

FROM 1.19.2-alpine3.16 as builder

WORKDIR /go/src/github.com/kinduff/prometheus_exporter_template

RUN apk --no-cache --virtual .build-deps add git alpine-sdk

COPY . .

RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -ldflags '-s -w' -o binary .

FROM $IMAGE

LABEL name="prometheus_exporter_template"

WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kinduff/prometheus_exporter_template/binary prometheus_exporter_template

CMD ["./prometheus_exporter_template"]
