// +build !windows

package examples

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func ServeReuse() {
	lc := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			if err := c.Control(func(fd uintptr) {
				err = syscall.SetsockoptInt(
					int(fd),
					unix.SOL_SOCKET,
					unix.SO_REUSEPORT,
					1,
				)
			}); err != nil {
				return err
			}
			return err
		},
	}

	ln, err := lc.Listen(
		context.Background(),
		"tcp",
		os.Getenv("HOST")+":"+os.Getenv("PORT"),
	)
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from %s\n", os.Getenv("INSTANCE"))))
	})

	if err := http.Serve(ln, nil); err != nil {
		panic(err.Error())
	}
}
