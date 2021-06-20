import socket
import requests
import os

LOCAL_IP_ENV_KEY = "LOCAL_MACHINE"
ADDRESS_TO_GET_PUBLIC_IP = "https://api.ipify.org/"
LOCAL_IP_ADDRESS_VALUE = None


# todo zaza note.
# this code only for testing on machines with non public ip
# code is taken from https://stackoverflow.com/a/23822431
def get_local_ip():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect(('1.1.1.1', 0))
    return s.getsockname()[0]


def get_public_ip():
    return requests.get(ADDRESS_TO_GET_PUBLIC_IP).text


def get_ip():
    if os.getenv(LOCAL_IP_ENV_KEY) is not None and os.getenv(LOCAL_IP_ENV_KEY) == "true":
        return get_local_ip()
    else:
        return get_public_ip()
