package service

import (
	"context"
	"fmt"
	"log"
	"main/model"
	"mime/multipart"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

var conf model.QNconfig

func Init() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	conf = model.QNconfig{
		AccessKey: viper.GetString("oss.access_key"),
		SecretKey: viper.GetString("oss.secret_key"),
		Bucket:    viper.GetString("oss.bucket_name"),
		Domain:    viper.GetString("oss.domain_name"),
	}
}

func UploadProfilePhoto(id int, file multipart.File, size int64) error {
	keyToOverwrite := strconv.Itoa(id) + ".jpg"
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", conf.Bucket, keyToOverwrite),
	}
	mac := qbox.NewMac(conf.AccessKey, conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	})

	ret := storage.PutRet{}

	if err := formUploader.Put(context.Background(), &ret,
		upToken, keyToOverwrite, file, size, nil); err != nil {
		return err
	}
	return nil
}
