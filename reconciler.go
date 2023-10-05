package main

import (
	"fmt"
	"os"
	"sort"
)

func reconcileAWSGroupsToJCGroups(jcGroups, awsGroups UserGroupCollection) {
	outstandingAWSGroups := []string{}
	for _, awsGroup := range awsGroups.Groups {
		isOutstanding := true
		for _, jcGroup := range jcGroups.Groups {
			if awsGroup.Name == jcGroup.Name {
				isOutstanding = false
			}
		}
		if isOutstanding {
			outstandingAWSGroups = append(outstandingAWSGroups, awsGroup.Name)
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

func reconcileJCGroupMembersToAWSGroupMembers(jcGroups, awsGroups UserGroupCollection) {
	// 1. Loop over all JC groups
	// 2. For each group find the associated AWS Group
	// 3. Ensure that the members in both groups are the same
	// 4. If not, print aws group name and members that exist
	//    in AWS but not in JC

	outOfSyncExists := false
	for _, jcGroup := range jcGroups.Groups {
		awsGroup := awsGroups.findByName(jcGroup.Name)
		if awsGroup == nil {
			panic(fmt.Sprintf("could not find matching AWS group named: %s", jcGroup.Name))
		}
		outstandingUsers := compareUserGroupMembers(awsGroup, &jcGroup)
		if len(outstandingUsers) == 0 {
			fmt.Printf("In Sync: %s\n", awsGroup.Name)
			break
		}
		// out of sync
		outOfSyncExists = true
		fmt.Printf("OUT OF SYNC: %s\n", awsGroup.Name)
		for _, u := range outstandingUsers {
			fmt.Println("\t-", u)
		}
	}

	if outOfSyncExists {
		fmt.Println("WARNING: There are groups in AWS that are out of sync with JumpCloud!")
		fmt.Println("Manually reconcile the list of groups and user memberships above.")
		fmt.Println("The users listed above should be remove from the specified AWS group.")
	}
}
