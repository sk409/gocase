package gocase_test

import (
	"testing"

	"github.com/sk409/gocase"
)

func TestUpperCamelCase(t *testing.T) {
	faild := func(e, g string) {
		t.Errorf("Failed to UpperCamelCase: Expected %s but got %s", e, g)
	}
	words := []string{"id", "userId", "userIdle", "commitSha1"}
	expectation := []string{"Id", "UserId", "UserIdle", "CommitSha1"}
	expectationGoLike := []string{"ID", "UserID", "UserIdle", "CommitSHA1"}
	for index, word := range words {
		b := []byte(word)
		u := gocase.UpperCamelCase(b, false)
		s := string(u)
		e := expectation[index]
		if s != e {
			faild(e, s)
			return
		}
		u = gocase.UpperCamelCase(b, true)
		s = string(u)
		e = expectationGoLike[index]
		if s != e {
			faild(e, s)
			return
		}
	}
}

func TestLowerCamelCase(t *testing.T) {
	faild := func(e, g string) {
		t.Errorf("Failed to LowerCamelCase: Expected %s but got %s", e, g)
	}
	words := []string{"ID", "UserID", "userIdle", "commitSHA1"}
	expectation := []string{"id", "userId", "userIdle", "commitSha1"}
	expectationGoLike := []string{"id", "userID", "userIdle", "commitSHA1"}
	for index, word := range words {
		b := []byte(word)
		u := gocase.LowerCamelCase(b, false)
		s := string(u)
		e := expectation[index]
		if s != e {
			faild(e, s)
			return
		}
		u = gocase.LowerCamelCase(b, true)
		s = string(u)
		e = expectationGoLike[index]
		if s != e {
			faild(e, s)
			return
		}
	}
}
