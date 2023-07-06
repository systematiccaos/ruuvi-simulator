FROM alpine:latest

RUN apk add gcompat

WORKDIR /app
RUN adduser --home /app --disabled-password app
ADD sim .
RUN chmod +x /app/sim
RUN chown -R app /app
USER app

CMD /app/sim