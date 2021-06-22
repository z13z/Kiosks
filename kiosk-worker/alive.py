import logging
import os.path
import schedule
import time
import threading

import requests

import ip_provider

SERVER_ADDRESS = "http://{}:8080".format(os.getenv("KIOSK_SERVER", "localhost"))
CONNECTOR_SERVICE_ADDRESS = "/kiosksConnector"
AUTHENTICATION_HEADER_KEY = "Authentication"
SERVICE_CALL_INTERVAL_IN_SECONDS = 30
SENT_FROM_IP_HEADER_KEY = 'X-From-Ip'


def worker_job(jwt, controller_service_port):
    schedule.every(SERVICE_CALL_INTERVAL_IN_SECONDS).seconds.do(call_create_method, jwt=jwt,
                                                                controller_service_port=controller_service_port)
    while True:
        schedule.run_pending()
        time.sleep(1)


def start_status_update_worker(jwt, controller_service_port):
    threading.Thread(target=worker_job, args=(jwt, controller_service_port)).start()


def call_create_method(jwt, controller_service_port):
    try:
        session = requests.Session()
        session.headers.update({AUTHENTICATION_HEADER_KEY: jwt})
        session.headers.update({SENT_FROM_IP_HEADER_KEY: ip_provider.get_ip() + ":" + str(controller_service_port)})
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
