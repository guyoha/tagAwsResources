package resutil

type resource struct {
	Resourcetype   string
	ResourceID     string
	AwsAccountID   string
	AwsAccountName string
	AwsRegion      string
}

type Resources struct {
	Resources []resource
}
