FROM golang:1.25 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY src src/
COPY docs docs/
RUN CGO_ENABLED=0 go build -o /bin/api .

FROM scratch
COPY --from=build /bin/api /bin/api
CMD ["/bin/api"]
