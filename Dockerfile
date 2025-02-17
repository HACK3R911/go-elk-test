FROM ubuntu:latest
LABEL authors="afana"

ENTRYPOINT ["top", "-b"]