# sendem

A command-line tool for sending emails via SMTP.

## Installation

`sendem` will run on most Linux and Mac OS X systems.

To install it, just find the appropriate one-liner below - based on the destination O/S and architecture - and copy-paste it into your terminal.

Feel free to change the install dir - `$HOME/bin` in the examples below - to be something more appropriate for your needs.

### Linux (32-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/sendem/releases/latest/download/sendem-linux-386.tar.gz | tar -zxf - -C $HOME/bin
```

### Linux (64-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/sendem/releases/latest/download/sendem-linux-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Intel)

```
curl -s -L -o - https://github.com/alasdairmorris/sendem/releases/latest/download/sendem-darwin-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o - https://github.com/alasdairmorris/sendem/releases/latest/download/sendem-darwin-arm64.tar.gz | tar -zxf - -C $HOME/bin
```

### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/sendem@latest
```

## Usage

```
A command-line tool for sending emails via SMTP.

Usage:
  sendem [-f FROM] [-c CC ...] [-b BCC ...] [-s SUBJECT] [-m FILE] [-a FILE ...]
            [-u USERNAME] [-p PASSWORD] [-x HOST:PORT] [-H] RECIPIENT...
  sendem -h | --help
  sendem --version

Options:
  -h, --help              Show this screen.
  --version               Show version.
  -f, --from ADDR         From this address.
  -c, --cc ADDR           CC this address (for multiple, use "-c a1 -c a2 ...")
  -b, --bcc ADDR          BCC this address (for multiple, use "-b a1 -b a2 ...")
  -s, --subject SUBJECT   Subject
  -m, --message FILE      File to use for message body [default is stdin]
  -a, --attach FILE       File to attach (for multiple, use "-a f1 -a f2 ...")
  -H, --html              Treat message body as HTML.
  -u, --user USERNAME     Username for SMTP connection.
  -p, --pass PASSWORD     Password for SMTP connection.
  -x, --server HOST:PORT  Server to use for SMTP [default: localhost:25]

Homepage: https://github.com/alasdairmorris/sendem
```

## Examples

```
$ date | sendem -a /tmp/file.pdf -f sender@whatever.com -s Subject recipient@example.com
```
