FROM debian:12-slim
WORKDIR /app

COPY v2ray-plugin /usr/bin/
COPY sslocal /usr/bin/
COPY client.json /app

CMD ["/usr/bin/sslocal", "-c", "/app/client.json"]

