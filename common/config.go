package common

import "os"

type Args struct {
	Sm4key   string
	FileName string
}

// 示例：{"ak":"xxxxx","sk":"bbbb","BucketURL":"https://cos.ap-beijing.myqcloud.com","BatchURL":"https://xxxxx1xx.cos.ap-beijing.myqcloud.com"}
type OssCfg struct {
	TmpSecretID  string `json:"TmpSecretId,omitempty"`
	TmpSecretKey string `json:"TmpSecretKey,omitempty"`
	SessionToken string `json:"Token,omitempty"`
	BucketURL    string `json:"BucketURL"`
	BatchURL     string `json:"BatchURL"`
}

var ContectKey = "10086.com"

type Config struct {
	TOTPKEY     string
	SECRETID    string
	SECRETKEY   string
	REGION      string
	BUCKET      string
	APPID       string
	OSSFILEPATH string
}

// 全局配置信息
var ServerCfg Config

func init() {
	ServerCfg = Config{
		TOTPKEY:     os.Getenv("TOTPKEY"),
		SECRETID:    os.Getenv("SECRETID"),
		SECRETKEY:   os.Getenv("SECRETKEY"),
		REGION:      os.Getenv("REGION"),
		BUCKET:      os.Getenv("BUCKET"),
		APPID:       os.Getenv("APPID"),
		OSSFILEPATH: os.Getenv("OSSFILEPATH"),
	}
}
