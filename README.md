# BotwayGo

> Golang client package for Botway.


## Installation

```bash
go get github.com/abdfnx/botwaygo
```

## Usage

> this is an example of botway slack go template

```go
package main

import (
	"context"
	"log"

	"github.com/abdfnx/botwaygo"
	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient(botwaygo.GetToken(), botwaygo.GetAppToken())

	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
```
