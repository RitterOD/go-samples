#!/usr/bin/sh
git clone https://github.com/antlr/antlr4.git
cd antlr4/docker
docker build -t antlr/antlr4 --platform linux/amd64 .