FROM golang:alpine3.19 as base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

FROM base as builder

WORKDIR /home/app

COPY . .

RUN go mod download

RUN go run github.com/steebchen/prisma-client-go generate

RUN go build .


FROM builder as release

WORKDIR /home/node

COPY --from=builder ./main /home/node/main

CMD [ "./main" ]