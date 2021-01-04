package ox

var mainTemplate = `
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/wawandco/oxpecker/cli"
	wawandco "github.com/wawandco/oxplugins"
)

func main() {
  	fmt.Print("[info] Using {{.Name}}/cmd/ox \n\n")
	ctx := context.Background()
    
  	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
  	}
    
	cl := cli.New()
	// append your plugins here
	cl.Plugins = append(cl.Plugins, wawandco.All...)
    
    err = cl.Run(ctx, pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
`
