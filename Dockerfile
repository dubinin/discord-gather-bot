FROM golang:latest

ADD . /go/src/github.com/dubinin/discord-gather-bot

RUN go get github.com/bwmarrin/discordgo

RUN go get github.com/Sirupsen/logrus

RUN go install github.com/dubinin/discord-gather-bot

ENTRYPOINT /go/bin/discord-gather-bot
