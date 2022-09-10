# Balance notifier application

This application is for tracking balance changes of 
account and send notification via channels. App integrated with provider banks open apis.
Currently, `Kapital` supported as a provider and `Telegram` as a channel.

# For running application

First you need to configure environment variables inside [`.env`](.env) file.

- `USERNAME` - is your bank username;
- `PASSWORD` - is your bank password;
- `PROVIDER.KAPITAL.ACCOUNT_NO` - if provider is Kapital then your account no which is needed for balance tracking;
- `NOTIFICATION.TELEGRAM.API_KEY` - is for sending notification via `telegram` you'll need api key. You can generate via [bot father](https://t.me/BotFather);
- `NOTIFICATION.TELEGRAM.CHAT_ID` - is for sending message to which chat/user. You can get chat id via [Raw data bot](https://t.me/RawDataBot). you'll neeed `chat.id` field.

After that run [`docker-compose.yml`](docker-compose.yml) file:
    
    docker-compose up -d

Application will build and start new docker container.

For stopping application run:

    docker-compose down --rmi all
