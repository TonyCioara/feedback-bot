FROM golang:1.11.4

RUN mkdir /feedback-bot-run
WORKDIR /feedback-bot-run

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# RUN go get -d -v ./...

# ENV GO111MODULE=on
ENV PORT=3000

EXPOSE 3000

CMD ["go", "run", "main.go"]