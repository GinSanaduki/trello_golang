<p align="center">
    <a href="https://opensource.org/licenses/BSD-3-Clause"><img src="https://img.shields.io/badge/license-bsd-orange.svg" alt="Licenses"></a>
</p>

# trello_golang

# Introduction
Copyright 2024 The trello_golang Project Authors, GinSanaduki.
All rights reserved.

trello-cliがいつの間にかシンタックスエラーなどでまともに動かなくなっていたので、
golangのapiで接続できるようにしたもの。
The following librarys need to be installed.
* adlio/trello
```
go get github.com/adlio/trello
```
* caarlos0/env
```
go get github.com/caarlos0/env/v10
```

# 環境変数
以下の環境変数を設定する必要があります。  
* ENV_TRELLO_APPKEY
* ENV_TRELLO_TOKEN
* ENV_TRELLO_USERNAME

Trello API KeyとTokenの取得 #API - Qiita  
https://qiita.com/HiguchiTakahiro/items/94f52eee07d257f8f995  
を参考にAPI KeyとTokenを取得し、  
それぞれをENV_TRELLO_APPKEYとENV_TRELLO_TOKENに設定してください。  
ENV_TRELLO_USERNAMEには、メンバーとして記載されている、
@から始まるメンバーIDの、先頭の@を除いた値を設定します。
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

# trello_AddCard
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリストに、引数のCARD_NAMEで設定されたカード名のカードを追加します。
カードのオプションは個人的にデフォルトで困らなかったので設定していません。使いたかったら
https://pkg.go.dev/github.com/adlio/trello  
を読んで自分で対応してください。

# trello_DeleteCard
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリストに、引数のCARD_NAMEで設定されたカード名のカードをすべて削除します。

# trello_AddList
ENV_TRELLO_USERNAMEで設定されたメンバの、引数のBOARD_NAMEで設定されたボード内の、引数のLIST_NAMEで設定されたリスト名のリストを追加します。
リストのオプションは個人的にデフォルトで困らなかったので設定していません。使いたかったら
https://pkg.go.dev/github.com/adlio/trello  
を読んで自分で対応してください。
リストのアーカイブのメソッドはきっとどこかにあるんでしょう、たぶん・・・。

# trello_Extract
ENV_TRELLO_USERNAMEで設定されたメンバのボード内にあるリストとカードを探索し、引数のCARD_NAMEで設定されたカード名のカードをすべて削除します。

