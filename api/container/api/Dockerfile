
FROM golang:1.25-alpine AS backend

WORKDIR /src

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux

COPY src/ .

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/mach-mlops \
    ./cmd/api


FROM scratch

COPY --from=backend /out/mach-mlops /mach-mlops

EXPOSE 8080

ENTRYPOINT ["/mach-mlops"]