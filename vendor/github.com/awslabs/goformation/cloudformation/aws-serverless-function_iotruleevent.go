package cloudformation

import (
	"encoding/json"
)

// AWSServerlessFunction_IoTRuleEvent AWS CloudFormation Resource (AWS::Serverless::Function.IoTRuleEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
type AWSServerlessFunction_IoTRuleEvent struct {

	// AwsIotSqlVersion AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
	AwsIotSqlVersion *Value `json:"AwsIotSqlVersion,omitempty"`

	// Sql AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
	Sql *Value `json:"Sql,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_IoTRuleEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.IoTRuleEvent"
}

func (r *AWSServerlessFunction_IoTRuleEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
