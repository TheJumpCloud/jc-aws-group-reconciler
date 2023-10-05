package main

type User struct {
	Email string
}

type UserGroup struct {
	Name  string
	Users []User
}

func (g UserGroup) findByName(name string) bool {
	return g.Name == name
}

type UserGroupCollection struct {
	Groups []UserGroup
}

func (gc UserGroupCollection) findByName(name string) *UserGroup {
	for _, group := range gc.Groups {
		if group.Name == name {
			return &group
		}
	}
	return nil
}

// compareUserGroupMembers compares the members in group a to
// the members in group b. Any any members that are in a but
// not b are returned in a list.
func compareUserGroupMembers(a, b *UserGroup) []string {
	difference := []string{}
	for _, memberA := range a.Users {
		isFound := false
		for _, memberB := range b.Users {
			if memberA == memberB {
				isFound = true
				break
			}
		}
		if !isFound {
			difference = append(difference, memberA.Email)
		}
	}
	return difference
}
