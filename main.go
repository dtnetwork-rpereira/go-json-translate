package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gt "github.com/bas24/googletranslatefree"
)

func main() {
	languageFile, err := os.Open("pt-BR.json")

	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo")
		return
	}

	defer languageFile.Close()

	languageObject, err := io.ReadAll(languageFile)

	if err != nil {
		fmt.Println("Não foi possível carregar os dados do arquivo")
		return
	}

	var result map[string]string

	json.Unmarshal([]byte(languageObject), &result)

	exportedLanguageCode := "az"

	for _, v := range result {
		newTranslation, _ := gt.Translate(v, "pt-BR", exportedLanguageCode)
		fmt.Println(newTranslation)
	}
}
