#!/usr/bin/env python3
"""
Simple SLACK bot base, based on the bolt-python tutorial at:
https://slack.dev/bolt-python/tutorial/getting-started.
"""

import os
import logging

from slack_bolt import App
from slack_bolt.adapter.socket_mode import SocketModeHandler


APP = App(
    token=os.environ.get("SLACK_BOT_TOKEN")
)

@APP.command("/calm")
def hello_command(ack, body):
    """
    Replies to the /calm slash command.
    """
    user_id = body["user_id"]
    ack(f"Thanks <@{user_id}>. I feel better now!")


@APP.event("app_mention")
def event_mention(event, say):
    """
    Replies to the `app_mention` event. This is triggered when some one mentions
    the bot by @<name>
    """
    say(f"Hi <@{event['user']}>!")


# pylint: disable=unused-argument
@APP.event("message")
def event_message(event, say):
    """
    Replies to the `message` event. This is triggered when the bot gets a direct
    chat message.

    More advanded examples of message events can be found at:
    https://github.com/slackapi/bolt-python/blob/main/examples/message_events.py
    """
    say(f"Haha, yeah.")


if __name__ == "__main__":

    logging.basicConfig(level=logging.INFO)
    SocketModeHandler(APP, os.environ["SLACK_APP_TOKEN"]).start()
