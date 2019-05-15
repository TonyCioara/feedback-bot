FROM golang:1.11.4

WORKDIR /github.com/TonyCioara/feedback-bot
COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]