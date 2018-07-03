package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/johnmccabe/go-vmpooler/token"
	"golang.org/x/crypto/ssh/terminal"
)

const vmpoolerEndpoint = "https://vmpooler.mydomain.com/api/v1"

func main() {
	log.Print("New Token Client ===")
	username, password, err := getCredentials()
	if err != nil {
		log.Fatal(err.Error())
	}

	t := token.NewClient(vmpoolerEndpoint, username, password)

	log.Print("Create Token ===")
	token, err := t.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Token generated: %s", token)

	log.Print("Get Token ===")
	tok, err := t.Get(token.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Token retrieved: %s", tok)

	log.Print("GetAll Tokens ===")
	toks, err := t.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Tokens retrieved: %s", toks)

	log.Print("Delete Token ===")
	if err := t.Delete(token.Token); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Token deleted: %s", token)

}

func getCredentials() (string, string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println()

	return username, password, err
}
