package services

import (
	"testing"
)

func Test_Authorization_CreateToken(t *testing.T) {

	td, err := CreateToken("somekey")

	if err != nil {
		t.Error("Error when creatingg token details")
	}

	if len(td.AccessToken) == 0 {
		t.Error("Invalid token length")
	}

	if len(td.RefreshToken) == 0 {
		t.Error("Must have refresh token")
	}

}
