FROM golang:1.15.3 as build

WORKDIR /go/src

RUN go get "gorm.io/gorm" && go get "gorm.io/driver/postgres" && go get "github.com/gin-gonic/gin" && go get "github.com/gin-contrib/cors"
COPY . /go/src
RUN go build -o /go/app

FROM gcr.io/distroless/base
COPY --from=build /go/app /
CMD ["/app"]