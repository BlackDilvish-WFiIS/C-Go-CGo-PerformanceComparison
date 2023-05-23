FROM golang:latest

RUN apt-get update && \
    apt-get install -y build-essential

WORKDIR /app

COPY . .

CMD ["./run.sh"]
