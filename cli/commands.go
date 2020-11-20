package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/urfave/cli"
)

var (
	url string = "https://cve.circl.lu/api/browse"
)

type responseVendor struct {
	Product string   `json:"product"`
	Vendor  []string `json:"vendor"`
}

type responseProduct struct {
	Product []string `json:"product"`
}

func buscarVendors(c *cli.Context) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	var responseVendor responseVendor
	err = json.Unmarshal(respBody, &responseVendor)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	vendor := c.String("vendor")

	if vendor == "" {
		for _, v := range responseVendor.Vendor {
			fmt.Println(v)
		}

	} else {
		for _, v := range responseVendor.Vendor {
			if strings.Contains(v, vendor) {
				fmt.Println(v)
			}
		}
	}
}

func buscarProdutos(c *cli.Context) {
	vendor := c.String("vendor")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url, vendor), nil)
	if err != nil {
		log.Fatal("Erro ao buscar produtos => ", err)
	}

	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	var response responseProduct
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		log.Fatal("Erro ao buscar vendors => ", err)
	}

	for _, p := range response.Product {
		fmt.Println(p)
	}
}

func buscarCve(c *cli.Context) {
	var (
		vendor  = c.String("vendor")
		product = c.String("product")
		id      = c.String("id")
		url     string
	)

	if id != "" {
		url = fmt.Sprintf("https://cve.circl.lu/api/cve/%s", id)
	} else {
		url = fmt.Sprintf("https://cve.circl.lu/api/search/%s/%s", vendor, product)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Erro ao buscar cve => ", err)
	}

	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao buscar cve => ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao buscar cve => ", err)
	}

	err = ioutil.WriteFile("cve.json", respBody, 0644)
	if err != nil {
		log.Fatal("Erro ao salvar cve em arquivo JSON => ", err)
	}

	log.Println("CVE salva no arquivo cve.json")
}
