package stoprdsdbinstance

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"fmt"
)


var log = logger.GetLogger("activity-tibco-stoptrdsdbinstance")

const (
	ivRegion    					= "region"
	ivAccessKey 					= "accessKey"
	ivSecretKey 					= "secretKey"
	ivDBInstanceIdentifier 			= "dbInstanceIdentifier"
	ivDBSnapshotIdentifier 			= "dbSnapshotIdentifier"

	ovDBResponse	    		= "response"
	ovDBInstanceStatus	   		= "status"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	if len(context.GetInput(ivAccessKey).(string)) < 1 || len(context.GetInput(ivSecretKey).(string)) < 1 || len(context.GetInput(ivRegion).(string)) < 1 || len(context.GetInput(ivDBInstanceIdentifier).(string)) < 1 {
		log.Error("Required variables have not been set !")
		return false, fmt.Errorf("required variables have not been set")
	}
	
	//Creating the session
	var accessKey, secretKey = "", ""
	if context.GetInput(ivAccessKey) != nil {
		accessKey = context.GetInput(ivAccessKey).(string)
	}
	if context.GetInput(ivSecretKey) != nil {
		secretKey = context.GetInput(ivSecretKey).(string)
	}
	var config *aws.Config
	region := context.GetInput(ivRegion).(string)
	if accessKey != "" && secretKey != "" {
		config = aws.NewConfig().WithRegion(region).WithCredentials(credentials.NewStaticCredentials(accessKey, secretKey, ""))
	} else {
		config = aws.NewConfig().WithRegion(region)
	}
	log.Debug("Session created")

	svc := rds.New(session.New(config))

	log.Debug("Setting StopDBInstanceInput parameters")
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(context.GetInput(ivDBInstanceIdentifier).(string)),
	}
	//Set the parameter if only it is present else null
	if len(context.GetInput(ivDBSnapshotIdentifier).(string))  >  1{
		input.DBSnapshotIdentifier = aws.String(context.GetInput(ivDBSnapshotIdentifier).(string))
	}

	result, err := svc.StopDBInstance(input)

	if err != nil {
		log.Error(err)
		return false, err
	}
	//On success Create Response in the structure.

	log.Debugf("RDS DB Instance successfully Stopped With DBInstanceIdentifier : ",result.DBInstance.DBInstanceIdentifier,"\n and Response : \n",result.DBInstance)
	context.SetOutput(ovDBResponse,result.DBInstance)
	context.SetOutput(ovDBInstanceStatus,"Success")

	return true, nil
}
