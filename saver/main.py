import paho.mqtt.client as mqtt
import os

def on_message(client, userdata, message):
    print(f"{message.topic}: {message.payload.decode()}")

host = os.getenv("HOST")
assert host is not None

if __name__ == "__main__":
    client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    client.on_message = on_message
    client.connect(host, 1883)
    client.subscribe("home")
    client.loop_forever()
