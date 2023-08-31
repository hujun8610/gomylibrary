package net

import "testing"

func TestPostFormRequest(t *testing.T) {
	url := "https://dev.coc.10086.cn/staging-coc/gateway/api/qimen/user/smsCode?mobile=15967116560&app_key=23558957&target_appkey=32502322&format=json&ext=%7B%22templateNo%22%3A%224641%22%2C%22smsCodeTypeEnum%22%3A%22503%22%2C%22touchCode%22%3A%22P00000016463%22%2C%22pageId%22%3A%221650310735709323264%22%7D&timestamp=2023-08-31+18%3A02%3A43&product_code=100000094566&method=qimen.alibaba.alicom.opentrade.identifycode.get&sign_method=md5&source_appkey=23558957&request_id=15sxrkkqnhclo&sign=FFCDED6779F842F8AF73A3A0416A3605"
	result := PostFormRequest(url)
	log.WithField("response", result).Info("result")
}
