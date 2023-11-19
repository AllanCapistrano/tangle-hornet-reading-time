package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/allancapistrano/tangle-client-go/messages"
	"github.com/kyokomi/emoji/v2"
)

const DIRECTORY_NAME = "files"

func main() {
	var amountMessagesInString string
	var amountMessages int
	var index string

	nodeURL := "http://127.0.0.1:14265"

	amountMessagesParameter := flag.Int("qtm", -1, "Quantidade de mensagens")
	indexParameter := flag.String("idx", "", "Índice das mensagens")

	flag.Parse()
	
	if (*amountMessagesParameter == -1) {
		var err error
		fmt.Print("Digite quantas mensagens você quer gerar: ")
		fmt.Scanln(&amountMessagesInString)

		amountMessages, err = strconv.Atoi(amountMessagesInString)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		amountMessages = *amountMessagesParameter
	}

	if amountMessages < 0 {
		log.Fatal("invalid amount of messages")
	}

	if (*indexParameter == "") {
		fmt.Print("Digite o índice para as mensagens: ")
		fmt.Scanln(&index)
	} else {
		index = *indexParameter
	}

	message := "{\"available\":true,\"avgLoad\":3,\"createdAt\":1695652263921,\"group\":\"group3\",\"lastLoad\":4,\"publishedAt\":1695652267529,\"source\":\"source4\",\"type\":\"LB_STATUS\"}"

	fmt.Printf("Mensagem que será publicada: %s\n", message)

	files, err := os.ReadDir(DIRECTORY_NAME)
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf("tangle-hornet-reading-time_%d.csv", len(files))
	filePath := fmt.Sprintf("%s/%s", DIRECTORY_NAME, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing header to csv file.
	header := []string{"Índice", "Tempo de consulta (s)"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	fmt.Println(emoji.Sprint("\n:hourglass:Inciando a publicação e leitura das mensagens."))

	for i := 0; i < amountMessages; i++ {
		// Submitting a message
		messages.SubmitMessage(nodeURL, index, message, 15)

		start := time.Now()

		// Getting all messages from an index.
		_, err := messages.GetAllMessagesByIndex(nodeURL, index)
		if err != nil {
			log.Fatal(err)
		}

		elapsed := time.Since(start)
		elapsedInString := strconv.FormatFloat(elapsed.Seconds(), 'f', -1, 64)

		// Writing data to csv file
		row := []string{strconv.Itoa(i + 1), elapsedInString}
		if err := writer.Write(row); err != nil {
			log.Fatal(err)
		}

		if i == amountMessages/4 {
			fmt.Println(emoji.Sprint(":heavy_check_mark: 25% das mensagens já foram publicadas e consultadas."))
		} else if i == amountMessages/2 {
			fmt.Println(emoji.Sprint(":heavy_check_mark: 50% das mensagens já foram publicadas e consultadas."))
		} else if i == amountMessages/4+amountMessages/2 {
			fmt.Println(emoji.Sprint(":heavy_check_mark: 75% das mensagens já foram publicadas e consultadas."))
		}
	}

	fmt.Println(emoji.Sprintf("\n:white_check_mark:Experimento concluído, o arquivo '%s' encontra-se do diretório '%s'.", fileName, DIRECTORY_NAME))
}
