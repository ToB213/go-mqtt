## Go言語を使用したMosquitto MQTTパブリッシャーとサブスクライバー
このリポジトリは，Go言語を使ってMosquittoブローカーと通信するMQTTパブリッシャーとサブスクライバーのサンプルプログラムです．

## 前提条件
MacBook M1
Goのインストール（推奨バージョン：v1.16以降）
Mosquittoのインストールとローカル実行

## 1. MacBookにMosquittoをインストール

まだMosquittoがインストールされていない場合，Homebrewを使用して以下のコマンドでインストールします：

```
brew install mosquitto
```

次に，Mosquittoを起動します：


```
mosquitto
```

このコマンドで，Mosquittoが正常に起動しているか確認できます．ローカルホスト（localhost:1883）で動作していることが確認できます．

## 2. Goのインストール

Goがまだインストールされていない場合は，Goの公式サイトから最新バージョンをダウンロードしてインストールしてください．

インストールが完了したら，以下のコマンドでインストールが成功しているか確認します：

```
go version
```


## 3. Go用のMQTTクライアントライブラリのインストール

次に，GoでMQTT通信を行うために，Eclipse Paho MQTTクライアントライブラリをインストールします．以下のコマンドを実行します：

```
go get github.com/eclipse/paho.mqtt.golang
```


## プロジェクト構成

```
tree .
.
├── publisher
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── subscriber
    ├── go.mod
    ├── go.sum
    └── main.go
```

## 実行方法
### 1. サブスクライバーの実行

まず，メッセージを受信するサブスクライバーを起動します．以下のコマンドをターミナルで実行してください：

```
go run subscriber/main.go
```

これにより，サブスクライバーがtest/topicというトピックで待機し，パブリッシャーからのメッセージを受信する準備が整います．

### 2. パブリッシャーの実行
次に，別のターミナルを開き，パブリッシャーを実行してメッセージを送信します．以下のコマンドを実行します：

```
go run publisher/publisher.go
```

パブリッシャーは，test/topicというトピックに対して10回メッセージを送信します．

## 実行結果

### サブスクライバーの出力例

サブスクライバーを実行し，パブリッシャーからメッセージが送信されると，次のような出力が表示されます：

```
接続成功
購読成功
受信: トピック [test/topic] メッセージ: Hello, World! 1
受信: トピック [test/topic] メッセージ: Hello, World! 2
受信: トピック [test/topic] メッセージ: Hello, World! 3
受信: トピック [test/topic] メッセージ: Hello, World! 4
受信: トピック [test/topic] メッセージ: Hello, World! 5
受信: トピック [test/topic] メッセージ: Hello, World! 6
受信: トピック [test/topic] メッセージ: Hello, World! 7
受信: トピック [test/topic] メッセージ: Hello, World! 8
受信: トピック [test/topic] メッセージ: Hello, World! 9
受信: トピック [test/topic] メッセージ: Hello, World! 10
...
```

### パブリッシャーの出力例

パブリッシャー側では，次のようなメッセージ送信のログが表示されます：

```
接続成功
メッセージ送信:  Hello, World! 1
メッセージ送信:  Hello, World! 2
メッセージ送信:  Hello, World! 3
メッセージ送信:  Hello, World! 4
メッセージ送信:  Hello, World! 5
メッセージ送信:  Hello, World! 6
メッセージ送信:  Hello, World! 7
メッセージ送信:  Hello, World! 8
メッセージ送信:  Hello, World! 9
メッセージ送信:  Hello, World! 10
切断
```
