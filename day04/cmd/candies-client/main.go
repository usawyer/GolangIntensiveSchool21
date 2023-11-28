package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Request struct {
	CandyCount int64  `json:"candyCount"`
	CandyType  string `json:"candyType"`
	Money      int64  `json:"money"`
}

const url = "https://127.0.0.1:3333/buy_candy"

func main() {
	var order Request
	flag.Int64Var(&order.CandyCount, "c", 0, "count of candy to buy")
	flag.StringVar(&order.CandyType, "k", "", "two-letter abbreviation for the candy type CE/AA/NT/DE/YR")
	flag.Int64Var(&order.Money, "m", 0, "amount of money you \"gave to machine\"")
	flag.Parse()

	//validation
	//var candies = map[string]int64{"CE": 10, "AA": 15, "NT": 17, "DE": 21, "YR": 23}
	//fmt.Println(order)

	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	caCert, _ := os.ReadFile("cert/minica.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, _ := tls.LoadX509KeyPair("cert/client/cert.pem", "cert/client/key.pem")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	r, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
