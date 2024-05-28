FROM golang:1.22.3-alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin ./...

FROM scratch

COPY --from=build /usr/local/bin/tui-cv /tui-cv

CMD ["/tui-cv"]
