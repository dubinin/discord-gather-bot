FROM golang:1.7-wheezy

ADD . /go/src/bitbucket.org/puumku/discord-gather-bot

RUN go get github.com/bwmarrin/discordgo && go get github.com/Sirupsen/logrus

RUN go install github.com/dubinin/discord-gather-bot

ENTRYPOINT /go/bin/discord-gather-bot
