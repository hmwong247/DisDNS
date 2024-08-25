### DisDNS
___
Nothing to due with DNS service

DisDNS is just a simple proxy application to retrieve your remote (or home) IP address, when you cannot afford a static IP address and a domain name.

DisDNS use Discord bot as the front-end application, and ipify API to get your public IP address

### Usage
___
1. create your Discord bot
2. create a `.env` file in the project root directory
3. copy and paste your Discord bot token to the `.env` file
```sh
export BOT_TOKEN = [YOUR DISCORD BOT TOKEN]
```
4. source the `.env` file
```sh
source .env
```
5. run the application
```sh
./DisDNS
```

Just send a message to the bot with "ip" to get your remote IP address
