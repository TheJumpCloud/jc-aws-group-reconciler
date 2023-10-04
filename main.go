package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"

	"github.com/aws/aws-sdk-go/aws"
	awscreds "github.com/aws/aws-sdk-go/aws/credentials"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	awsidstore "github.com/aws/aws-sdk-go/service/identitystore"
)

const contentType string = "application/json"
const accept string = "application/json"

var requiredEnvVars []string = []string{
	"JUMPCLOUD_API_KEY",
	"JUMPCLOUD_APPLICATION_IDS",
	"AWS_REGION",
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
	"AWS_ID_STORE_ID",
}

var optionalEnvVars []string = []string{
	"AWS_SESSION_TOKEN",
}

// checkEnvironmentVariables ensures that all the required environment
// variables are set and have a value
func checkEnvironmentVariables() string {
	for _, name := range requiredEnvVars {
		if value, isSet := os.LookupEnv(name); !isSet || value == "" {
			return name
		}
	}
	return ""
}

func printUsage(missing string) {
	fmt.Println("The following environment variables are required to be set:")
	for _, name := range requiredEnvVars {
		fmt.Println(name)
	}
	fmt.Println("\nIf your aws connection uses SSO then AWS_SESSION_TOKEN must be set as well")

	if missing != "" {
		fmt.Println("environment variable is not set:", missing)
	}
}

// getBoundJumpCloudGroups returns a list of all of the user groups bound to the
// supplied (aws) application ids
func getBoundJumpCloudGroups() []string {
	apiKey := os.Getenv("JUMPCLOUD_API_KEY")
	applicationIDs := strings.Split(os.Getenv("JUMPCLOUD_APPLICATION_IDS"), ",")

	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())
	ctx := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: apiKey,
	})

	boundGroups := []string{}
	for _, appID := range applicationIDs {
		allBoundUserGroups := []jcapiv2.GraphObjectWithPaths{}
		var count int32 = 0
		optionalParams := map[string]interface{}{"skip": count}
		for {
			boundUserGroups, _, err := client.ApplicationsApi.GraphApplicationTraverseUserGroup(
				ctx, appID, contentType, accept, optionalParams)

			if err != nil {
				panic(fmt.Sprintf(
					"could not get user groups bound to application id: %s, count: %d, error: %v",
					appID, count, err))
			}
			allBoundUserGroups = append(allBoundUserGroups, boundUserGroups...)

			if len(boundUserGroups) < 100 {
				break
			} else {
				count++
				optionalParams = map[string]interface{}{"skip": count * 100}
			}
		}

		for _, boundUserGroup := range allBoundUserGroups {
			group, _, err := client.UserGroupsApi.GroupsUserGet(
				ctx, boundUserGroup.Id, contentType, accept, nil)
			if err != nil {
				panic(fmt.Sprintf(
					"could not get user group name for id: %s: %v", boundUserGroup.Id, err))
			}
			boundGroups = append(boundGroups, group.Name)
		}
	}

	return boundGroups
}

// getAWSGroups returns a list of all of the user groups in the specified AWS region
func getAWSGroups() []string {
	session := awssession.Must(awssession.NewSession())
	creds := awscreds.NewCredentials(&awscreds.EnvProvider{})
	idStoreSvc := awsidstore.New(session, &aws.Config{Credentials: creds})

	idStoreID := os.Getenv("AWS_ID_STORE_ID")
	groups, err := idStoreSvc.ListGroups(&awsidstore.ListGroupsInput{IdentityStoreId: &idStoreID})
	// fmt.Println(awsGroups, err)

	if err != nil {
		panic(fmt.Sprintf("failed to get aws groups: %v", err))
	}

	if groups.Groups == nil {
		panic("no groups returned from AWS: groups is nil")
	}

	// HACK/TODO: get pagination working for the AWS ListGroups call. Initial attempts proved
	// unsuccessful.
	if len(groups.Groups) > 99 {
		fmt.Println("***WARNING*** More than 99 AWS groups found, results may be incomplete")
		fmt.Println("Re-run this command after cleaning up AWS groups")
	}

	groupNames := []string{}
	for _, group := range groups.Groups {
		groupNames = append(groupNames, *group.DisplayName)
	}

	return groupNames
}

func reconcileAWStoJC(jcGroups, awsGroups []string) {
	outstandingAWSGroups := []string{}
	for _, awsGroup := range awsGroups {
		// fmt.Println(*group.DisplayName)
		isOutstanding := true
		for _, jcGroup := range jcGroups {
			if awsGroup == jcGroup {
				isOutstanding = false
			}
		}
		if isOutstanding {
			outstandingAWSGroups = append(outstandingAWSGroups, awsGroup)
		}
	}

	sort.Strings(outstandingAWSGroups)

	fmt.Printf("There are %v groups in AWS that are NOT bound to ", len(outstandingAWSGroups))
	fmt.Printf("the JumpCloud application ids:\n %v\n", os.Getenv("JUMPCLOUD_APPLICATION_IDS"))
	fmt.Println("The following groups should be removed from AWS:")
	for _, group := range outstandingAWSGroups {
		fmt.Println(group)
	}
}

func main() {
	if missing := checkEnvironmentVariables(); missing != "" {
		printUsage(missing)
		return
	}

	jumpCloudGroupsBoundToAWS := getBoundJumpCloudGroups()
	awsGroups := getAWSGroups()
	reconcileAWStoJC(jumpCloudGroupsBoundToAWS, awsGroups)
}
