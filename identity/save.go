package identity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (ID *Identity) Svae() {
	ctx := context.Background()
	collection := Client.Database("Citizenship").Collection("Identities")
	ID.ID = primitive.NewObjectID()
	IDs = append(IDs, *ID)
	SvaeIDs

	collection.InsertOne(ctx, ID)
}