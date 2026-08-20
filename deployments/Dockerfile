FROM golang:1.23 AS build
WORKDIR /src
COPY go.mod .
COPY . .
RUN go build -trimpath -o /out/workflowd ./cmd/workflowd
FROM gcr.io/distroless/static-debian12
COPY --from=build /out/workflowd /workflowd
EXPOSE 8088
ENTRYPOINT ["/workflowd"]
