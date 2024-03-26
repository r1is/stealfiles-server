package utils

import (
	"stealfiles-server/common"
	"time"

	"encoding/json"

	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

func Sts() (string, error) {
	// 临时密钥服务配置信息
	appid := common.ServerCfg.APPID
	bucket := common.ServerCfg.BUCKET
	region := common.ServerCfg.REGION
	filepath := common.ServerCfg.OSSFILEPATH
	tak := common.ServerCfg.SECRETID
	tsk := common.ServerCfg.SECRETKEY

	c := sts.NewClient(
		tak,
		tsk,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
						//下载
						"name/cos:GetObject",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + filepath,
					},
					// 开始构建生效条件 condition
					// 关于 condition 的详细设置规则和COS支持的condition类型可以参考https://cloud.tencent.com/document/product/436/71306
				},
			},
		},
	}

	// case 1 请求临时密钥
	res, err := c.GetCredential(opt)
	if err != nil {
		return "", err
	}
	// fmt.Println(res)
	_res := &common.OssCfg{
		TmpSecretID:  res.Credentials.TmpSecretID,
		TmpSecretKey: res.Credentials.TmpSecretKey,
		SessionToken: res.Credentials.SessionToken,
		BucketURL:    "https://cos." + region + ".myqcloud.com",
		BatchURL:     "https://" + bucket + ".cos." + region + ".myqcloud.com",
	}
	strJson, _ := json.Marshal(_res)
	return string(strJson), nil
}
