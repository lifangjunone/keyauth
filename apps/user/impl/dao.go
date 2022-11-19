package impl

import (
	"context"
	"github.com/lifangjunone/keyauth/apps/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func save(col *mongo.Collection, ctx context.Context, ins *user.User) error {
	_, err := col.InsertOne(ctx, ins)
	if err != nil {
		return err
	}
	return nil
}
