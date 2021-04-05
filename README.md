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

[![WSB](https://logo.clearbit.com/gamestop.com)](https://discord.com/api/oauth2/authorize?client_id=828417475624435712&permissions=0&scope=bot)
[![FAANG](https://logo.clearbit.com/gamestop.com)](https://discord.com/api/oauth2/authorize?client_id=828417475624435712&permissions=0&scope=bot)

## Premium

For advanced features like faster update times, custom stock lists, and having the price in the Name vs the Activity, you can subscribe to my premuim offering. I will host individual instances for your discord server at a cost of $2/month.

There is a full logging stack that includes loki & promtail with grafana for visualization.

If you are interested please see the [contact info on my github page](https://github.com/rssnyder) and send me a messgae via your platform of choice (discord perferred). For a live demo, join the support discord linked at the top or bottom of this page.

![Really cool grafana dashboard](https://s3.cloud.rileysnyder.org/public/assets/grafana.png)

### Self-Hosting

Install python for your target operating system.

Clone down the repo locally:

```
git clone git@github.com:rssnyder/discord-stock-tickerboard.git && cd discord-stock-tickerboard
```

Register a new application in the discord developer portal and copy the bot token:

```
export DISCORD_BOT_TOKEN=<token>
```

Set your desired list source.

For custom on-the-fly lists:

```
export STATIC_LIST=PFG,GME,BB
```

Or for built in lists that come with the repo (sourced from `builtin.json`):

```
export BUILTIN_LIST=FAANG
```

To change the nickname of the bot instead of the activity (will need change nickname permissions in the server):

```
export SET_NICKNAME=1
```

If you want to adjust the time between the different stocks:

```
export FREQUENCY=3
```

Or if you want to set a static header in front of the bot name:

```
export NICKNAME_HEADER=1.
```

Other options:

```
export LOG_FILE=log.log  # log to file instead of stdout
```

Once all your options are set, simply install the dependencies and run the bot (virtual environments might be a smart idea):

```
pip3 install -r requirements.txt
python3 main.py
```

### Docker

You can also run these bots using docker. This can make running multiple bots esier. Here is an example docker compose file for the basic feature set (please check for the latest release and update the tags accordingly):

```
---
version: "2"
services:
  tickerboard-custom:
    image: ghcr.io/rssnyder/discord-stock-tickerboard:0.1.0
    container_name: tickerboard-custom
    environment:
      - DISCORD_BOT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
      - STATIC_LIST=PFG,GME,BB
      - SET_NICKNAME=1
    restart: unless-stopped

  tickerboard-meme:
    image: ghcr.io/rssnyder/discord-stock-tickerboard:0.1.0
    container_name: tickerboard-meme
    environment:
      - DISCORD_BOT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
      - BUILTIN_LIST=MEME
      - SET_NICKNAME=1
    restart: unless-stopped
```

```
docker-compose-up -d
```

## Support

If you have a request for a new ticker or issues with a current one, please open a github issue or find me on discord at `jonesbooned#1111` or [join the support server](https://discord.gg/CQqnCYEtG7).

Love these bots? Maybe [buy me a coffee](https://ko-fi.com/rileysnyder)! Or send some crypto to help keep these bots running:

eth: 0x27B6896cC68838bc8adE6407C8283a214ecD4ffE

doge: DTWkUvFakt12yUEssTbdCe2R7TepExBA2G

bch: qrnmprfh5e77lzdpalczdu839uhvrravlvfr5nwupr

btc: 1N84bLSVKPZBHKYjHp8QtvPgRJfRbtNKHQ
