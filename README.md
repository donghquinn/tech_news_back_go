# Scraper Backend

- This Code is Golang ported program of [my nest.js program](https://github.com/donghquinn/tech_news_backend)

- I converted my original nest.js based codes to golang because of response time.

- Furthermore it seems lighter than nestjs codes to me.

## Dependencies

- go-gin

```shell
go get -u github.com/gin-gonic/gin
```

- prisma

```shell
go get -u github.com/steebchen/prisma-client-go
```

- godotenv

```shell
go get -u github.com/joho/godotenv
```

## Prisma

- Data Entities are defined at prisma/schema.prisma
- Can generate and activate prisma by the command below

```shell
go run github.com/steebchen/prisma-client-go generate
```
