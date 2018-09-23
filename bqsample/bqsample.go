package bqsample

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// テーブル作るくん
func makeTmp(ctx context.Context, client *bigquery.Client) error {
	return client.Dataset("master").Table("tmp_advertiser").Create(ctx, nil)
}

func Do() {
	ctx := context.Background()
	client, error := bigquery.NewClient(ctx, "handson-1061")

	if error != nil {
		panic(error)
	}

	// error = makeTmp(ctx, client)
	// if error != nil {
	// 	panic(error)
	// }

	query := client.Query("select name from master.advertiser limit 2;")
	// query.AllowLargeResults = true
	query.WriteDisposition = bigquery.WriteTruncate // 書き込み先にデータがあった場合はoverrideする
	// query.CreateDisposition // クエリー結果テーブルを作成するか

	// query.DryRun = true
	query.Dst = client.Dataset("master").Table("tmp_advertiser") // Dstはdestinationtable

	job, error := query.Run(ctx)

	if error != nil {
		log.Println("Failed to Run Query Job:%v", error)
	}

	it, error := job.Read(ctx)

	if error != nil {
		log.Println("Failed to Read Query:%v", error)
	}

	for {
		// BigQueryの結果から、中身を格納するためのBigQuery.Valueのsliceを宣言
		// BigQuery.Valueはinterface{}型
		var values []bigquery.Value

		// 引数に与えたvaluesにnextを格納する
		// Iteratorを返す
		// これ以上結果が存在しない場合には、iterator.Doneを返す
		// iterator.Doneが返ってきたら、forを抜ける
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println("Failed to Iterate Query:%v", err)
		}

		fmt.Println(values)
	}
}
