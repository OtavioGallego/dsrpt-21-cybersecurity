package cli

import (
	"github.com/urfave/cli"
)

// Gerar retorna a aplicação de linha de comando pronta para ser utilizada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de CVEs na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "vendor",
			Usage: "nome do vendor",
			Value: "",
		},
	}

	flagsCve := []cli.Flag{
		cli.StringFlag{
			Name:  "product",
			Usage: "nome do produto",
			Value: "",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "id no formato CVE-9999-9999",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "vendors",
			Usage:  "Busca de vendors registrados",
			Flags:  flags,
			Action: buscarVendors,
		},
		{
			Name:   "products",
			Usage:  "Busca de produtos de um determinado vendor",
			Flags:  flags,
			Action: buscarProdutos,
		},
		{
			Name:   "cve",
			Usage:  "Busca de cve por vendor e produto, ou por CVE-ID",
			Flags:  append(flags, flagsCve...),
			Action: buscarCve,
		},
	}

	return app
}
