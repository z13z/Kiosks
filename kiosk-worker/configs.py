import random
import string
import json
import common
import requests
import os.path
import logging

# todo zaza change address
SERVER_ADDRESS = "http://localhost:8080"
CONNECTOR_SERVICE_ADDRESS = "/kiosksConnector"
CONFIG_FILE_NAME = "kiosk_config.json"
SERVICE_PASSWORD_LENGTH = 16


def get_config_variables_from_json(configs):
    if configs is None:
        exit(221)
    return configs["id"], configs["password"], configs["servicePassword"]


def call_create_method():
    possible_chars = string.ascii_letters + string.digits
    password_plaintext = "".join(random.choices(possible_chars, k=SERVICE_PASSWORD_LENGTH))
    response = requests.put(SERVER_ADDRESS + CONNECTOR_SERVICE_ADDRESS,
                            json={"kioskImageId": 1, "kioskAddress": "localhost:5000",
                                  "servicePassword": password_plaintext})
    must_save_config = None
    if response.status_code == 200:
        f = open(CONFIG_FILE_NAME, "w")
        must_save_config = response.json()
        logging.info("Got config from server: " + str(must_save_config))
        must_save_config["servicePassword"] = common.get_sha256(password_plaintext)
        config_json = json.dumps(must_save_config)
        f.write(config_json)
        f.close()
    return get_config_variables_from_json(must_save_config)


def config_file_exists():
    return os.path.isfile(CONFIG_FILE_NAME)


def load_configs_from_file():
    f = open(CONFIG_FILE_NAME, "r")
    json_data = f.read()
    f.close()
    return get_config_variables_from_json(json.loads(json_data))
