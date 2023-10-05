package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	awscreds "github.com/aws/aws-sdk-go/aws/credentials"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	awsidstore "github.com/aws/aws-sdk-go/service/identitystore"
)

// getAWSGroups returns a list of all of the user groups in the specified AWS region
func getAWSGroups() UserGroupCollection {
	fmt.Println("Collecting AWS groups and members (this may take some time)...")
	session := awssession.Must(awssession.NewSession())
	creds := awscreds.NewCredentials(&awscreds.EnvProvider{})
	idStoreSvc := awsidstore.New(session, &aws.Config{Credentials: creds})

	idStoreID := os.Getenv("AWS_ID_STORE_ID")
	allAWSGroups := []*awsidstore.Group{}
	var nextToken *string = nil
	var maxResults int64 = 100
	for {
		fmt.Print(">")
		input := &awsidstore.ListGroupsInput{
			IdentityStoreId: &idStoreID,
			MaxResults:      &maxResults,
		}
		if nextToken != nil {
			input.SetNextToken(*nextToken)
		}
		groups, err := idStoreSvc.ListGroups(input)

		if err != nil {
			panic(fmt.Sprintf("failed to get aws groups: %v", err))
		}
		if groups.Groups == nil {
			panic("no groups returned from AWS: groups is nil")
		}

		allAWSGroups = append(allAWSGroups, groups.Groups...)
		nextToken = groups.NextToken

		if len(groups.Groups) < 100 {
			break
		}
	}

	groupNames := []UserGroup{}
	for _, group := range allAWSGroups {
		fmt.Print(">")
		members, err := idStoreSvc.ListGroupMemberships(
			&awsidstore.ListGroupMembershipsInput{
				GroupId:         group.GroupId,
				IdentityStoreId: &idStoreID,
			},
		)
		if err != nil {
			panic(fmt.Sprintf(
				"could not list (AWS) group members for: %s: %v", *group.DisplayName, err))
		}
		users := []User{}
		for _, member := range members.GroupMemberships {
			// fmt.Println(member)
			user, err := idStoreSvc.DescribeUser(&awsidstore.DescribeUserInput{
				IdentityStoreId: &idStoreID,
				UserId:          member.MemberId.UserId,
			})
			if err != nil {
				panic(fmt.Sprintf(
					"could not describe (AWS) user: %s: %v", *member.MemberId.UserId, err))
			}
			// fmt.Println(*user.UserName)
			users = append(users, User{Email: *user.UserName})
		}
		groupNames = append(groupNames,
			UserGroup{Name: *group.DisplayName,
				Users: users,
			})
	}
	fmt.Println("\nDONE: total # AWS groups:", len(allAWSGroups))
	return UserGroupCollection{Groups: groupNames}
}
