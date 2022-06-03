package identity

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	SecretData Secret
	Client	*mongo.Client
)

func LoadSecret() {
	data, _ := ioutil.ReadFile("./secret.json")
	json.Unmarshal(data, &SecretData)
}

func (Secret *Secret) Save() {
	data, err := json.Marshal(Secret)
	handleError(err)
	ioutil.WriteFile("./secret.json", data, 0755)
}

func ConnectToDB() {
	var err error
	clientOptions := options.Client().
		ApplyURI(SecretData.MongoDBUrl)
	ctx := context.Background()
	
	Client, err = mongo.Connect(ctx, clientOptions)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}