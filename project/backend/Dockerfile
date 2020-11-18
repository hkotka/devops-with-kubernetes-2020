FROM golang:1.15.3 as build

WORKDIR /go/src

COPY . /go/src
RUN go build -ldflags "-s -w" -o /go/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/app /
CMD ["/app"]