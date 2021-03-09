package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
        config, err := pgx.ParseConfig("postgresql://abc@localhost:26257/bank?sslcert=client.abc.crt&sslkey=client.abc.key&sslmode=verify-full&sslrootcert=ca.crt")
		if err != nil {
			log.Fatal("error configuring the database: ", err)
		}

		conn, err := pgx.ConnectConfig(context.Background(), config)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
		
        defer conn.Close(context.Background())

		rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
		if err != nil {
			log.Fatal(err)
		}
		
        defer rows.Close()
        
        fmt.Fprintf(w, "Initial balances: \n")
        
        for rows.Next() {
			var id, balance int
			if err := rows.Scan(&id, &balance); err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "%d - ", id)
			fmt.Fprintf(w, "%d \n", balance)
		} 
	})
	http.ListenAndServe(":8088", nil)
}
