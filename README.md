# QuickEmail
A command-line interface for sendind emails (Gmail) quickly and easily. This software abstracts the connection to the Gmail SMTP server to make your life easier.

- Progress:

| Feature                                   | Status      |
|-------------------------------------------|-------------|
| Add support to Gmail                      | Done        |
| Add support to Outlook                    | In progress |
| Add support to email with file attachment | In progress |

## Requirements

1. [Golang@1.17.6](https://go.dev/dl/) or more recent.

## Installation

1. Install the latest version of the program

```bash
go install github.com/HicaroD/QuickEmail@latest
```

This will add your program to the `~/go/bin/` folder. You can run the program with `~/go/bin/QuickEmail`

## Usage

First of all, you need to allow less secure app in your Gmail configuration. In order to do that, click [here](https://myaccount.google.com/), click on "Security" in the left side, scroll down until "Less secure app access". After you find it, click on it and enable that feature. Otherwise, Gmail wouldn't allow you to send e-mail using QuickEmail.

Now you can run this command to send an e-mail.

```
~/go/bin/QuickEmail -topic "The topic of your e-mail" -send "Your message" -to "Email recipient"
```

After run this command, the program will ask you two things: e-mail and password. Your password will not be showed in the console while you type!

## License
[MIT](./LICENSE)
