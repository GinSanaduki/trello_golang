// trello_AddList.go
// GOOS=windows GOARCH=amd64 go build -o trello_AddList.exe trello_AddList.go
// GOOS=linux GOARCH=amd64 go build -o trello_AddList trello_AddList.go
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
	)
	flag.Parse();
	fmt.Println("TargetBoardName：", *TargetBoardName);
	fmt.Println("TargetListName：", *TargetListName);
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
	
	for _, ListData := range lists {
		if ListData.Name == *TargetListName {
			IsExistTargetListData = true;
			break;
		}
	}
	
	if IsExistTargetListData {
		fmt.Println("TargetListName IS ALREADY EXISTS");
		os.Exit(99);
	}
	
	newLists, err := board.CreateList(*TargetListName);
	
	if err != nil {
		panic(err);
		os.Exit(99);
	}
	
	_ = newLists;
	
	fmt.Println("CURRENT BOARD STATE...");
	fmt.Println("-------------------------------------------");
	
	lists, err = board.GetLists(trello.Defaults());
	
	for j, ListData := range lists {
		fmt.Println(j, ListData.ID, ListData.Name);
	}
	
	fmt.Println("-------------------------------------------");
	
	os.Exit(0);
}

