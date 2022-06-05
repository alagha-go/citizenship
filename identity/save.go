package identity

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (ID *Identity) Save() {
	if ID.Exists() {
		return
	}
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


func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}