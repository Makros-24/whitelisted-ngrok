#### [WORK IN PROGRESS] 
# NGROK Whitelist Server

This is a Go application for creating a ngrok tunnel with the option of specifying a whitelist. The server only accepts connections from whitelisted IP addresses and forwards them to a specified destination. Additionally, it could send an email with the ngrok URL when a tunnel is established.

## Prerequisites

Before running this application, you need to have the following prerequisites installed:

- Go (Golang) version
- The `go-mail` library for sending notifications
- The `ngrok` library creating tunnels
- The `cobra` library for cli
- The `viper` library for configuration
- A ngrok account and authentication token

## Installation

To install all the required dependencies, you can use:
```bash
go get [packages]
```

alternatively, you install required dependencies manually using the following commands:
```bash
go mod tidy
```

## Building

```bash
go build main.go
```

## Configuration
Make sure to set the following environment variables for ngrok:

ngrok.token: Your ngrok authentication token
notification.active: Whether to send an email notification when a tunnel is established if so the next configurations must be specified
    smtp.username: Your email username
    smtp.password: Your email password
    smtp.server.host: smtp.server.com
    smtp.server.port: ex:587


## Usage
```bash
go run main.go tcp <address:port> 
```

<address:port> The destination address and port where the connections will be forwarded.

## Whitelist
You can specify the IP addresses that are allowed to connect to the server by adding them to the whitelist slice in the main function:

```bash
-w <ip-address-1> -w <ip-address-n> 
```

## Running the Server
Run the server with the command mentioned in the "Usage" section. It will create an ngrok tunnel and start listening for incoming connections from whitelisted IP addresses.

## Handling Connections
The server will accept incoming connections and forward them to the specified destination. If an IP address is not in the whitelist, the connection will be refused.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
