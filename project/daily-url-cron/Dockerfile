FROM golang:1.15.3 as build

WORKDIR /go/src

RUN go get -u "gorm.io/gorm" && go get -u "gorm.io/driver/postgres"
COPY . /go/src
RUN go build -o /go/app

FROM golang:1.15.3
COPY --from=build /go/app /
CMD ["/app"]