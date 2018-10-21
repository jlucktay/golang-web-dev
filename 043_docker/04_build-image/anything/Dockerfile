# Changes whalesay to speak a fortune
FROM docker/whalesay:latest
RUN apt-get -y update \
    && apt-get install -y fortunes=1.99.1-7build1 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
CMD /usr/games/fortune -a | cowsay
