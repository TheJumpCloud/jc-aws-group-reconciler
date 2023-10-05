package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

const contentType string = "application/json"
const accept string = "application/json"
const maxResults int32 = 100

// getBoundJumpCloudGroups returns a list of JCUserGroups that are bound to the
// supplied (aws) application ids. Each JCUserGroup contains the users in the group
func getBoundJumpCloudGroups() UserGroupCollection {
	fmt.Println("Collecting JumpCloud groups that are bound to AWS (this may take some time)...")
	apiKey := os.Getenv("JUMPCLOUD_API_KEY")
	applicationIDs := strings.Split(os.Getenv("JUMPCLOUD_APPLICATION_IDS"), ",")

	clientV2 := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())
	ctxV2 := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: apiKey,
	})

	clientV1 := jcapiv1.NewAPIClient(jcapiv1.NewConfiguration())
	ctxV1 := context.WithValue(context.TODO(), jcapiv1.ContextAPIKey, jcapiv1.APIKey{
		Key: apiKey,
	})

	boundGroups := []UserGroup{}
	for _, appID := range applicationIDs {
		fmt.Print(">")
		allBoundUserGroups := []jcapiv2.GraphObjectWithPaths{}
		var count int32 = 0
		optionalParams := map[string]interface{}{"skip": count}
		for {
			boundUserGroups, _, err := clientV2.ApplicationsApi.GraphApplicationTraverseUserGroup(
				ctxV2, appID, contentType, accept, optionalParams)

			if err != nil {
				panic(fmt.Sprintf(
					"could not get user groups bound to application id: %s, count: %d, error: %v",
					appID, count, err))
			}
			allBoundUserGroups = append(allBoundUserGroups, boundUserGroups...)

			if len(boundUserGroups) < int(maxResults) {
				break
			} else {
				count++
				optionalParams = map[string]interface{}{"skip": count * maxResults}
			}
		}

		for _, boundUserGroup := range allBoundUserGroups {
			fmt.Print(">")
			group, _, err := clientV2.UserGroupsApi.GroupsUserGet(
				ctxV2, boundUserGroup.Id, contentType, accept, nil)
			if err != nil {
				panic(fmt.Sprintf(
					"could not get user group name for id: %s: %v", boundUserGroup.Id, err))
			}

			members, _, err := clientV2.UserGroupsApi.GraphUserGroupMembersList(
				ctxV2, boundUserGroup.Id, contentType, accept, nil)
			if err != nil {
				panic(fmt.Sprintf(
					"could not list (JC) group members for: %s, %v", group.Name, err))
			}
			// fmt.Println("members for", group.Name, len(members))
			groupMembers := []User{}
			for _, member := range members {
				user, _, err := clientV1.SystemusersApi.SystemusersGet(
					ctxV1, member.To.Id, contentType, accept, nil)
				if err != nil {
					panic(fmt.Sprintf("could not get user: %v", err))
				}
				groupMembers = append(groupMembers, User{Email: user.Email})
			}
			boundGroups = append(boundGroups, UserGroup{Name: group.Name, Users: groupMembers})
		}
	}
	fmt.Println("\nDONE: total # JumpCloud groups bound to AWS:", len(boundGroups))
	return UserGroupCollection{Groups: boundGroups}
}
