package rouser_test

import "testing"

func TestFetchAuthenticatedUser(t *testing.T) {
	authenticatedUser, err := ruser.FetchAuthenticatedUser()
	if err != nil {
		t.Errorf("error fetching authenticated user: %s", err)
	}
	if authenticatedUser.ID == 0 {
		t.Error("authenticated user had an id of 0")
	}
}
