FROM golang:1.15.3 as build

WORKDIR /go/src
COPY . /go/src

RUN go get -d -v ./...
RUN go build -o /go/app

FROM gcr.io/distroless/base
COPY --from=build /go/app /
COPY templates /templates
CMD ["/app"]