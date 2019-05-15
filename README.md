# <B>Feedback Bot</B>

<img width="647" alt="screenshot 2019-03-07 19 00 26" src="https://user-images.githubusercontent.com/29664539/54006079-c3192d00-4110-11e9-8cfe-b54dfe74c3c8.png">

Feedback Bot is an automated way of delivering and receiving one-to-one feedback over Slack.

The feedback format is encouraging and at the same time constructive (Good-Better-Best).

## <B>Usage:</B>

- #### Send feedback: Generate a feedback survey and send feedback to a user in your slack.
<img width="511" alt="screenshot 2019-03-07 19 42 28" src="https://user-images.githubusercontent.com/29664539/54006211-3327b300-4111-11e9-8356-c168d2b0d998.png">

- #### Receive all your feedback weekly or on request in CSV format.
<img width="656" alt="screenshot 2019-03-07 19 42 10" src="https://user-images.githubusercontent.com/29664539/54006212-3622a380-4111-11e9-948b-aab1c3821b34.png">

- #### Chat with the bot and do things such as:
  - #### Subscribe and unsubscribe from weekly feedback notifications.
  - #### Query feedback feedback by one or multiple parameters.
  - #### Delete feedback you have received.

## <B>Installation</B>
#### 1. Go to https://api.slack.com/apps and create a new App
#### 2. Enable `Incoming Webhooks` and `Interactive Components`
#### 3. Fork and clone the project
#### 4. Create a PostgresQL database and name it `feedback-bot-db`
#### 5. Create an .env file in the root of your project, with the following environment variables:
```
BOT_OAUTH_ACCESS_TOKEN=xoxb-11439472923-880521654032-Sx7Qv46ofo9XODBqAc9pQ5Cl
VERIFICATION_TOKEN=9BP11qDFQwPP6seaZQeE9QLC
DBUSER=database-user
DBPASSWORD=database-password
DBHOST=database-host
DBPORT=5432
PORT=3000
DBNAME=database-name
```
You can get the BOT_OAUTH_ACCESS_TOKEN and VERIFICATION_TOKEN from the Slack Dashboard
#### 6 Run the server
- ##### Option 1: Deploy to Heroku
- ##### Option 2: Or run using Docker `docker-compose up`
- ##### Then Expose the local URL using Ngrok
            
#### 7. Go to your Apps `Interactive Components` section and set the request URL to: `your-app-url/events-endpoint`

## <B>More</B>
#### If you want to change the feedback type options in the surveys:
1. Navigate to `utils/attachments.go`
2. Navigate to the `GenerateFeedbackSurvey()`
3. In `dialogElement2` change the options to your liking.





