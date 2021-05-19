package resutil

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sts"
	. "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	"os"
	"sort"
	"strings"
)

func GetUserInput(message string, lst []string, toUpperValidation bool) string {
	log.Debug("in util=>GetUserInput")
	reader := bufio.NewReader(os.Stdin)
	isValidInput := false
	returnedValue := ""

	for !isValidInput {
		fmt.Print(message, Bold(Green(lst)), ":")
		log.Info(message, Bold(Green(lst)), ":")
		text, _ := reader.ReadString('\n')
		fmt.Println("") //just print extra space on screen so the whole flow will be more readable
		if toUpperValidation {
			text = strings.ToUpper(text)
		}
		//returnedValue = strings.Replace(text, "\r\n", "", -1) // I need this line becuse isValidInput method reset the text value
		returnedValue = strings.Replace(text, "\n", "", -1) // I need this line becuse isValidInput method reset the text value
		isValidInput = validateInput(text, lst)
	}

	log.Info(returnedValue)

	return returnedValue
}

func GetUserInputNoValidation(message string) string {
	log.Debug("in util=>GetUserInput")
	reader := bufio.NewReader(os.Stdin)
	returnedValue := ""

	fmt.Print(message, ":")
	log.Info(message, ":")
	text, _ := reader.ReadString('\n')
	t := strings.Replace(text, "\n", "", -1)
	returnedValue = t
	log.Info(returnedValue)

	return returnedValue
}

func validateInput(input string, lst []string) bool {
	log.Debug("in util=>validateInput")
	//fmt.Printf("ENvironment=%s",os.Environ())

	//t := strings.Replace(input,"\n","",-1)
	//t := strings.Replace(input, "\r\n", "", -1)
	t := strings.Replace(input, "\n", "", -1)

	//fmt.Println("my input after conversion is "+t)

	sort.Strings(lst)
	i := sort.SearchStrings(lst, t)
	if !(i < len(lst) && lst[i] == t) {
		fmt.Println(White(t).BgRed(), White(" is not valid input... please try again").BgRed())
	}

	return i < len(lst) && lst[i] == t
}

func GetAwsAccountNameByProfile(profile string) string {
	accountName := ""
	aws_account_profile_to_name := map[string]string{
		"lpap":    "LPA_PROD",
		"lpas":    "LPA_STG",
		"lpa-dev": "LPA_DEV",
		"lpacdn":  "LPA_CDN",
	}

	for key, val := range aws_account_profile_to_name {

		if profile != key {
			delete(aws_account_profile_to_name, key)
		} else {
			accountName = val
		}

	}

	return accountName
}

func getAwsAccountInfo(ses client.ConfigProvider) []string {
	svc := sts.New(ses)
	input := &sts.GetCallerIdentityInput{}
	acctInfo := []string{}
	aws_account_to_name := map[string]string{
		"809699088482": "LPA_PROD",
		"218020736831": "LPA_STG",
		"178600608703": "LPA_DEV",
		"444671552647": "LPA_CDN",
	}

	result, err := svc.GetCallerIdentity(input)
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

	}
	println(result)
	for key, val := range aws_account_to_name {
		if *result.Account != key {
			//awsAccountName = value
			delete(aws_account_to_name, key)
		} else {
			acctInfo = append(acctInfo, key, val)
		}

	}

	return acctInfo

}
