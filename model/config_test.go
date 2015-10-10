package model

import(
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := GetConfig()

	if conf.BackupPath != "C:\\Some\\Path\\To\\Dropbox" {
		t.Errorf("Fail to read conf: %s is not equal to C:\\Some\\Path\\To\\Dropbox", conf.BackupPath)
	}

	if conf.CheckInterval != 2 {
		t.Errorf("Failing assert that %d is not equal to 2", conf.CheckInterval)
	}
}