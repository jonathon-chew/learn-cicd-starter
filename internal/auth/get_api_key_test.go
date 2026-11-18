package auth

import (
	"net/http"
	"testing"
)

func Test_api_key_non_existant(t *testing.T) {
	f := make(http.Header, 10)

	_, err := GetAPIKey(f)

	if err != nil {
		t.Logf("Successfully failed with: %v", err)
	} else {
		t.Fail()
	}
}

func Test_api_key_blank(t *testing.T) {
	f := make(http.Header, 10)
	f.Add("Authorization", "")

	_, err := GetAPIKey(f)

	if err != nil {
		t.Logf("Successfully failed with: %v", err)
	} else {
		t.Fail()
	}
}

func Test_api_key_nonesense(t *testing.T) {
	f := make(http.Header, 10)
	f.Add("Authorization", "ofijosifodskfkdns")

	_, err := GetAPIKey(f)

	if err != nil {
		t.Logf("Successfully failed with: %v", err)
	} else {
		t.Fail()
	}
}
