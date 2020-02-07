FROM golang:1.12-alpine as builder

ENV GO111MODULE=on

RUN apk add --no-cache git

WORKDIR /svc

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /svc main.go

FROM scratch
COPY --from=builder /svc /svc
ENTRYPOINT [ "/svc"]