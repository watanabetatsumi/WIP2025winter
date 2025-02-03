# README

## プロジェクト概要
このディレクトリはTCP/IPコマンドの自作・実装用です

## ディレクトリ構成
```
/home/tatsumi/WIP2025winter/TCPIP/
├── cmd/        # mainファイル/実行可能ファイル
├── entity/       # 各プロトコルのヘッダー等の要件
└── protocol/     # 各プロトコルのロジック実装
```

## 使用言語
- Golang

## 使い方
1. リポジトリをクローンします。
    ```bash
    $ git clone https://github.com/watanabetatsumi/WIP2025winter
    ```
2. cmdディレクトリに移動します
    ```bash
    $ cd ./TCPIP/cmd
    ```
3. これらの実行可能ファイルをライブラリにコピーします。
    ```bash
    $ sudo cp ../<実行可能ファイル> /usr/local/bin 
    $ sudo chmod +x /usr/local/bin/<実行可能ファイル>
    ```
4. sudo権限をつけて実行します。
    ```bash
    $ sudo ft_arp 192.168.32.1
    ARP Reply : 0 15 5d b9 f6 78
    ```
