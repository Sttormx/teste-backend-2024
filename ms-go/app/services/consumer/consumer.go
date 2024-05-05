package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-go/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Handler(message []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(message, &data); err != nil {
		fmt.Println("error:", err)
		return
	}

	db.Connection().UpdateOne(context.TODO(), bson.M{"id": data["id"]}, bson.M{"$set": data}, options.Update().SetUpsert(true))
	defer db.Disconnect()
}
