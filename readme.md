# Busca de CVEs
Esta é uma aplicação de linha de comando escrita em Go, cujo objetivo é buscar informações sobre CVEs disponíveis na API pública [CVE Search](https://www.cve-search.org/api/#public-web-api-of-cve-search)

## Utilização
Para utilizar a API é necessário ter o [Go](https://golang.org/dl/) instalado no sistema operacional de sua preferência

Após a instalação, basta utilizar o seguinte comando na raiz do projeto para instalar todas as dependências

```bash
go mod download
```
Após isso, será possível utilizar a aplicação utilizando os comandos abaixos:

### Busca de vendors
Para buscar a lista completa de vendors registrados na API, basta utilizar o comando *vendors*, conforme especificado abaixo
```bash
go run main.go vendors
```

Caso queira filtrar os resultados, é possível adicionar a flag *vendors* para correspondência completa ou parcial
```bash
go run main.go vendors --vendor micr
```

### Busca de produtos
Para buscar a lista de produtos associados a um vendor, é possível utilizar o comando *products*. Note que neste caso a flag *vendor* é obrigatória
```bash
go run main.go products --vendor microsoft
```

### Busca de CVE
Existem duas formas de buscar uma CVE, através do comando *cve*. 

A primeira delas é utilizando as flags *vendors* e *product*, conforme descrito abaixo
```bash
go run main.go cve --vendor microsoft --product outlook
```

A segunda forma é utilizando o ID da CVE, no formato **CVE-9999-9999**, através do parâmetro *id* conforme descrito abaixo:
```bash
go run main.go cve --id CVE-2010-3333
```

Por conter muitas informações, a CVE obtida é sempre salva em um arquivo chamado **cve.json**, na raiz do projeto, independente de qual das duas formas tenha sido utilizada para buscá-la.

## Componentes do Grupo

Gabriel Diamantino Batista - RM 79262\
Leonardo Ulisses Amorim Castanho - RM 79491\
Otávio Augusto Gallego - RM 77955\
Victor Wilson Costa Lamana - RM 77835\
Vinicius Rodrigues de Barros - RM 77936\


## Licença
[MIT](https://choosealicense.com/licenses/mit/)
