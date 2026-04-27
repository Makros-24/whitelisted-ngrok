# ngka — Whitelisted ngrok Tunnels

A CLI tool that creates [ngrok](https://ngrok.com/) tunnels with IP whitelisting. Only connections from whitelisted IP addresses are forwarded to the destination; all others are rejected. Supports both TCP and HTTP tunnel types, with optional email notifications when a tunnel URL is created.

## Prerequisites

- Go 1.21+
- A [ngrok account](https://dashboard.ngrok.com/signup) and authentication token

## Installation

```bash
git clone https://github.com/Makros-24/whitelisted-ngrok.git
cd whitelisted-ngrok
go mod tidy
go build -o ngka .
```

## Configuration

Edit `config.yml` with your values before running:

```yaml
ngrok:
  token: your_ngrok_auth_token

smtp:
  username: you@gmail.com
  password: your_app_password
  server:
    host: smtp.gmail.com
    port: 587

notification:
  active: false
  url:
    recipients: notify@example.com
```

| Key | Required | Description |
|-----|----------|-------------|
| `ngrok.token` | Yes | Your ngrok authentication token |
| `notification.active` | No (default: `false`) | Send an email when a tunnel URL is created |
| `smtp.*` | Only when `notification.active: true` | SMTP credentials for email delivery |
| `notification.url.recipients` | Only when `notification.active: true` | Recipient email address |

> **Note:** `config.yml` is listed in `.gitignore` — do not rename it or you risk committing credentials.

## Usage

### TCP tunnel

```bash
./ngka tcp -w <ip1> [-w <ip2> ...] <address:port>
```

### HTTP tunnel

```bash
./ngka http -w <ip1> [-w <ip2> ...] <address:port>
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--whitelist` | `-w` | IP address allowed to connect (repeatable, required) |

### Examples

Expose a local SSH server, only reachable from one IP:

```bash
./ngka tcp -w 203.0.113.42 localhost:22
```

Expose a local web server to two IPs:

```bash
./ngka http -w 203.0.113.42 -w 198.51.100.7 localhost:8080
```

## How it works

1. An ngrok tunnel is created using your auth token.
2. The tunnel URL is logged (and emailed if notifications are enabled).
3. For each incoming connection, the remote IP is checked against the whitelist.
   - **Allowed:** the connection is proxied transparently to `<address:port>`.
   - **Rejected:** the connection is closed immediately.

## License

MIT — see [LICENSE](LICENSE).
