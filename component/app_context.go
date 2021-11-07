package component

import (
	"food_delivery_be/component/uploadprovider"
	"food_delivery_be/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
	pb         pubsub.Pubsub
}

func NewAppContent(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, pb pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey, pb: pb}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.pb
}
