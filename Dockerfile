FROM golang:alpine3.19 as base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

FROM base as builder

WORKDIR /app

COPY . .

RUN go build -o scraper .


FROM builder as release

WORKDIR /home/node

COPY --from=builder /app/scraper ./scraper

CMD [ "./scraper" ]