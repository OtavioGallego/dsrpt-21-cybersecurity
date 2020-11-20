package main

import (
	"log"
	"os"

	"github.com/OtavioGallego/dsrpt-21-cybersecurity/cli"
)

func main() {
	aplicacao := cli.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
