# log-ag

A lightweight CLI log aggregator built in Go. Collect, search, and monitor logs from any app or service — all stored locally in SQLite.

## Install

### From source

```bash
git clone https://github.com/mannansainicyber/log-aggregator
cd log-aggregator
go build -o log-ag
```

Move the binary to your PATH so you can run it from anywhere:

```bash
sudo mv log-ag /usr/local/bin/
```

### Requirements

- Go 1.21+
- GCC (required for SQLite — `sudo apt install gcc` on Linux)

---

## Usage

### Send a log

```bash
log-ag send "server started"
log-ag send --level error --service api "user login failed"
log-ag send -l warn -s db "slow query detected"
```

Flags:

| Flag | Short | Default | Description |
|---|---|---|---|
| `--level` | `-l` | `info` | Log level (info, warn, error) |
| `--service` | `-s` | `default` | Service name |

---

### Search logs

```bash
log-ag search
log-ag search --level error
log-ag search --service api
log-ag search --since "1 hour"
log-ag search --limit 10
```

Flags:

| Flag | Short | Default | Description |
|---|---|---|---|
| `--level` | `-l` | | Filter by level |
| `--service` | `-s` | | Filter by service |
| `--since` | `-t` | | Time filter e.g. `1 hour`, `30 minutes` |
| `--limit` | `-n` | `50` | Max results to return |

---

### Watch logs live

```bash
log-ag watch
```

Tails incoming logs in real time. Press `Ctrl+C` to stop.

---

### Stats

```bash
log-ag stats
```

Shows log counts grouped by level and service.

```
─────────────────────────────────────────
  BY LEVEL
─────────────────────────────────────────
  error      5
  warn       3
  info       12
─────────────────────────────────────────
  BY SERVICE
─────────────────────────────────────────
  api        10
  db         6
  default    4
─────────────────────────────────────────
```

---

### Clear all logs

```bash
log-ag clear
```

Wipes the entire log database.

---

## Storage

Logs are stored locally at:

```
~/.log-ag/logs.db
```

No server, no cloud, no config needed.

---

## Project structure

```
log-ag/
├── main.go
├── cmd/
│   ├── root.go
│   ├── send.go
│   ├── search.go
│   ├── watch.go
│   ├── stats.go
│   └── clear.go
├── db/
│   └── db.go
├── go.mod
└── README.md
```

---

## Built with

- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — SQLite driver

---

## License

MIT