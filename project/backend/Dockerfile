FROM golang:1.15.3 as build

WORKDIR /go/src

RUN go get -u github.com/gin-gonic/gin && go get -u github.com/gin-contrib/cors && go get -u "gorm.io/gorm" && go get -u "gorm.io/driver/postgres" && go get -u "gorm.io/driver/sqlite" && go get -u "github.com/gin-contrib/logger" && go get -u "github.com/rs/zerolog" && go get -u "github.com/rs/zerolog/log"
COPY . /go/src
RUN go build -o /go/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/app /
CMD ["/app"]