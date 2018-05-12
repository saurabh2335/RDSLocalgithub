---
title: Stop AWS RDS Instance
---

# Stop RDS Instance
This activity allows you to Stop AWS RDS via DBInstanceIdentifier and provide the access key and secret for authentication.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/
```

## Schema
Inputs and Outputs:

```json
{
 "inputs":[
    {
      "name": "accessKey",
      "type": "string",
      "required": "true"
    },
    {
      "name": "secretKey",
      "type": "string",
      "required": "true"
    },
    {
      "name": "region",
      "type": "string",
      "required": "true",
      "allowed" : ["us-east-2","us-east-1","us-west-1","us-west-2","ap-south-1","ap-northeast-2","ap-southeast-1","ap-southeast-2","ap-northeast-1","cn-northwest-1","ca-central-1","eu-central-1","eu-west-1","eu-west-2","sa-east-1"]
    },
	{
      "name": "dbInstanceIdentifier",
      "type": "string",
      "required": "true"
    },
	{
      "name": "dbSnapshotIdentifier",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "response",
      "type": "string"
    },
	{
      "name": "status",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     				| Required | Description 																	|
|:----------------------	|:---------|:-------------------------------------------------------------------------------|
| region      				| True     | The AWS region in which you want to invoke the function 						|
| accessKey   				| True     | AWS access key for the user to invoke the function 							|
| secretKey   				| True     | AWS secret key for the user to invoke te function 								|
| dbInstanceIdentifier  	| True     | String value dbInstanceIdentifier for existing AWS RDS Instance				|
| dbSnapshotIdentifier 		| False    | String value dbSnapshotIdentifier for existing AWS RDS Instance 				|
| response  				| False    | Response of STOP RDS DB Instance												|
| status      				| False    | Current status of RDS DB Instance												|

## Examples
Coming soon...