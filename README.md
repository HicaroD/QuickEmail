# QuickEmail
A command-line interface for sendind emails quickly and easily. This software abstracts the connection to the Gmail SMTP server to make your life easier.

| Feature                                                  | Status             |
|----------------------------------------------------------|--------------------|
| **Support for Gmail**                                    | :heavy_check_mark: |
| **Support for sending an e-mail to multiple recipients** | :heavy_check_mark: |


## Summary
1. [Requirements](#requirements)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Help](#help)
5. [FAQ (Frequently Asked Questions)](#faq)
6. [License](#license)


## Requirements

1. [Golang@1.17.6](https://go.dev/dl/) or more recent.

## Installation

1. Install the latest version of the program in `usr/bin`

    ```bash
    GOBIN=/usr/bin/ go install github.com/HicaroD/QuickEmail@latest
    ```

    After that, you can simply run the program using `QuickEmail` in the command-line directly :smiley:.

    If you're having problems with permission, try adding `sudo` at the beginning of the command. Now you can run it again and everything should work properly.

## Usage

First of all, you need to allow less secure app in your Gmail configuration. In order to do that, click [here](https://myaccount.google.com/), click on "Security" in the left side, scroll down until "Less secure app access". After you find it, click on it and enable that feature. Otherwise, Gmail wouldn't allow you to send e-mail using QuickEmail. 

After run any command, the program will ask you two things: e-mail and password. Your password will not be showed in the console while you type. After that, if you see "E-mail was sent successfully!", everything worked fine, but please check your "Sent" e-mails to make sure.

1. **Send a simple e-mail**

    ```
    QuickEmail -from "Your name" -topic "The topic of your e-mail" -send "The message" -to "Recipient's email"
    ```

2. **Send an e-mail to multiple recipients**

    The command is basically the same, but in the `to` flag, you have to separate the recipient's e-mails with a semicolon.

    ```
    ./QuickEmail -from "Your name" -topic "The topic of your e-mail" -send "The message" -to "example1@gmail.com;example2@gmail.com"
    ```

## Help

If you want to know how these flags work, run `QuickEmail -h`.

```
Usage of QuickEmail:
  -from string
    	Your username
  -send string
    	The actual message that you want to send
  -to string
    	The recipient's e-mail
  -topic string
    	The topic of the e-mail
```

## FAQ

1. **Where is Outlook support?**

    As far as I can say, Outlook doesn't allow me to send e-mail easily. I'll be working on that and do my best to implement it, but I might give up on it. We'll see.

2. **Does QuickEmail supports file attachments?**

    It currently doesn't support this feature, maybe in the future it will! 

## License
[MIT](./LICENSE)
