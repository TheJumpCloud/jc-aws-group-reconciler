package main

import (
	"fmt"
	"os"
)

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

func main() {
	if missing := checkEnvironmentVariables(); missing != "" {
		printUsage(missing)
		return
	}

	jumpCloudGroupsBoundToAWS := getBoundJumpCloudGroups()
	awsGroups := getAWSGroups()
	reconcileAWSGroupsToJCGroups(jumpCloudGroupsBoundToAWS, awsGroups)
	reconcileJCGroupMembersToAWSGroupMembers(jumpCloudGroupsBoundToAWS, awsGroups)
	fmt.Print("\n\nDONE\n")
}
