// trello_AddCard.go
// GOOS=windows GOARCH=amd64 go build -o trello_AddCard.exe trello_AddCard.go
// GOOS=linux GOARCH=amd64 go build -o trello_AddCard trello_AddCard.go
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
		TargetBoardName = flag.String("BOARD_NAME", "", "string")
		TargetListName = flag.String("LIST_NAME", "", "string")
		TargetCardName = flag.String("CARD_NAME", "", "string")
	)
	flag.Parse();
	fmt.Println("TargetBoardName：", *TargetBoardName);
	fmt.Println("TargetListName：", *TargetListName);
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
		panic(err);
		os.Exit(99);
	}
	
	IsExistTargetBoardData := false;
	TargetBoardID := "";
	
	for _, BoardData := range boards {
		if BoardData.Name == *TargetBoardName {
			IsExistTargetBoardData = true;
			TargetBoardID = BoardData.ID;
			break;
		}
	}
	
	if !IsExistTargetBoardData {
		fmt.Println("TargetBoardName IS NOT EXISTS");
		os.Exit(99);
	}
	
	fmt.Println("TargetBoardID：", TargetBoardID);
	
	board, err := client.GetBoard(TargetBoardID, trello.Defaults());
	
	if err != nil {
		panic(err);
		os.Exit(99);
	}
	
	lists, err := board.GetLists(trello.Defaults());
	
	if err != nil {
		panic(err);
		os.Exit(99);
	}
	
	IsExistTargetListData := false;
	TargetListID := "";
	
	for _, ListData := range lists {
		if ListData.Name == *TargetListName {
			IsExistTargetListData = true;
			TargetListID = ListData.ID;
			break;
		}
	}
	
	if !IsExistTargetListData {
		fmt.Println("TargetListName IS NOT EXISTS");
		os.Exit(99);
	}
	
	fmt.Println("TargetListID：", TargetListID);
	
	list, err := client.GetList(TargetListID, trello.Defaults());
	
	if err != nil {
		panic(err);
		os.Exit(99);
	}
	
	list.AddCard(&trello.Card{ Name: *TargetCardName }, trello.Defaults());
	fmt.Println("CURRENT LIST STATE...");
	cards, err := list.GetCards(trello.Defaults());
	fmt.Println("-------------------------------------------");
	
	for k, CardData := range cards {
		fmt.Println(k, CardData.ID, CardData.Name);
	}
	
	fmt.Println("-------------------------------------------");
	
	os.Exit(0);
}

