package main

import (
	"fmt"
	"log"

	"github.com/Dedo-Finger2/notion-file-helper/utils"
)

/*
	Este projeto consiste em um
	helper para o Notion
	onde ele:

	Com áudios, ele divide em vários outros
	arquivos de áudio com um tamanho
	máximo de 4,9MB por arquivo.
	Depois, coloca tudo dentro de um zip
	e joga isso numa pasta do usuário.

	Com imagens, ele reduzirá a qualidade
	da imagem até onde der, tentando deixar
	ela com no máximo 4,9MB de tamanho.

	Com vídeos, ele fará o mesmo processo
	feito com os áudios.
*/

func main() {
	fmt.Println("Hello, World!")

	var p string

	fmt.Scan(&p)

	sb, smb, err := utils.GetFileSizeFromPath(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sb)
	fmt.Printf("%.2f", smb)
}
