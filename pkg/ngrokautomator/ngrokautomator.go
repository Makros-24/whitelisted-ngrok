package ngrokautomator

import (
	"context"
	"io"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, dest string, whitelist []string, tunnelType config.Tunnel) error {
	tun, err := ngrok.Listen(ctx, tunnelType, ngrok.WithAuthtoken(viper.GetString("ngrok.token")))
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	if viper.GetBool("notification.active") {
		if err := sendNewUrl(
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
		); err != nil {
			return err
		}
	}

	allowed := make(map[string]struct{}, len(whitelist))
	for _, ip := range whitelist {
		allowed[ip] = struct{}{}
	}

	for {
		conn, err := tun.Accept()
		if err != nil {
			return err
		}

		remoteIP := strings.Split(conn.RemoteAddr().String(), ":")[0]
		log.Println(remoteIP, "requesting connection")

		if _, ok := allowed[remoteIP]; ok {
			log.Println("accepted connection from", conn.RemoteAddr())
			go func() {
				if err := handleConn(ctx, dest, conn); err != nil {
					log.Println("connection closed:", err)
				}
			}()
		} else {
			log.Println(remoteIP, "connection refused")
			conn.Close()
		}
	}
}

func handleConn(ctx context.Context, dest string, src net.Conn) error {
	defer src.Close()

	dst, err := net.DialContext(ctx, "tcp", dest)
	if err != nil {
		return err
	}
	defer dst.Close()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer dst.Close()
		_, err := io.Copy(dst, src)
		return err
	})
	g.Go(func() error {
		defer src.Close()
		_, err := io.Copy(src, dst)
		return err
	})

	return g.Wait()
}
