package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/keyauth/apps/user"
	"github.com/lifangjunone/keyauth/conf"
	"github.com/lifangjunone/keyauth/service_registry"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	srv = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	user.UnimplementedServiceServer
}

func (s *impl) Name() string {
	return user.AppName
}

func (s *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *impl) Registry(server *grpc.Server) {
	user.RegisterServiceServer(server, srv)
}

func init() {
	service_registry.RegistryGrpcApp(srv)
}
