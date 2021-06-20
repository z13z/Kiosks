import logging
import os
import subprocess

import flask

LATEST_SCREENSHOT_FILE_NAME = "screenshot.png"


def get_screenshot():
    os.system("export DISPLAY=:1 && scrot -z -o {} -e 'mv $f .'".format(LATEST_SCREENSHOT_FILE_NAME))
    return flask.send_file(LATEST_SCREENSHOT_FILE_NAME)


def execute_command(command):
    process = None
    try:
        process = subprocess.Popen(command.split(), stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        stdout, stderr = process.communicate()
        return stdout
    except subprocess.CalledProcessError:
        logging.error("Error in command execution '{}'".format(command))
        return process.stderr.read()
    except FileNotFoundError:
        logging.error("Executable not found for command '{}'".format(command))
        return "Executable not found for command '{}'".format(command)
    except:
        logging.error("Error while executing command '{}'".format(command))
        return "Error while executing command '{}'".format(command)
