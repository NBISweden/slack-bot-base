
// Simple SLACK bot base, based on the bolt-js tutorial at:
// https://slack.dev/bolt-js/tutorial/getting-started

const { App } = require('@slack/bolt');

const app = new App({
  token: process.env.SLACK_BOT_TOKEN,
  socketMode: true,
  appToken: process.env.SLACK_APP_TOKEN,
});

/**
 * Replies to the /calm slash command.
 */
app.command('/calm', async ({ command, ack, respond }) => {
  // Acknowledge command request
  await ack();

  await respond(`Thanks <@${command.user_id}>. I feel better now!`);
});

/**
 * Replies to the `app_mention` event. This is triggered when some one mentions
 * the bot by @<name>.
 */
app.event('app_mention', async ({ event, client, logger }) => {
  try {
    // Call chat.postMessage with the built-in client
    const result = await client.chat.postMessage({
      channel: event.channel,
      text: `Hi <@${event.user}>!`
    });
    logger.info(result);
  }
  catch (error) {
    logger.error(error);
  }
});

/**
 * Replies to the `message` event. This is triggered when the bot gets a direct
 * chat message.
 *
 */
app.message('', async ({ message, say }) => {
  // say() sends a message to the channel where the event was triggered
  await say(`Haha, yeah.`);
});


(async () => {
  // Start your app
  await app.start();

  console.log('⚡️ Bolt app is running!');
})();

