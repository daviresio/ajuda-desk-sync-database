package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/daviresio/sync-database/model"
	"github.com/olivere/elastic/v7"
	"strconv"
)

var (
	esclient *elastic.Client
)

func init() {
	var err error
	esclient, err = getESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic(err)
	}
}

func getESClient() (*elastic.Client, error) {

	client, err :=  elastic.NewClient(elastic.SetURL("http://35.209.165.160:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}


func InsertData(session sarama.ConsumerGroupSession, message **sarama.ConsumerMessage) {

	var genericData model.GerenicData
	var idData model.IdData

	err := json.Unmarshal((*message).Value, &genericData)

	if err != nil {
		println(err)
		panic(err)
	}

	err = json.Unmarshal((*message).Value, &idData)

	if err != nil {
		println(err)
		panic(err)
	}


	ctx := context.Background()

	v, err := esclient.Index().
		Index(genericData.Payload.Source.Table).
		Id(strconv.Itoa(idData.Payload.After.Id)).
		BodyJson(genericData.Payload.After).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(genericData.Payload.After))
	fmt.Println(v)

	session.MarkMessage(*message, "")


}
