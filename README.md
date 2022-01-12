# QuickEmail
A command-line interface for sendind emails (Gmail) quick and easily. This software abstracts the connection to the Gmail SMTP server to make your life easier.

## Requirements

1. [Golang@1.17.6](https://go.dev/dl/) or more recent.

## Installation

1. Install the latest version of the program

```bash
go install github.com/HicaroD/QuickEmail@latest
```

This will add your program to the `~/go/bin/` folder. You can run the program with `~/go/bin/QuickEmail`

## Usage

You can run this command to send an e-mail.

```
~/go/bin/QuickEmail -topic "The topic of your e-mail" -send "Your message" -to "Email recipient"
```

After run this command, the program will ask you two things: e-mail and password. Your password will not be showed in the console while you type!

## License
[MIT](./LICENSE)
