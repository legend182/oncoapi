package mongodb

import (
	"context"
	"oncoapi/setting"

	"github.com/qiniu/qmgo"
)

var mg *qmgo.Client
var coll *qmgo.Collection
func Init(cfg *setting.MongodbConf)(err error) {
	ctx := context.Background()
	uri := "mongodb://"+cfg.Host+":"+cfg.Host
	mg,err = qmgo.NewClient(ctx,&qmgo.Config{Uri: uri})
	coll = mg.Database(cfg.DB).Collection(cfg.Collection)
	return
}

func Close() {
	_ = mg.Close(context.Background())
}