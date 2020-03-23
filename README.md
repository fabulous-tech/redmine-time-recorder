# Redmine time recorder

Redmine time recorderは作業時間を対話式cliで登録するためのツールです。

## 使い方。
実行時の必須flag
```cassandraql
-e redmineのエンドポイントを指定します。
-k redmineのAPI Keyを指定します。
```

実行すると対話形式で作業時間を登録したいプロジェクトやチケットを確認していくので指示に従って入力していきます。

## 注意

現在、とりあえず動くバージョンです。  
入力値のエラーチェックやチケットが多いときのページネーションにも対応していません。

## Licence
This software is released under the MIT License, see LICENSE.txt.