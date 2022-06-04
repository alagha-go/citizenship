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
	IDs []Identity
)

func Main() {
	LoadSecret()
	LoadIDs()
	ConnectToDB()
}

func LoadSecret() {
	data, _ := ioutil.ReadFile("./secret.json")
	json.Unmarshal(data, &SecretData)
}

func LoadIDs() {
	data, _ := ioutil.ReadFile("./DB/ids.json")
	json.Unmarshal(data, &IDs)
}

func (Secret *Secret) Save() {
	data := JsonMarshal(Secret)
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