FROM golang:1.11.4

WORKDIR /go/src/github.com/TonyCioara/feedback-bot
COPY . .

# RUN go get -d -v ./...

ENV GOMODULES111=on
RUN go install -v ./...

ENV PORT=3000
CMD ["go", "run", "main.go"]