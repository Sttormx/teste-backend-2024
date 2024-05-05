package products

import (
	"context"
	"encoding/json"
	"log"
	"ms-go/app/helpers"
	"ms-go/app/messager"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(data models.Product, isAPI bool) (*models.Product, error) {

	if data.ID == 0 {
		var max models.Product

		opts := options.FindOne()

		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

		db.Connection().FindOne(context.TODO(), bson.D{}, opts).Decode(&max)

		data.ID = max.ID + 1
	}

	if err := data.Validate(); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusUnprocessableEntity}
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = data.CreatedAt

	_, err := db.Connection().InsertOne(context.TODO(), data)

	if err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	defer db.Disconnect()

	if isAPI {
		msg, err := json.Marshal(&data)
		if err != nil {
			return nil, err
		}

		go func() {
			inst, err := messager.New().Setup(messager.TOPIC_GO_TO_RAILS, messager.PARTITION_DEFAULT).Write(msg)
			if err != nil {
				log.Println(err)
			}

			inst.Close()
		}()
	}

	return &data, nil
}
