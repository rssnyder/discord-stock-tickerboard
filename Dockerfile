FROM python:3.8.4
LABEL org.opencontainers.image.source https://github.com/rssnyder/discord-stock-tickerboard

COPY requirements.txt /

RUN pip install -r requirements.txt

COPY main.py /
COPY builtin.json /
COPY utils /utils

CMD [ "python", "./main.py" ]
