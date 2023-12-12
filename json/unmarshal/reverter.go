package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	nome string `json:"-"`
	raça string `json:"-"`
}

func main() {

	cachorroJSON := `{"nome":"bili","raça":"putbul"}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}
