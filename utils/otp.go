package utils

import (
	"stealfiles-server/common"

	"github.com/pquerna/otp/totp"
)

func Validate(timecode string) bool {
	return totp.Validate(timecode, common.ServerCfg.TOTPKEY)
}
