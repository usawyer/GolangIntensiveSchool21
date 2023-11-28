package main

import (
	"bytes"
	"candies/restapi/operations"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
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

func validateInput() (Request, error) {
	var order Request
	flag.Int64Var(&order.CandyCount, "c", 0, "count of candy to buy")
	flag.StringVar(&order.CandyType, "k", "", "two-letter abbreviation for the candy type CE/AA/NT/DE/YR")
	flag.Int64Var(&order.Money, "m", 0, "amount of money you \"gave to machine\"")
	flag.Parse()

	var candies = map[string]int64{"CE": 10, "AA": 15, "NT": 17, "DE": 21, "YR": 23}
	if _, ok := candies[order.CandyType]; !ok {
		return Request{}, errors.New("invalid candy type, should be CE/AA/NT/DE/YR")
	}
	if order.CandyCount < 0 || order.Money < 0 {
		return Request{}, errors.New("invalid value")
	}
	return order, nil
}

func makeRequest(order Request) ([]byte, *http.Response, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, r, nil
}

func printAnswer(body []byte, r *http.Response) {
	if r.StatusCode == 201 {
		var success operations.BuyCandyCreatedBody
		err := json.Unmarshal(body, &success)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s You change is %d!\n", success.Thanks, success.Change)
	} else if r.StatusCode == 402 {
		var fail operations.BuyCandyPaymentRequiredBody
		err := json.Unmarshal(body, &fail)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(fail.Error)
	} else if r.StatusCode == 400 {
		var fail operations.BuyCandyBadRequestBody
		err := json.Unmarshal(body, &fail)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(fail.Error)
	} else {
		fmt.Print(string(body))
	}
}

func main() {
	order, err := validateInput()
	if err != nil {
		log.Fatal(err)
	}

	body, r, err := makeRequest(order)
	if err != nil {
		log.Fatal(err)
	}

	printAnswer(body, r)
}
