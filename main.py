'''discord-stock-tickerboard'''
from os import getenv
from json import load
import logging

import asyncio
import discord
from redis import Redis, exceptions

from utils.yahoo import get_stock_price

CURRENCY = 'usd'
BUILTIN_FILENAME = 'builtin.json'


class TickerBoard(discord.Client):
    '''
    Discord client for watching stock/crypto prices
    '''


    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        # Get source of securities
        dynamic_list = getenv('DYNAMIC_LIST')
        static_list = getenv('STATIC_LIST')
        builtin_list = getenv('BUILTIN_LIST')

        if dynamic_list:
            logging.info('dynamic list')
            return

        elif static_list:
            logging.info('static list')

            self.bg_task = self.loop.create_task(
                self.static_list_board(
                    static_list.split(','),
                    getenv('LIST_NAME', 'UNKNOWN'),
                    getenv('SET_NICKNAME'),
                    getenv('NICKNAME_HEADER', ''),
                    getenv('FREQUENCY', 30)
                )
            )
        
        elif builtin_list:
            logging.info('built in list')

            # load builtin list
            with open(BUILTIN_FILENAME, 'r') as json_file:
                data = load(json_file)
            
            if builtin_list not in data:
                logging.error(f'{builtin_list} not an option in {BUILTIN_FILENAME}')
                return

            self.bg_task = self.loop.create_task(
                self.static_list_board(
                    data[builtin_list],
                    getenv('LIST_NAME', builtin_list),
                    getenv('SET_NICKNAME'),
                    getenv('NICKNAME_HEADER', ''),
                    getenv('FREQUENCY', 30)
                )
            )


    async def on_ready(self):
        '''
        Log that we have successfully connected
        '''

        logging.info('logged in')

        # We want to know some stats
        servers = [x.name for x in list(self.guilds)]

        redis_server = getenv('REDIS_URL')
        if redis_server:

            # Use redis to store stats
            r = Redis(host=redis_server, port=6379, db=0)

            try:
                for server in servers:
                    r.incr(server)
            except exceptions.ConnectionError:
                logging.info('No redis server found, not storing stats')

        logging.info('servers: ' + str(servers))


    async def static_list_board(self, symbols: list, name: str, change_nick: bool = False, nick_header: str = '', frequency: int = 60):
        '''
        Update the bot activity based on stock price
        symbols = list of symbols to cycle through
        name = name of the list
        change_nick = flag for changing nickname vs activity
        frequency = how often to cycle through
        '''

        await self.wait_until_ready()
        logging.info('starting static list board job...')

        # Loop as long as the bot is running
        index = 0
        while not self.is_closed():

            # Go back to start
            if index == len(symbols):
                index = 0
            
            target = symbols[index]

            logging.info(f'fetching stock price of {target}...')
            
            # Grab the current price data w/ day difference
            data = get_stock_price(target)
            price_data = data.get('quoteSummary', {}).get('result', []).pop().get('price', {})
            price = price_data.get('regularMarketPrice', {}).get('raw', 0.00)
            price = "{:.2f}".format(price)

            # If after hours, get change
            if price_data.get('postMarketChange'):

                # Get difference or new price
                if getenv('POST_MARKET_PRICE'):
                    post_market_target = 'postMarketPrice'
                else:
                    post_market_target = 'postMarketChange'

                raw_diff = price_data.get(post_market_target, {}).get('raw', 0.00)
                diff = round(raw_diff, 2)

                activity_content = f'${price} AHT {diff}'
                logging.info(f'stock after hours price retrived: {activity_content}')
            else:
                raw_diff = price_data.get('regularMarketChange', {}).get('raw', 0.00)
                diff = round(raw_diff, 2)
                if diff >= 0.0:
                    diff = '+' + str(diff)

                activity_content = f'${price} / {diff}'
                logging.info(f'stock price retrived: {activity_content}')
            
            # Add name to content
            activity_content = target + ' ' + activity_content

            # Change name via nickname if set
            if change_nick:

                for server in self.guilds:

                    try:
                        await server.me.edit(
                            nick=' '.join([nick_header, activity_content])
                        )

                    except discord.HTTPException as e:
                        logging.error(f'updating nick failed: {e.status}: {e.text}')
                    except discord.Forbidden as f:
                        logging.error(f'lacking perms for chaning nick: {f.status}: {f.text}')

                    logging.info(f'stock updated nick in {server.name}')
                
                # Use activity for name
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=name
                        )
                    )

                    logging.info(f'stock activity updated: {name}')

                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

            else:

                # Use activity for stock prices
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=activity_content
                        )
                    )

                    logging.info(f'stock activity updated: {activity_content}')

                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

            logging.info(f'stock sleeping for {frequency}s')
            await asyncio.sleep(int(frequency))
            logging.info('stock sleep ended')
            index += 1


if __name__ == "__main__":

    logging.basicConfig(
        filename=getenv('LOG_FILE'),
        level=logging.INFO,
        datefmt='%Y-%m-%d %H:%M:%S',
        format='%(asctime)s %(levelname)-8s %(message)s',
    )

    token = getenv('DISCORD_BOT_TOKEN')
    if not token:
        logging.error('DISCORD_BOT_TOKEN not set!')

    client = TickerBoard()
    client.run(token)
