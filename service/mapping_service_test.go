package service

import (
	"github.com/cobre420/mapper-go/domain"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_ProcessMappingGet(t *testing.T) {
	ProcessMapping(
		"../resource/_test/mapping-001.yaml",
		"../resource/_test/cfg.json",
		initStateVector(t, "../resource/_test/initial-vector-001.json"),
	)
}

func Test_ProcessMappingPut(t *testing.T) {
	ProcessMapping(
		"../resource/_test/mapping-002.yaml",
		"../resource/_test/cfg.json",
		initStateVector(t, "../resource/_test/initial-vector-002.json"),
	)
}

func initStateVector(t *testing.T, file string) *domain.StateVector {
	sv := domain.NewStateVector()

	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	sv.Vector("main").Text = string(data)

	return &sv
}
