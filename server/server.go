package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var config []byte
	var conn *sql.DB
	config, err = initialize()
	connMysql(config)
	defer conn.Close()
	fmt.Println("Welcome to MobChecker")
	mux := http.NewServeMux()
	mux.HandleFunc("/", copyright)
	mux.HandleFunc("/click", click)
	http.ListenAndServe(":8080", mux)

}

func copyright(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Powered by MobChecker!"))
}

func click(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/test", http.StatusFound)

}

func initialize() ([]byte, err) {
	configFile := "config/config.json"
	return ioutil.ReadFile(configFile)
}

func connMysql(config []byte) (*sql.DB, err) {
	var info struct {
		Mysql string
	}
	return sql.Open("mysql", info.Mysql)
}
