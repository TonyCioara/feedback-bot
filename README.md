# Feedback Bot

### Feedback Bot is an automated way of delivering and receiving one-to-one feedback over Slack.
### The feedback format is encouraging and at the same time constructive (Good-Better-Best).

### <B>Usage:</B>
- Send feedback: Generate a feedback survey and send feedback to a user in your slack.
- Receive all your feedback weekly or on request in CSV format.
- Chat with the bot and do things such as:
  - Subscribe and unsubscribe from weekly feedback notifications.
  - Query feedback feedback by one or multiple parameters.
  - Delete feedback you have received.

### <B>Installation</B>
#### 1. Go to https://api.slack.com/apps and create a new App
#### 2. Create a PostgresQL database and name it `feedback-bot-db`
#### 3. Set your environment variables:
```
BOT_OAUTH_ACCESS_TOKEN=xoxb-11439472923-880521654032-Sx7Qv46ofo9XODBqAc9pQ5Cl
VERIFICATION_TOKEN=9BP11qDFQwPP6seaZQeE9QLC
DBUSER=database-user
DBPASSWORD=database-password
DBHOST=database-host
DBPORT=5432
DBNAME=database-name
```
You can get the BOT_OAUTH_ACCESS_TOKEN and VERIFICATION_TOKEN from the Slack Dashboard




