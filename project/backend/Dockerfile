FROM golang:1.15.3 as build

WORKDIR /go/src

RUN go get -u github.com/gin-gonic/gin && go get -u github.com/gin-contrib/cors && go get -u "gorm.io/gorm" && go get -u "gorm.io/driver/postgres" && go get -u "gorm.io/driver/sqlite"
COPY . /go/src
RUN go build -o /go/app

FROM golang:1.15.3
COPY --from=build /go/app /
CMD ["/app"]