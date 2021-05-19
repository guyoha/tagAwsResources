package resutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func UpdateEc2Tags(ses client.ConfigProvider, awsProfile string, awsRegion string, fp string, sheetName string, resourceIdName string) {

	ec2Svc := ec2.New(ses)
	ids := []*string{}
	ress := Resources{}
	acctName := GetAwsAccountNameByProfile(awsProfile)
	GetDetailsFromExcel(fp, resourceIdName, "AwsAccountName", "AwsRegion", sheetName, acctName, awsRegion, &ress)

	for i := range ress.Resources {
		ids = append(ids, aws.String(ress.Resources[i].ResourceID))
	}

	//input := &ec2.DescribeInstancesInput{
	//	InstanceIds: instids,
	//}

	input := &ec2.CreateTagsInput{
		Resources: ids,
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Deprecation"),
				Value: aws.String("PreDeprecationCleanUp"),
			},
		},
	}

	result, err := ec2Svc.CreateTags(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	println(result)

}

//func UpdateEbsTags(ses client.ConfigProvider, awsProfile string, awsRegion string, fp string) {
//	ec2Svc := ec2.New(ses)
//	ebsids := []*string{}
//	ress := Resources{}
//	acctName := GetAwsAccountNameByProfile(awsProfile)
//	GetDetailsFromExcel(fp, "VolumeId", "AwsAccountName","AwsRegion", "EBS", acctName,awsRegion, &ress)
//
//	for i := range ress.Resources {
//		ebsids = append(ebsids, aws.String(ress.Resources[i].ResourceID))
//	}
//
//	input := &ec2.CreateTagsInput{
//		Resources: ebsids,
//		Tags: []*ec2.Tag{
//			{
//				Key:   aws.String("Deprecation"),
//				Value: aws.String("PreDeprecation"),
//			},
//		},
//	}
//
//	result, err := ec2Svc.CreateTags(input)
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//			switch aerr.Code() {
//			default:
//				fmt.Println(aerr.Error())
//			}
//		} else {
//			// Print the error, cast err to awserr.Error to get the Code and
//			// Message from an error.
//			fmt.Println(err.Error())
//		}
//		return
//	}
//
//	println(result)
//}
//
//func UpdateSnapShotsTags(ses client.ConfigProvider, awsProfile string, awsRegion string, fp string) {
//	ec2Svc := ec2.New(ses)
//	ebsids := []*string{}
//	ress := Resources{}
//	acctName := GetAwsAccountNameByProfile(awsProfile)
//	GetDetailsFromExcel(fp, "id", "AwsAccountName","AwsRegion", "SnapShots", acctName,awsRegion, &ress)
//
//	for i := range ress.Resources {
//		ebsids = append(ebsids, aws.String(ress.Resources[i].ResourceID))
//	}
//
//	input := &ec2.CreateTagsInput{
//		Resources: ebsids,
//		Tags: []*ec2.Tag{
//			{
//				Key:   aws.String("Deprecation"),
//				Value: aws.String("PreDeprecation"),
//			},
//		},
//	}
//
//	result, err := ec2Svc.CreateTags(input)
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//			switch aerr.Code() {
//			default:
//				fmt.Println(aerr.Error())
//			}
//		} else {
//			// Print the error, cast err to awserr.Error to get the Code and
//			// Message from an error.
//			fmt.Println(err.Error())
//		}
//		return
//	}
//
//	println(result)
//}
