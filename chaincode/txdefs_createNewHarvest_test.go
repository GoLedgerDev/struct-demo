package main_test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/goledgerdev/cc-tools/mock"
	cc "github.com/goledgerdev/struct-demo/chaincode"
)

func TestCreateNewHarvest(t *testing.T) {
	stub := mock.NewMockStub("org3MSP", new(cc.CCDemo))

	expectedResponse := map[string]interface{}{
		"@key":         "harvest:59d6420c-9df2-5ea4-9e1c-6f3f7801d231",
		"@lastTouchBy": "org3MSP",
		"@lastTx":      "createNewHarvest",
		"@assetType":   "harvest",
		"date":         "2002-10-02T12:00:00Z",
		"harvestTeam":  []interface{}{"Team A"},
		"type":         1.0,
		"used":         false,
	}
	req := map[string]interface{}{
		"date":        "2002-10-02T12:00:00Z",
		"type":        1,
		"harvestTeam": []interface{}{"Team A"},
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		t.FailNow()
	}

	res := stub.MockInvoke("createNewHarvest", [][]byte{
		[]byte("createNewHarvest"),
		reqBytes,
	})

	if res.GetStatus() != 200 {
		log.Println(res)
		t.FailNow()
	}

	var resPayload map[string]interface{}
	err = json.Unmarshal(res.GetPayload(), &resPayload)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(resPayload, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", resPayload)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}

	var state map[string]interface{}
	stateBytes := stub.State["harvest:59d6420c-9df2-5ea4-9e1c-6f3f7801d231"]
	err = json.Unmarshal(stateBytes, &state)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(state, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", state)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}
}
