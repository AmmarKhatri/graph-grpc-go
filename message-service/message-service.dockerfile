FROM alpine:latest

RUN mkdir /app

COPY messageApp /app

CMD [ "/app/messageApp"]