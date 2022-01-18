# QuickEmail
A command-line interface for sendind emails quickly and easily. This software abstracts the connection to the Gmail SMTP server to make your life easier.

| Feature                                                  | Status             |
|----------------------------------------------------------|--------------------|
| Support for Gmail                                        | :heavy_check_mark: |
| Support for sending an e-mail to multiple recipients     | :heavy_check_mark: |
| Support for Outlook                                      | :x:                |
| Support for file attachments                             | :x:                |

## Requirements

1. [Golang@1.17.6](https://go.dev/dl/) or more recent.

## Installation

1. Install the latest version of the program in `usr/bin`

```bash
GOBIN=/usr/bin/ go install github.com/HicaroD/QuickEmail@latest
```

After that, you can simply run the program using `QuickEmail` in the command-line directly :smiley:.

## Usage

First of all, you need to allow less secure app in your Gmail configuration. In order to do that, click [here](https://myaccount.google.com/), click on "Security" in the left side, scroll down until "Less secure app access". After you find it, click on it and enable that feature. Otherwise, Gmail wouldn't allow you to send e-mail using QuickEmail. 

After run any command, the program will ask you two things: e-mail and password. Your password will not be showed in the console while you type. After that, if you see "Operation successfully completed!", everything worked fine, but please check your "Sent" e-mails to make sure.

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

```bash
$ ./QuickEmail -h
Usage of ./QuickEmail:
  -from string
    	Your username
  -send string
    	The actual message that you want to send
  -to string
    	The recipient e-mail
  -topic string
    	The topic of the e-mail
```

## License
[MIT](./LICENSE)
