FROM golang:1.19-alpine AS go

RUN apk update && apk add --no-cache git

ENV USER=appuser
ENV UID=1001
ENV APP=uml-caddy

RUN adduser \
    --disabled-password\
    --shell "sbin/nologin"\
    --no-create-home\
    --uid "${UID}"\
    "${USER}"

WORKDIR $GOPATH/src/$APP

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/$APP ./cmd/$APP

FROM golang:1.19-alpine

WORKDIR $GOPATH/bin

COPY --from=go /etc/passwd /etc/passwd
COPY --from=go /etc/group /etc/group

COPY --from=go /go/bin/$APP /go/bin/$APP

COPY ./templates/ /go/bin/templates/

USER $USER

ENTRYPOINT [ "/go/bin/uml-caddy" ]
