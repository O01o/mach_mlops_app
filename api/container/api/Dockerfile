
FROM golang:1.25-alpine AS backend

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux

COPY cmd/ ./cmd/
COPY internal/ ./internal/

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