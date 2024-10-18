package main

import (
	"testing"
	"v2lbsyun/config"
	"v2lbsyun/public/utils"
	v2lcontext "v2lbsyun/v2l_context"
)

func TestGetDesEncrypt(t *testing.T) {
	config.InitConfig()
	t.Logf("config: %v\n", config.Cfg)
	enscrypt := utils.GetDesEncrypt("yanxianhe")
	t.Logf("encrypt: %s\n", enscrypt)
	decrypt := utils.GetDesDecrypt(enscrypt)
	t.Logf("decrypt: %s\n", decrypt)
}

func TestPing(t *testing.T) {
	v2lcontext.Ping(nil)
	t.Logf("test ping:\n")
}
