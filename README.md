NBIS chatbot demo
=================

This repository contains getting-started code for working with slack chatbots
for the NBIS developers' mini-retreat 2022.

## Creating and setting up your bot

You can go to [this page](https://api.slack.com/apps?new_app=1) to create a new
bot.

### Basic Information

You need to create a token under **App-Level Tokens** to use socket mode for
your bot. Add both scopes.

Under **Display Information** you can change your bot name, add a description,
and set app icon. Only the name is mandatory.

Then go to the following pages to configure your bot for some basic tasks:

### Socket Mode

Enable socket mode to use the examples in this repo.

### Slash Commands

Here you can add new commands for your bot to react to. The demo bot reacts to
`/calm`, but you can change this easily.

You have to add all commands you want to use in the app using the **Create New
Command** button.

### Event Subscriptions

You need:
 - `message.im` to chat with your bot
 - `app_mention` to be able to @ your bot

NOTE: Bots need to be explicitly added to a channel to react to mentions. An
easy way to add bots to a channel is to write their @name, then clicking the
name and selecting "Add this app to a channel".

### OAuth & Permissions

You need to add some permissions under **Bot Token Scopes**:
 - `app_mentions:read` for the `app_mention` event
 - `chat:write` to reply to messages to your bot
 - `commands` to use slash commands
 - `im:history` to see messages in chats

## Running your bot

The bots reads token information from environment variables. You will need:
 - `SLACK_BOT_TOKEN`, found under **Install App** (starts with `xoxb-`).
 - `SLACK_APP_TOKEN`, found under **Basic Information** -> **App-Level Tokens**.
   Click the token name to see the token value (starts with `xapp-`).

### Golang

Make sure that you have exported the needed environment variables mentioned
above.

First, go to the `golang` directory and install the dependencies:
```
go get
```

Then start the bot as:
```
go run .
```
or
```
go build
./chatbot
```

The golang chatbot has a lot more debug messages than the other two, and is the
only one that explicitly tells you if if connected successfully.

### Javascript

Make sure that you have exported the needed environment variables mentioned
above.

First, go to the `javascript` directory and install the dependencies:
```
npm install
```

Then start the bot as:
```
node chatbot.js
```

If your tokens are correct, you should see:
```
⚡️ Bolt app is running!
```

You can now test talking to your bot on slack!


### Python

Make sure that you have exported the needed environment variables mentioned
above.

First, go to the `python` directory, create and enable a virtual environment,
and install the dependencies:
```
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

Then start the bot as:
```
python3 chatbot.py
```

If your tokens are correct, you should see:
```
⚡️ Bolt app is running!
```

You can now test talking to your bot on slack!
