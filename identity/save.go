package identity

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (ID *Identity) Save() {
	ctx := context.Background()
	collection := Client.Database("Citizenship").Collection("Identities")
	ID.ID = primitive.NewObjectID()
	IDs = append(IDs, *ID)
	SaveIDs()

	collection.InsertOne(ctx, ID)
}

func SaveIDs() {
	data, err := json.Marshal(IDs)
	handleError(err)
	ioutil.WriteFile("./DB/ids.json", data, 0755)
}