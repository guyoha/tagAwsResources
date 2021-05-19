package resutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"strings"
)

func UpdateElbTags(ses client.ConfigProvider, awsProfile string, awsRegion string, fp string, sheetName string, resourceIdName string) {

	svc := elbv2.New(ses)
	//elbArns := []*string{}
	ress := Resources{}
	acctName := GetAwsAccountNameByProfile(awsProfile)
	GetDetailsFromExcel(fp, resourceIdName, "AwsAccountName", "AwsRegion", sheetName, acctName, awsRegion, &ress)

	for i := range ress.Resources {
		elbs := strings.Split(ress.Resources[i].ResourceID, "|")

		for x := range elbs {
			if elbs[x] != "" {
				//elbArns = append(elbArns, aws.String(elbs[x]))

				input := &elbv2.AddTagsInput{
					ResourceArns: []*string{
						aws.String(elbs[x]),
					},
					Tags: []*elbv2.Tag{
						{
							Key:   aws.String("Deprecation"),
							Value: aws.String("PreDeprecationCleanUp"),
						},
					},
				}

				result, err := svc.AddTags(input)
				if err != nil {
					if aerr, ok := err.(awserr.Error); ok {
						switch aerr.Code() {
						case elb.ErrCodeAccessPointNotFoundException:
							fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
						case elb.ErrCodeTooManyTagsException:
							fmt.Println(elb.ErrCodeTooManyTagsException, aerr.Error())
						case elb.ErrCodeDuplicateTagKeysException:
							fmt.Println(elb.ErrCodeDuplicateTagKeysException, aerr.Error())
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

				fmt.Println(result)

			}

		}
	}
}

//func UpdateTGsTags(ses client.ConfigProvider, awsProfile string, awsRegion string, fp string) {
//
//	svc := elbv2.New(ses)
//	//elbArns := []*string{}
//	ress := Resources{}
//	acctName := GetAwsAccountNameByProfile(awsProfile)
//	GetDetailsFromExcel(fp, "TargetArn", "AwsAccountName","AwsRegion", "ELB", acctName,awsRegion, &ress)
//
//	for i := range ress.Resources {
//		tgs := strings.Split(ress.Resources[i].ResourceID,"|")
//
//		for x := range tgs{
//			if tgs[x] != ""{
//				//elbArns = append(elbArns, aws.String(elbs[x]))
//
//				input := &elbv2.AddTagsInput{
//					ResourceArns: []*string{
//						aws.String(tgs[x]),
//					},
//					Tags: []*elbv2.Tag{
//						{
//							Key:   aws.String("Deprecation"),
//							Value: aws.String("PreDeprecation"),
//						},
//					},
//				}
//
//				result, err := svc.AddTags(input)
//				if err != nil {
//					if aerr, ok := err.(awserr.Error); ok {
//						switch aerr.Code() {
//						case elb.ErrCodeAccessPointNotFoundException:
//							fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
//						case elb.ErrCodeTooManyTagsException:
//							fmt.Println(elb.ErrCodeTooManyTagsException, aerr.Error())
//						case elb.ErrCodeDuplicateTagKeysException:
//							fmt.Println(elb.ErrCodeDuplicateTagKeysException, aerr.Error())
//						default:
//							fmt.Println(aerr.Error())
//						}
//					} else {
//						// Print the error, cast err to awserr.Error to get the Code and
//						// Message from an error.
//						fmt.Println(err.Error())
//					}
//					return
//				}
//
//				fmt.Println(result)
//			}
//
//		}
//	}
//}
