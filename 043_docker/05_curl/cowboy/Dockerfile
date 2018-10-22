# This docker file builds an image that runs curl
FROM ubuntu:bionic
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends curl=7.58.0-2ubuntu3.3 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
