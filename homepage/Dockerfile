FROM golang:1.25 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY docs docs/
RUN CGO_ENABLED=0 go build -o /bin/homepage .

FROM scratch
COPY static ./static
COPY templates ./templates
COPY --from=build /bin/homepage /bin/homepage
CMD ["/bin/homepage"]
