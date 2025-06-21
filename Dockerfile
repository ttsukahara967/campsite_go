FROM golang:1.24

RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/app

# go.mod, go.sum, .air.tomlなどを app/ からCOPY
COPY app/go.mod app/go.sum ./
COPY app/.air.toml ./
RUN go mod download

COPY app/ ./
# Swagger docsもapp/docs以下に生成されていればOK
# COPY app/docs/ ./docs/

RUN ls -l /go/src/app

CMD ["air", "-c", ".air.toml"]
