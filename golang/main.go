package main

// Golang doesn't have an official SDK, but there are unofficial ones!
// This chatbot is based on a tutorial at:
// https://towardsdatascience.com/develop-a-slack-bot-using-golang-1025b3e606bc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// handleSlashCommand handles known slash commands.
// Currently it handles the /calm command.
func handleSlashCommand(command slack.SlashCommand, client *slack.Client) error {
	// We need to switch depending on the command
	if command.Command == "/calm" {

		msg := fmt.Sprintf("Thanks <@%s>. I feel better now!", command.UserName)

		_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionText(msg, true))
		if err != nil {

			return fmt.Errorf("failed to post message: %w", err)
		}
	}

	return nil
}

// handleEvent handles all types of bot events defined under
// "Event Subscriptions" in the slack API page.
func handleEvent(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	// First we check if this is an CallbackEvent
	case slackevents.CallbackEvent:

		innerEvent := event.InnerEvent
		// Finally check event type
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:

			// Mention events are when someone uses @<name> in chat

			msg := fmt.Sprintf("Hi <@%s>!", ev.User)
			_, _, err := client.PostMessage(ev.Channel, slack.MsgOptionText(msg, false))
			if err != nil {

				return fmt.Errorf("failed to post message: %w", err)
			}

		case *slackevents.MessageEvent:

			// Message events are direct messages to the bot

			// Since the bot replies are also posted in the chat, we need to
			// ignore them so we don't get into a reply-loop. Unless you
			// specifically wants bot chatting, this is easily done by just
			// removing all messages with a bot-id.
			if ev.BotID != "" {

				return nil
			}

			msg := "Haha, yeah..."

			_, _, err := client.PostMessage(ev.Channel, slack.MsgOptionText(msg, false))
			if err != nil {

				return fmt.Errorf("failed to post message: %w", err)
			}
		}
	default:

		return errors.New("unsupported event type")
	}

	return nil
}

func main() {

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	debugLog := true

	// Create a client with the bot and app tokens
	client := slack.New(
		botToken,
		slack.OptionDebug(debugLog),
		slack.OptionAppLevelToken(appToken),
	)

	// create a socket client with the regular client
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(debugLog),
		socketmode.OptionLog(
			log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags),
		),
	)

	// Create a context that can be used to cancel goroutine
	ctx, cancel := context.WithCancel(context.Background())
	// Make this cancel called properly in a real program , graceful shutdown etc
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		// runtime loop
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")

				return
			case event := <-socketClient.Events: // Check events

				switch event.Type {

				case socketmode.EventTypeSlashCommand:
					// Type cast to a SlashEvent
					command, ok := event.Data.(slack.SlashCommand)
					if !ok {
						log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)

						continue
					}
					// Acknowledge the request
					socketClient.Ack(*event.Request)

					err := handleSlashCommand(command, client)
					if err != nil {
						log.Fatal(err)
					}

				case socketmode.EventTypeEventsAPI:

					// Type cast to a EventAPI event
					eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)

						continue
					}
					// Send an Acknowledge to the slack server so that it knows
					// that we received the message
					socketClient.Ack(*event.Request)
					err := handleEvent(eventsAPIEvent, client)
					if err != nil {
						log.Fatal(err)
					}

				default:
					log.Printf("Unknown event type: %s", event.Type)
				}

			}
		}
	}(ctx, client, socketClient)

	err := socketClient.Run()
	if err != nil {
		log.Printf("failed to run socketClient: %v", err)
	}
}
