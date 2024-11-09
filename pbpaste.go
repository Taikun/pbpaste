package main

import (
    "fmt"
    "log"

    "github.com/atotto/clipboard"
)

func main() {
    // Leer el contenido del portapapeles
    text, err := clipboard.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Imprimir el contenido
    fmt.Println(text)
}