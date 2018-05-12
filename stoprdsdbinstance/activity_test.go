package stoprdsdbinstance

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"fmt"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	tc.SetInput("accessKey", "******************")
	tc.SetInput("secretKey", "******************")
	tc.SetInput("region", "ap-southeast-2")
	tc.SetInput("dbInstanceIdentifier", "flogordsinstance")
	tc.SetInput("dbSnapshotIdentifier", "")
	success, err := act.Eval(tc)

	if err != nil {
		t.Error("Error while terminating the instance")
		t.Fail()
		return
	}
	if success {
		response := tc.GetOutput("response")
		status := tc.GetOutput("status")
		fmt.Printf("Response ID : %v\n", response)
		fmt.Printf("status ID : %v\n", status)
	} else {
		t.Error("Unknown Error")
		t.Fail()
		return
	}
}
