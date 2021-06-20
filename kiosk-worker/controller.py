import os

import flask

LATEST_SCREENSHOT_FILE_NAME = "screenshot.png"


def get_screenshot():
    os.system("export DISPLAY=:1 && scrot -z -o {} -e 'mv $f .'".format(LATEST_SCREENSHOT_FILE_NAME))
    return flask.send_file(LATEST_SCREENSHOT_FILE_NAME)
