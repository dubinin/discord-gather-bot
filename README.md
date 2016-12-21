GameType (3v3, 2v2, 1v1, 2s, 3s)
CommandType (on, off, help, add, del, info)

---

Add bot to server:
    https://discordapp.com/oauth2/authorize?client_id={client_id}&scope=bot&permissions=3072

Get token:
    https://discordapp.com/oauth2/token

---

docker build -t discord-gather-bot .

docker run -d --name bot -e "BOT_TOKEN={token}" -e "BOT_CHANNEL={channel_id}" -e "BOT_ADMINS={comma separated ids of admins}" discord-gather-bot
