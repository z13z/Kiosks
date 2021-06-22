#!/bin/python3
# pip install requests
# pip install Flask
# sudo apt install -y scrot
from flask import request
import common
import configs
import flask

import controller
import alive

api = flask.Flask(__name__)

AUTHENTICATION_HEADER = 'Authentication'
CONTROLLER_SERVICE_PORT = 5000


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
        kioskId, serverPassword, servicePassword = configs.call_create_method(CONTROLLER_SERVICE_PORT)
    else:
        kioskId, serverPassword, servicePassword = configs.load_configs_from_file()
    alive.start_status_update_worker(serverPassword, CONTROLLER_SERVICE_PORT)
    api.run(host="0.0.0.0", port=CONTROLLER_SERVICE_PORT)
