package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"go_team01/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
)

type clientData struct {
	warehouseClient pb.WarehouseClient
}

func (c *clientData) handler() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		txt := scanner.Text()
		str := strings.Split(txt, " ")

		switch str[0] {
		case "GET":
			if len(str) == 2 {
				res, err := c.warehouseClient.GetDocument(context.Background(), &pb.GetDocumentRequest{Key: str[1]})
				if err != nil {
					log.Fatal("Getting document failed ", err)
				}
				fmt.Println("GET operation")
				_ = res
			}
		case "SET":
			fmt.Println("SET operation")
		case "DELETE":
			fmt.Println("DELETE operation")
		default:
			fmt.Println("Unknown operation")
		}
	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	defer conn.Close()

	warehouseClient := pb.NewWarehouseClient(conn)

	client := clientData{warehouseClient: warehouseClient}
	client.handler()
}
