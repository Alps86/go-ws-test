package main

import (
	"net/http"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"fmt"
)

func main() {
	http.ListenAndServe(":1718", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w, nil)
		if err != nil {
			fmt.Println(err)
		}

		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println(err)
				}

				s := string(msg[:])
				fmt.Println(s)

				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}))
}