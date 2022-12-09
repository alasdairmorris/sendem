package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"gopkg.in/gomail.v2"
)

const version = "v1.0.0"

const usage = `A command-line tool for sending emails via SMTP.

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
`

type Config struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Message     []byte
	Attachments []string
	IsHtml      bool
	Username    string
	Password    string
	Host        string
	Port        int
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		config Config
		opts   docopt.Opts
		infile string
		file   *os.File
		server string
		err    error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	if err != nil {
		log.Fatal(err)
	}

	// From address
	config.From, err = opts.String("--from")
	if err != nil || config.From == "" {
		var currentuser *user.User
		currentuser, err = user.Current()
		if err != nil {
			log.Fatal(err)
		}

		var hostname string
		hostname, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		config.From = fmt.Sprintf("%s@%s", currentuser.Username, hostname)
	}

	// To address(es)
	if recipients, ok := opts["RECIPIENT"].([]string); ok {
		for _, r := range recipients {
			_, err = mail.ParseAddress(r)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error parsing email address %s - %s", r, err))
			}
			config.To = append(config.To, r)
		}
	}

	// CC address(es)
	if recipients, ok := opts["--cc"].([]string); ok {
		for _, r := range recipients {
			_, err = mail.ParseAddress(r)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error parsing email address %s - %s", r, err))
			}
			config.Cc = append(config.Cc, r)
		}
	}

	// BCC address(es)
	if recipients, ok := opts["--bcc"].([]string); ok {
		for _, r := range recipients {
			_, err = mail.ParseAddress(r)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error parsing email address %s - %s", r, err))
			}
			config.Bcc = append(config.Bcc, r)
		}
	}

	// Subject
	config.Subject, err = opts.String("--subject")

	// Message body
	infile, err = opts.String("--message")
	if infile == "" || infile == "-" {
		config.Message, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config.Message, err = os.ReadFile(infile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Attachments
	if attachments, ok := opts["--attach"].([]string); ok {
		for _, a := range attachments {
			if file, err = os.Open(a); err != nil {
				log.Fatal(err)
			} else {
				config.Attachments = append(config.Attachments, a)
				file.Close()
			}
		}
	}

	// Is HTML?
	config.IsHtml, err = opts.Bool("--html")

	// Username
	config.Username, err = opts.String("--user")

	// Password
	config.Password, err = opts.String("--pass")

	// Host:Port
	server, err = opts.String("--server")
	if err != nil {
		log.Fatal(err)
	}
	bits := strings.Split(server, ":")
	if len(bits) != 2 {
		log.Fatal("Invalid format of server parameter: " + server)
	}
	config.Host = bits[0]
	config.Port, err = strconv.Atoi(bits[1])
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {

	log.SetFlags(0)

	config := getConfig()

	m := gomail.NewMessage()

	m.SetHeader("From", config.From)

	m.SetHeader("To", config.To...)

	if len(config.Cc) > 0 {
		m.SetHeader("Cc", config.Cc...)
	}

	if len(config.Bcc) > 0 {
		m.SetHeader("Bcc", config.Bcc...)
	}

	m.SetHeader("Subject", config.Subject)

	if config.IsHtml {
		m.SetBody("text/plain", "Please use an HTML-capable email client to view this message.")
		m.AddAlternative("text/html", string(config.Message))
	} else {
		m.SetBody("text/plain", string(config.Message))
	}

	for _, a := range config.Attachments {
		m.Attach(a)
	}

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

}
