package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	nome string `json:"-"`
	raça string `json:"-"`
}

func main() {
	c := cachorro{"Bili", "putbul"}

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

}
