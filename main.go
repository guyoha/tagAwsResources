package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"tagAwsResources/main/resutil"
)

func main() {

	//awsProfiles := []string{"lpap", "lpas", "lpa-dev", "lpcdn"}
	//dcRegions := []string{"us-east-1", "us-east-2", "us-west-2", "ap-southeast-2", "eu-west-1"}
	//actions := []string{"TagEC2", "TagEBS", "TagELB", "TagTg", "TagSnapShots"}
	validateparams := []string{"Yes", "No"}

	//profile := resutil.GetUserInput("Aws Profile: ", awsProfiles, false)
	//region := resutil.GetUserInput("Aws Region: ", dcRegions, false)
	//action := resutil.GetUserInput("What activity: ", actions, false)
	//filePath := resutil.GetUserInputNoValidation("Enter Path to Excel")

	profile := os.Args[1]
	region := os.Args[2]
	action := os.Args[3]
	filePath := os.Args[4]

	println("profile:" + os.Args[1])
	println("region:" + os.Args[2])
	println("action:" + os.Args[3])
	println("filePath:" + os.Args[4])

	validateparam := resutil.GetUserInput("Confirm Values: ", validateparams, false)
	println(validateparam)

	println(profile, region)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config:  aws.Config{Region: aws.String(region)},
	}))

	switch action {

	case "TagEC2":
		resutil.UpdateEc2Tags(sess, profile, region, filePath, "EC2", "InstanceId")
	case "TagEBS":
		resutil.UpdateEc2Tags(sess, profile, region, filePath, "EBS", "VolumeId")
	case "TagSnapShots":
		resutil.UpdateEc2Tags(sess, profile, region, filePath, "SnapShots", "id")
	case "TagELB":
		resutil.UpdateElbTags(sess, profile, region, filePath, "ELB", "LbsArn")
	case "TagTg":
		resutil.UpdateElbTags(sess, profile, region, filePath, "ELB", "TargetArn")
	}
}
