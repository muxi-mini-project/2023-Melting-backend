package service

import (
	"context"
	"fmt"
	"main/model"
	"main/model/db"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var conf model.QNconfig

func Init() {
	db.OpenDB()
	conf = model.QNconfig{
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
		Bucket:    os.Getenv("bucket_name"),
		Domain:    os.Getenv("domain_name"),
	}
}

func UploadProfilePhoto(id int, file multipart.File, size int64) (string, error) {
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
		return "", err
	}
	return conf.Domain + "/" + ret.Key, nil
}
