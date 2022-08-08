FROM alpine:latest

RUN mkdir /app

COPY graphApp /app

CMD [ "/app/graphApp"]