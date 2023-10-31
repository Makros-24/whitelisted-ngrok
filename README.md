#### [WORK IN PROGRESS] 
# NGROK Whitelist Server

This is a Go application for creating a whitelist server using ngrok. The server only accepts connections from whitelisted IP addresses and forwards them to a specified destination. Additionally, it sends an email with the ngrok URL when a connection is established.

## Prerequisites

Before running this application, you need to have the following prerequisites installed:

- Go (Golang)
- The `go-mail` library
- An ngrok account and authentication token

## Installation

To install the required dependencies, you can use `go get`:

```bash
go get gopkg.in/gomail.v2
go get golang.ngrok.com/ngrok
```

## Configuration
Make sure to set the following environment variables for ngrok:

NGROK_AUTHTOKEN: Your ngrok authentication token

## Usage
```bash
go run main.go <address:port>
```

<address:port>: The destination address and port where the connections will be forwarded.

## Whitelist
You can specify the IP addresses that are allowed to connect to the server by adding them to the whitelist slice in the main function:

```bash
whitelist := []string{"193.*.*.*", "196.*.*.*"}
```

## Running the Server
Run the server with the command mentioned in the "Usage" section. It will create an ngrok tunnel and start listening for incoming connections from whitelisted IP addresses.

## Handling Connections
The server will accept incoming connections and forward them to the specified destination. If an IP address is not in the whitelist, the connection will be refused.

## Email Notification
The server will send an email with the ngrok URL to the specified recipient when a connection is established.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

