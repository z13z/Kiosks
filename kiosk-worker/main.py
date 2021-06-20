#!/usr/bin/python3
# pip install requests

import requests
import os.path
import logging

SERVER_ADDRESS = "http://localhost:8080"
CONNECTOR_SERVICE_ADDRESS = "/kiosksConnector"
CONFIG_FILE_NAME = "kiosk_config.json"


def call_create_method():
    response = requests.put(SERVER_ADDRESS + CONNECTOR_SERVICE_ADDRESS,
                            json={"kioskImageId": 1, "kioskAddress": "localhost:10013"})
    f = open(CONFIG_FILE_NAME, "w")
    logging.info("Got config from server: " + response.json())
    f.write(response.json())
    f.close()


def config_file_exists():
    return os.path.isfile(CONFIG_FILE_NAME)


if __name__ == '__main__':
    if not config_file_exists():
        call_create_method()
