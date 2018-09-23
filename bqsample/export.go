package bqsample

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
)

func ExportGCS() {
	log.Println("Start ExportGCS")

	ctx := context.Background()
	client, error := bigquery.NewClient(ctx, "handson-1061")

	if error != nil {
		panic(error)
	}

	tmpTable := client.Dataset("master").Table("tmp_advertiser_for_gcs")
	query := client.Query("select name from master.advertiser limit 2;")
	query.AllowLargeResults = true
	query.Dst = tmpTable // Dstはdestinationtable

	_, queryError := query.Run(ctx)

	if queryError != nil {
		log.Printf("Failed to Run Query Job:%v\n", queryError.Error())
	}

	gcsRef := bigquery.NewGCSReference(os.Getenv("CSV_GCS_PATH"))
	gcsRef.Compression = bigquery.Gzip
	gcsRef.DestinationFormat = bigquery.CSV

	extractor := tmpTable.ExtractorTo(gcsRef)
	extractor.ExtractConfig.DisableHeader = true
	_, extractError := extractor.Run(ctx)

	if extractError != nil {
		log.Printf("Failed to extract Job:%v\n", extractError.Error())
	}
	log.Println("End ExportGCS")

}
