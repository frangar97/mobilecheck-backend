FROM alpine:latest

RUN mkdir /app

COPY impulsadoraApp /app

CMD ["/app/impulsadoraApp"]