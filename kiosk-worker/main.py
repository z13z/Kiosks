#!/bin/python3
# pip install requests
# pip install Flask
# sudo apt install -y scrot
from flask import request
import common
import configs
import flask

import controller

api = flask.Flask(__name__)

AUTHENTICATION_HEADER = 'Authentication'


def check_authentication():
    if AUTHENTICATION_HEADER in request.headers:
        auth_header = request.headers[AUTHENTICATION_HEADER]
        return common.get_sha256(auth_header) == servicePassword
    return False


@api.route("/screenshot")
def get_screenshot():
    if not check_authentication():
        return "Unauthorized", 401
    return controller.get_screenshot(), 200


@api.route("/execute", methods=["POST"])
def execute_command():
    if not check_authentication():
        return "Unauthorized", 401
    return controller.execute_command(request.data.decode("utf-8")), 200


if __name__ == '__main__':
    global kioskId, serverPassword, servicePassword
    if not configs.config_file_exists():
        kioskId, serverPassword, servicePassword = configs.call_create_method()
    else:
        kioskId, serverPassword, servicePassword = configs.load_configs_from_file()
    api.run()
