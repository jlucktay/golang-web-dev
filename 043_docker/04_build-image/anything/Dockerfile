# Changes whalesay to speak a fortune
FROM docker/whalesay:latest
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
    fortune-mod=1:1.99.1-7 \
    fortunes=1:1.99.1-7 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
CMD ["bash", "-c", "/usr/games/fortune -a | cowsay"]
