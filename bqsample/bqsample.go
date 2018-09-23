package bqsample

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func Do() {
	ctx := context.Background()
	client, error := bigquery.NewClient(ctx, "handson-1061")

	if error != nil {
		panic(error)
	}

	query := client.Query("select name from master.advertiser;")

	it, error := query.Read(ctx)

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
