# go-growtopia-server 🚀

Simple Growtopia ENet server with HTTPS support built in Golang.

## 📖 Overview

This project is a simple Growtopia private server implementation using:

* 🌐 ENet for game networking (using cgo bindings)
* 🔐 HTTPS server (port 443)
* 🌍 HTTP server (port 80)
* ⚙️ Built entirely in Go (Golang)

Designed for learning, experimentation, and protocol research.

## ✨ Features

* 🎮 ENet-based game server (ENet via cgo)
* 🌍 Built-in HTTP (80)
* 🔐 Built-in HTTPS (443)
* ⚡ Lightweight and minimal setup
* 🧩 Written in pure Go

## 📦 Requirements

* Go 1.20+
* Git

## 🚀 Installation

```bash
git clone https://github.com/yourusername/go-growtopia-server.git
cd go-growtopia-server
go mod tidy
go run main.go
```

## 🔌 Ports

* 🌍 80   → HTTP
* 🔐 443  → HTTPS
* 🎮 ENet port → Configurable inside `ServerConfiguration.toml`

## ⚠️ Notes

This project is a simple implementation and is not affiliated with the official Growtopia game.

Use at your own risk.

---

Made with ❤️ using Go.

---

## 📜 License

This project is licensed under the MIT License.

You are free to use, modify, and distribute this software in accordance with the license terms.

---

## 🙏 Credits

* ENet Networking Library
  [https://github.com/lsalzman/enet](https://github.com/lsalzman/enet)

* [https://github.com/gtpshax](https://github.com/gtpshax)

* [https://github.com/yoruakio](https://github.com/yoruakio)

Special thanks to the open-source community for networking and protocol research contributions.
