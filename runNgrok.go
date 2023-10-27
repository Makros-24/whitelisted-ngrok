package main

import (
	"context"
	"github.com/go-mail/mail"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func usage(bin string) {
	log.Fatalf("Usage: %s <address:port>", bin)
}

func main() {
	if len(os.Args) != 2 {
		usage(os.Args[0])
	}
	if err := run(context.Background(), os.Args[1], []string{"ip1.*.*.*", "ip2.*.*.*"}); err != nil {
		log.Fatal(err)
	}
}

func sendMail(url string) error {
	m := mail.NewMessage()

	m.SetHeader("From", "makros.server@example.com")
	m.SetHeader("To", "makros@example.com")
	m.SetHeader("Subject", "Home Server")
	m.SetBody("text/html", "    <p>Hello,</p>\n    <p>Here is the link you requested:</p>\n    <a href=\""+url+"\">Click here to access the URL</a>\n    <p>If you have any questions or need further assistance, please hesitate to contact me, here is my fake number 25311184.</p>\n    <p>Best regards,</p>\n    <p>Your Name</p>")

	d := mail.NewDialer("smtp.gmail.com", 587, "makros.server@example.com", "passwordexample")

	if err := d.DialAndSend(m); err != nil {

		panic(err)
		return err
	}
	return nil
}

func checkWhitelist(whitelist []string, remoteIP string) bool {
	var isIPInWhitelist bool
	for _, ip := range whitelist {
		if ip == remoteIP {
			isIPInWhitelist = true
			break
		}
	}
	return isIPInWhitelist
}

func run(ctx context.Context, dest string, whitelist []string) error {
	tun, ngrokListenerError := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)

	if ngrokListenerError != nil {
		return ngrokListenerError
	}

	log.Println("tunnel created:", tun.URL())

	mailError := sendMail(tun.URL())
	if mailError != nil {
		return mailError
	}

	for {
		//Accept Connection
		conn, connectionError := tun.Accept()
		if connectionError != nil {
			return connectionError
		}

		//Whitelist Check
		remoteIP := strings.Split(conn.RemoteAddr().String(), ":")[0]
		log.Println(remoteIP + " requesting connection from server")

		if checkWhitelist(whitelist, remoteIP) {
			log.Println("whitelisted")
			log.Println("accepted connection from", conn.RemoteAddr())

			//handle connection
			go func() {
				err := handleConn(ctx, dest, conn)
				log.Println("connection closed:", err)
			}()
		} else {
			log.Println(remoteIP + " connection refused by server")
		}
	}
}

func handleConn(ctx context.Context, dest string, conn net.Conn) error {
	next, err := net.Dial("tcp", dest)
	if err != nil {
		return err
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := io.Copy(next, conn)
		return err
	})
	g.Go(func() error {
		_, err := io.Copy(conn, next)
		return err
	})

	return g.Wait()
}
