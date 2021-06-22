import logging
import os.path
import schedule
import time
import threading

import requests

SERVER_ADDRESS = "http://{}:8080".format(os.getenv("KIOSK_SERVER", "localhost"))
CONNECTOR_SERVICE_ADDRESS = "/kiosksConnector"
AUTHENTICATION_HEADER_KEY = "Authentication"
SERVICE_CALL_INTERVAL_IN_SECONDS = 30


def worker_job(jwt):
    schedule.every(SERVICE_CALL_INTERVAL_IN_SECONDS).seconds.do(call_create_method, jwt=jwt)
    while True:
        schedule.run_pending()
        time.sleep(1)


def start_status_update_worker(jwt):
    threading.Thread(target=worker_job, args=(jwt,)).start()


def call_create_method(jwt):
    try:
        session = requests.Session()
        session.headers.update({AUTHENTICATION_HEADER_KEY: jwt})
        response = session.post(SERVER_ADDRESS + CONNECTOR_SERVICE_ADDRESS)
        if response.status_code != 202:
            logging.warning(
                "Error status code returned while updating last online time, status code {}".format(
                    response.status_code))
        else:
            logging.info("Updating last online time finished successfully")
        session.close()
    except requests.exceptions.ConnectionError:
        logging.warning("Connection error while updating last online time")
