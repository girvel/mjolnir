import paho.mqtt.client as mqtt
import os
import psycopg
import logging
from datetime import datetime

def on_message(client, userdata, message):
    conn = userdata
    print(f"{message.topic}: {message.payload.decode()}")
    if message.topic.startswith("home/"):
        _, sensor, kind = message.topic.split("/")
        assert kind in ("temperature", "humidity")
        value = float(message.payload.decode())
        conn.execute(
            f"INSERT INTO {kind} (sensor, value, timestamp) VALUES (%s, %s, %s)",
            (sensor, value, datetime.now())
        )
        conn.commit()

postgres_host = os.getenv("POSTGRES_HOST")
assert postgres_host is not None

mosquitto_host = os.getenv("MOSQUITTO_HOST")
assert mosquitto_host is not None

if __name__ == "__main__":
    client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    client.on_message = on_message
    client.connect(mosquitto_host, 1883)
    client.subscribe("#")
    print(f"Connected to mosquitto on {mosquitto_host}:1883")

    with psycopg.connect(
        dbname="home",
        user="postgres",
        password="postgres",
        host=postgres_host,
        port=5432
    ) as conn:
        print(f"Connected to postgres on {postgres_host}:5432")

        client.user_data_set(conn)
        client.loop_forever()
