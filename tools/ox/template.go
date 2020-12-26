package ox

var mainTemplate = `
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/wawandco/oxpecker/cli"
	wawandco "github.com/wawandco/oxpecker-plugins"
)

func main() {
  	fmt.Print("~~~~ Using {{.Name}}/cmd/ox ~~~\n\n")
	ctx := context.Background()
    
  	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
  	}
    
	ox := cli.New()
	// append your plugins here
	ox.Plugins = append(wawandco.All, ...)
    
    err = ox.Run(ctx, pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
`
