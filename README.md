<p align="center">
    <a href="https://opensource.org/licenses/BSD-3-Clause"><img src="https://img.shields.io/badge/license-bsd-orange.svg" alt="Licenses"></a>
</p>

# trello_golang
This is an access tool for trello using the adlio/trello library.  

# Introduction
Copyright 2024 The trello_golang Project Authors, GinSanaduki.
All rights reserved.

trello-cliがいつの間にかシンタックスエラーなどでまともに動かなくなっていたので、
golangのapiで接続できるようにしたもの。  
Since trello-cli had stopped working properly due to syntax errors,  
I made it possible to connect using the golang API.  
The following librarys need to be installed.  
* adlio/trello
```
go get github.com/adlio/trello
```
https://github.com/adlio/trello
* caarlos0/env
```
go get github.com/caarlos0/env/v10
```
https://github.com/caarlos0/env

# 環境変数
以下の環境変数を設定する必要があります。  
The following environment variables must be set:  
* ENV_TRELLO_APPKEY
* ENV_TRELLO_TOKEN
* ENV_TRELLO_USERNAME

Trello API KeyとTokenの取得 #API - Qiita  
https://qiita.com/HiguchiTakahiro/items/94f52eee07d257f8f995  
を参考にAPI KeyとTokenを取得し、  
それぞれをENV_TRELLO_APPKEYとENV_TRELLO_TOKENに設定してください。  
ENV_TRELLO_USERNAMEには、メンバーとして記載されている、
@から始まるメンバーIDの、先頭の@を除いた値を設定します。

Obtaining Trello API Key and Token #API - Qiita  
https://qiita.com/HiguchiTakahiro/items/94f52eee07d257f8f995  
Refer to obtain an API Key and Token,  
and set them to ENV_TRELLO_APPKEY and ENV_TRELLO_TOKEN respectively.  
For ENV_TRELLO_USERNAME, set the member ID listed as a member,  
starting with @, without the leading @.  

```
#!/bin/sh
# trello.sh
# sh trello.sh

export ENV_TRELLO_APPKEY="API Key"
export ENV_TRELLO_TOKEN="Token"
export ENV_TRELLO_USERNAME="jhondoe"

./trello_Extract
./trello_AddCard -BOARD_NAME=BOARD_NAME -LIST_NAME=LIST_NAME -CARD_NAME=CARD_NAME
./trello_DeleteCard -BOARD_NAME=BOARD_NAME -LIST_NAME=LIST_NAME -CARD_NAME=CARD_NAME
./trello_AddList -BOARD_NAME=BOARD_NAME -LIST_NAME=LIST_NAME
./trello_Extract_Delete -CARD_NAME=CARD_NAME
```

# ビルド例
```
GOOS=windows GOARCH=amd64 go build -o trello_AddCard.exe trello_AddCard.go
GOOS=linux GOARCH=amd64 go build -o trello_AddCard trello_AddCard.go
```

# trello_Extract
ENV_TRELLO_USERNAMEで設定されたメンバのボード内にあるリストとカードをすべて表示します。  
引数：なし  
Displays all lists and cards in the board of the member specified by ENV_TRELLO_USERNAME.  
Arguments: None  

# trello_AddCard
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリストに、引数のCARD_NAMEで設定されたカード名のカードを追加します。  
カードのオプションは個人的にデフォルトで困らなかったので設定していません。使いたかったら  
https://pkg.go.dev/github.com/adlio/trello   
を読んで自分で対応してください。  

Adds a card with the card name specified by the CARD_NAME argument to the list specified by the LIST_NAME argument, on the board specified by the BOARD_NAME argument, of the member specified by ENV_TRELLO_USERNAME.  
I personally didn't have any problems with the default card options, so I didn't set them. If you want to use them, please read  
https://pkg.go.dev/github.com/adlio/trello  
and handle it yourself.  

# trello_DeleteCard
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリストに、引数のCARD_NAMEで設定されたカード名のカードをすべて削除します。  
This command deletes all cards with the card name specified by the CARD_NAME argument from the list specified by the LIST_NAME argument, in the board specified by the BOARD_NAME argument, of the member specified by ENV_TRELLO_USERNAME.  

# trello_AddList
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリスト名のリストを追加します。  
リストのオプションは個人的にデフォルトで困らなかったので設定していません。使いたかったら  
https://pkg.go.dev/github.com/adlio/trello    
を読んで自分で対応してください。   
リストのアーカイブのメソッドはきっとどこかにあるんでしょう、たぶん・・・。  

Adds a list of the list name specified by the LIST_NAME argument, in the board specified by the BOARD_NAME argument, of the members specified by ENV_TRELLO_USERNAME.  
I personally didn't have any problems with the default list options, so I didn't set them. If you want to use them, please read  
https://pkg.go.dev/github.com/adlio/trello  
and handle it yourself.  
There's probably a method for archiving lists somewhere, probably...  

# trello_Extract
ENV_TRELLO_USERNAMEで設定されたメンバのボード内にあるリストとカードを探索し、引数のCARD_NAMEで設定されたカード名のカードをすべて削除します。  
It searches the lists and cards in the board of the member specified by ENV_TRELLO_USERNAME, and deletes all cards with the card name specified by the argument CARD_NAME.  
