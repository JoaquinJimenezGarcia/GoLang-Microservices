FROM alpine:latest

RUN mkdir /app
COPY . /app

CMD [ "/app/loggerApp" ]