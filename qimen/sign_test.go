package qimen

import "testing"

func TestSignTopRequest(t *testing.T) {
	params := make(map[string]string)
	params["key"] = "value"
	secret := "919fa294a24538d2ae685aaf849d0a67"
	result := SignTopRequest(params, "", secret, "md5")
	expect := "7E57E0619CD614ACEBF2876C4705A66D"
	if expect != result {
		t.Errorf("result %s,expect %s", secret, expect)
	}

}
