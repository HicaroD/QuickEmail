# QuickEmail
A command-line interface for sendind emails quickly and easily. This software abstracts the connection to the Gmail SMTP server to make your life easier.

| Feature                                    | Status             |
|--------------------------------------------|--------------------|
| Add support to Gmail                       | :heavy_check_mark: |
| Add support to Outlook                     | :x:                |
| Add support to email with file attachments | :x:                |

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

QuickEmail is currently avaiable only for Gmail, but I have plans to extend it to Outlook in the future as well.

To send an e-mail, you can use this simple command below:
```
~/go/bin/QuickEmail from "Hicaro" -topic "The topic of your e-mail" -send "Your message" -to "Recipient's email"
```

After run this command, the program will ask you two things: e-mail and password. Your password will not be showed in the console while you type. After that, please check your "Sent" e-mails to make sure that everything was sent successfully.

## Help

If you want to know how these flags work. You can run `./QuickEmail -h`. That's gonna be the output.

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
