# Slack Stikers
![Banner](.github/files/sticker.gif)

Your slack conversations should not be boring. Add stickers to spice things up.

### What this app does
- Consumes the slack APIs to make the app a reality
- Expose endpoints(webhooks) for slack to send us requests
- Listen to those requests from slack and carry out the necessary actions
- Query the DB to get related stickers
- Return stickers that match to the user
- Send message to slack channel or private chat
- Interact with cloudinary
- Expose APIs for the FE to upload to cloudinary
- Many more

How does it work:
* First, add the slack app to your workspace
* Type `/stickers` or use the plus button shortcut to find a stickers
* Put in your search keyword
* Send that sticker!!!


### Built With

* [Goland](https://go.dev/)
* [Slack API](https://api.slack.com/)
* [Vue.js](https://vuejs.org/)
* [Cloudinary](https://cloudinary.com/)


## Getting Started

_Below is an example of how you can install and set up your app.

1. Get the .env file from [https://github.com/odetolakehinde](https://github.com/odetolakehinde)
2. Clone the repo
   ```sh
   git clone https://github.com/odetolakehinde/slack-stickers-be.git
   ```
##### Things to know and to get started with it's engineering.

- This project runs with a docker. You only need to execute `docker-compose up` from your favorite terminal. Keep in mind that the first time, it will run for a while and download some stuff, mafo!
- Golang is used with its ninja frameworks such gin framework for http handling. APIs are in rest API (gin)

##### So, welcome to Go!.

The main source code is in the `src` directory. Don't be scared!<br/>
###### Once docker is running, you can always access the app via the follow:<br/>

- Health check: `http://localhost:6001`
- Rest API endpoint: `http://localhost:6001/api/v1`

## Roadmap
See the [open issues](https://github.com/odetolakehinde/slack-stickers-be/issues) for a full list of proposed features (and known issues).


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


## Contact

Your Name - [@slackstickers](https://twitter.com/slackstickers) - useslackstickers@gmail.com

Project Link: [https://github.com/odetolakehinde/slack-stickers-be](https://github.com/odetolakehinde/slack-stickers-be)


## Support Us

* [Buy me a coffee](https://buymeacoffee.com/slackstickers)
