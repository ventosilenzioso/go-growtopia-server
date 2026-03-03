package httpsserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

var (
	cyan    = color.New(color.FgCyan).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
)

func logRequest(r *http.Request, statusCode int) {
	timestamp := time.Now().Format("15:04:05")
	method := r.Method
	path := r.URL.Path
	ip := r.RemoteAddr
	userAgent := r.UserAgent()

	var statusColor func(...interface{}) string
	if statusCode >= 200 && statusCode < 300 {
		statusColor = green
	} else if statusCode >= 300 && statusCode < 400 {
		statusColor = yellow
	} else {
		statusColor = red
	}

	fmt.Printf("%s %s %s %s from %s | %s\n",
		cyan(timestamp),
		magenta(method),
		path,
		statusColor(statusCode),
		yellow(ip),
		userAgent,
	)
}
