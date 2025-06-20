FROM golang:1.24

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY app/ ./
COPY docs/ ./docs/

RUN ls -l /go/src/app

CMD ["sh", "-c", "go run *.go"]
#CMD ["go", "run", "*.go"]
#CMD ["go", "run", "main.go"]
#CMD ["sleep", "600"]
#CMD ["tail", "-f", "/dev/null"]


