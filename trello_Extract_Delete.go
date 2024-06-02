// trello_Extract_Delete.go
// GOOS=windows GOARCH=amd64 go build -o trello_Extract_Delete.exe trello_Extract_Delete.go
// GOOS=linux GOARCH=amd64 go build -o trello_Extract_Delete trello_Extract_Delete.go
// Please set it as an environment variable.
// ENV_TRELLO_APPKEY
// ENV_TRELLO_TOKEN
// ENV_TRELLO_USERNAME
package main

import (
	"github.com/adlio/trello"
	"github.com/caarlos0/env/v10"
	"fmt"
	"os"
	"flag"
)

func main() {
	var (
		TargetCardName = flag.String("CARD_NAME", "", "string")
	)
	flag.Parse();
	fmt.Println("TargetCardName：", *TargetCardName);
	type config struct {
		AppKey string `env:"ENV_TRELLO_APPKEY"`
		Token string `env:"ENV_TRELLO_TOKEN"`
		UserName string `env:"ENV_TRELLO_USERNAME"`
	}
	
	cfg := config{}
	
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
		os.Exit(99);
	}
	
	client := trello.NewClient(cfg.AppKey, cfg.Token);
	member, err := client.GetMember(cfg.UserName, trello.Defaults());
	
	if err != nil {
		panic(err);
		os.Exit(99);
	}
	_ = member
	
	boards, err := member.GetBoards(trello.Defaults());
	
	if err != nil {
		// Handle error
	}
	
	for i, BoardData := range boards {
		fmt.Println("BOARD DATA");
		fmt.Println("-------------------------------------------");
		fmt.Println(i, BoardData.ID, BoardData.Name);
		fmt.Println("-------------------------------------------");
		lists, err := BoardData.GetLists(trello.Defaults());
		
		if err != nil {
			panic(err);
			os.Exit(99);
		}
		
		for j, ListData := range lists {
			fmt.Println("LIST DATA START");
			fmt.Println("-------------------------------------------");
			fmt.Println(j, ListData.ID, ListData.Name);
			cards, err := ListData.GetCards(trello.Defaults());
			
			if err != nil {
				panic(err);
				os.Exit(99);
			}
			
			fmt.Println("CARD DATA START");
			fmt.Println("-------------------------------------------");
			
			for k, CardData := range cards {
				fmt.Println(k, CardData.ID, CardData.Name);
				if CardData.Name == *TargetCardName {
					err := CardData.Delete();
					if err != nil {
						panic(err);
						os.Exit(99);
					}
				}
			}
			
			fmt.Println("-------------------------------------------");
			fmt.Println("CARD DATA END");
			fmt.Println("-------------------------------------------");
			fmt.Println("LIST DATA END");
		}
		
		fmt.Println("-------------------------------------------");
	}
	
	os.Exit(0);
}

