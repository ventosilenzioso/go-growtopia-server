<h1 align="center">go-growtopia-server 🚀</h1>

<p align="center">
  Simple Growtopia ENet (cgo) server with built-in HTTP (80) and HTTPS (443) support written in Go.
</p>

<hr/>

<h2>📖 Overview</h2>

<p>
This project is a lightweight Growtopia private server implementation using:
</p>

<ul>
  <li>🌐 ENet for game networking (via cgo bindings)</li>
  <li>🔐 HTTPS server (Port 443)</li>
  <li>🌍 HTTP server (Port 80)</li>
  <li>⚙️ Fully written in Golang</li>
</ul>

<p>
Designed for learning purposes, experimentation, and protocol research.
</p>

<hr/>

<h2>✨ Features</h2>

<ul>
  <li>🎮 ENet-based multiplayer server (cgo)</li>
  <li>🌍 Integrated HTTP service</li>
  <li>🔐 Integrated HTTPS service</li>
  <li>⚡ Minimal and lightweight architecture</li>
  <li>🧩 Clean Go project structure</li>
</ul>

<hr/>

<h2>📦 Requirements</h2>

<ul>
  <li>Go 1.20+</li>
  <li>Git</li>
</ul>

<hr/>

<h2>🚀 Installation</h2>

<pre><code>git clone https://github.com/yourusername/go-growtopia-server.git
cd go-growtopia-server
go mod tidy
go run main.go
</code></pre>

<hr/>

<h2>🔌 Ports</h2>

<ul>
  <li>🌍 <strong>80</strong> → HTTP</li>
  <li>🔐 <strong>443</strong> → HTTPS</li>
  <li>🎮 ENet Port → Configurable inside <code>ServerConfiguration.toml</code></li>
</ul>

<hr/>

<h2>📜 License</h2>

<p>
This project is licensed under the MIT License.
</p>

<hr/>

<h2>🙏 Credits</h2>

<ul>
  <li>ENet Networking Library</li>
  <li>https://github.com/gtpshax</li>
  <li>https://github.com/yoruakio</li>
</ul>

<p>
This project is not affiliated with the official Growtopia game.
</p>

<hr/>

<p align="center">
Made with ❤️ using Go
</p>
