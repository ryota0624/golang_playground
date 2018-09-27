# golang_playground
遊ぶようのリポジトリ

## やりたいこと
ロガーでstructを出してみる

cli作る用ライブラリを使ってみる

-サブコマンドを定義してみる
-引数で日付渡して適切なエラーを返したい

ストリーム処理をしてみる

sqlいじる

bqのライブラリいじる

- object名の頭に/つけるとダメ

- cliでのコレライブラリだと〜をまとめる

templateでsqlを作る

vgoなにできる？のdepとちゃうん？

日付操作ライブラリを探す



# start
go get github.com/rakyll/statik
statik -src=static

# Logger
```
	logger.Info("INFO") // Infoでる
	logger.Debug("DEBUG") // 本番だと出力されない
	logger.Error("ERROR") // stack trace出るけどプログラムは止まらない exit status 1
	logger.Fatal("Fatal") //プログラム止まる exit status 1
	logger.Warn("WARN")   // 開発だとstack traceが出る
	logger.Panic("PANIC") // 本番でもstacktraceが詳細に出る exit status 2
```
