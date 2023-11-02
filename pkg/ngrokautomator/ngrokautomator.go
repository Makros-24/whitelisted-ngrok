package ngrokautomator

import (
	"context"
	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	"strings"
)

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

func Run(ctx context.Context, dest string, whitelist []string, tunnelType config.Tunnel) error {
	tun, ngrokListenerError := ngrok.Listen(ctx,
		tunnelType,
		ngrok.WithAuthtoken(viper.GetString("ngrok.token")),
	)

	if ngrokListenerError != nil {
		return ngrokListenerError
	}

	log.Println("tunnel created:", tun.URL())

	if viper.GetBool("notification.active") {
		mailError := sendNewUrl(
			Email{
				sender:     viper.GetString("smtp.username"),
				recipients: viper.GetString("notification.url.recipients"),
				variables:  []string{tun.URL()},
			},
			SmtpConfig{
				server:   viper.GetString("smtp.server.host"),
				port:     viper.GetInt("smtp.server.port"),
				username: viper.GetString("smtp.username"),
				password: viper.GetString("smtp.password"),
			},
		)

		if mailError != nil {
			return mailError
		}
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
