# discord-stock-tickerboard

[![GitHub Super-Linter](https://github.com/rssnyder/discord-stock-tickerboard/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)

[![GitHub last commit](https://img.shields.io/github/last-commit/rssnyder/discord-stock-tickerboard.svg?style=flat)]()
[![GitHub stars](https://img.shields.io/github/stars/rssnyder/discord-stock-tickerboard.svg?style=social&label=Star)]()
[![GitHub watchers](https://img.shields.io/github/watchers/rssnyder/discord-stock-tickerboard.svg?style=social&label=Watch)]()

## Support
<a href='https://ko-fi.com/rileysnyder' target='_blank'><img height='35' style='border:0px;height:46px;' src='https://az743702.vo.msecnd.net/cdn/kofi3.png?v=0' border='0' alt='Buy Me a Coffee' />
[![Discord Chat](https://img.shields.io/discord/806606291798982678)](https://discord.gg/CQqnCYEtG7)

Live stock ticker boards for your discord server.

![Discord Board Sidebar w/ Bots](https://s3.cloud.rileysnyder.org/public/assets/sidebar-board.gif)

Love these bots? You can support this project by subscribing to the premium version or maybe [buy me a coffee](https://ko-fi.com/rileysnyder) or [hire me](https://github.com/rssnyder) to write/host **your** discord bot!

Related Projects:

Single Stock Ticker Bots: https://github.com/rssnyder/discord-stock-ticker

## Add free boards to your servers (click the image to add)

Don't see a board that you need or want to change a current list? Open a github issue/pr or join our discord server!

[![WSB](https://s3.cloud.rileysnyder.org/public/assets/wsb.jpeg)](https://discord.com/api/oauth2/authorize?client_id=828687744640352347&permissions=0&scope=bot)
[![FAANG](https://s3.cloud.rileysnyder.org/public/assets/faang.jpeg)](https://discord.com/api/oauth2/authorize?client_id=828687783190331392&permissions=0&scope=bot)
[![CC](https://s3.cloud.rileysnyder.org/public/assets/cc.jpeg)](https://discord.com/api/oauth2/authorize?client_id=828687818229284874&permissions=0&scope=bot)

## Premium

For advanced features like faster update times, custom stock lists, and having the price in the Name vs the Activity, you can subscribe to my premuim offering. I will host individual instances for your discord server at a cost of $2/month.

![Discord Board Premium Sidebar w/ Bots](https://s3.cloud.rileysnyder.org/public/assets/sidebar-board-premium.gif)

There is a full logging stack that includes loki & promtail with grafana for visualization.

If you are interested please see the [contact info on my github page](https://github.com/rssnyder) and send me a messgae via your platform of choice (discord perferred). For a live demo, join the support discord linked at the top or bottom of this page.

![Really cool grafana dashboard](https://s3.cloud.rileysnyder.org/public/assets/grafana.png)

### Self-Hosting

#### Running in a simple shell

Pull down the latest release for your OS [here](https://github.com/rssnyder/discord-stock-tickerboard/releases).

```
wget https://github.com/rssnyder/discord-stock-tickerboard/releases/download/v2.0.0/discord-stock-tickerboard-v2.0.0-linux-amd64.tar.gz

tar zxf discord-stock-tickerboard-v2.0.0-linux-amd64.tar.gz

./discord-stock-tickerboard
```

Add board to the service with an HTTP call:


Stock Payload: 

```
{
  "name": "Stocks",
  "frequency": 3,
  "items": ["PFG", "GME", "AMC"],
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "header": "1. ",  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "set_nickname": true,  # bool/OPTIONAL: put prices in nickname
  "set_color": true,  # bool/OPTIONAL: requires set_nickname
  "arrows": true  # bool/OPTIONAL: show arrows in ticker names
}
```


Crypto Payload: 

```
{
  "name": "Cryptos",
  "crypto": true,
  "frequency": 3,
  "items": ["bitcoin", "ethereum", "dogecoin"],
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "header": "2. ",  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "set_nickname": true,  # bool/OPTIONAL: put prices in nickname
  "set_color": true,  # bool/OPTIONAL: requires set_nickname
  "arrows": true  # bool/OPTIONAL: show arrows in ticker names
}
```

```
curl -X POST -H "Content-Type: application/json" --data '{
  "name": "Stocks",
  "frequency": 3,
  "set_nickname": true,
  "set_color": true,
  "percentage": true,
  "arrows": true,
  "discord_bot_token": "xxxxxxx",
  "items": ["PFG", "GME", "AMC"]
}' localhost:8080/tickerboard

```

#### List current running bots

```
curl localhost:8080/tickerboard
```

#### Remove a bot

```
curl -X DELETE localhost:8080/tickerboard/Stocks
```

```
curl -X DELETE localhost:8080/tickerboard/Cryptos
```

## Support

If you have a request for a new ticker or issues with a current one, please open a github issue or find me on discord at `jonesbooned#1111` or [join the support server](https://discord.gg/CQqnCYEtG7).

Love these bots? Maybe [buy me a coffee](https://ko-fi.com/rileysnyder)! Or send some crypto to help keep these bots running:

eth: 0x27B6896cC68838bc8adE6407C8283a214ecD4ffE

doge: DTWkUvFakt12yUEssTbdCe2R7TepExBA2G

bch: qrnmprfh5e77lzdpalczdu839uhvrravlvfr5nwupr

btc: 1N84bLSVKPZBHKYjHp8QtvPgRJfRbtNKHQ
