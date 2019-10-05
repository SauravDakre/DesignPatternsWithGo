package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	userDB := make(userRecords, 0)
	for i := 1; i <= 100; i++ {
		userDB = append(userDB, user{id: i})
	}

	expectedCache := []user{}
	proxy := newUserRecordsProxy(&userDB, 2)
	_, err := proxy.find(1)
	if err != nil {
		t.Error("expecting err to be nil but got", err)
	}
	expectedCache = append(expectedCache, user{1})
	equal := assertUserSlice(proxy.cache, expectedCache)
	if !equal {
		t.Error("expecting cache contents to be same but got different")
	}

	_, err = proxy.find(2)
	if err != nil {
		t.Error("expecting err to be nil but got", err)
	}
	expectedCache = append(expectedCache, user{2})
	equal = assertUserSlice(proxy.cache, expectedCache)
	if !equal {
		t.Error("expecting cache contents to be same but got different")
	}

	_, err = proxy.find(3)
	if err != nil {
		t.Error("expecting err to be nil but got", err)
	}
	expectedCache = expectedCache[1:]
	expectedCache = append(expectedCache, user{3})
	equal = assertUserSlice(proxy.cache, expectedCache)
	if !equal {
		t.Error("expecting cache contents to be same but got different")
	}

	_, err = proxy.find(2)
	if err != nil {
		t.Error("expecting err to be nil but got", err)
	}
	equal = assertUserSlice(proxy.cache, expectedCache)
	if !equal {
		t.Error("expecting cache contents to be same but got different")
	}
}

func assertUserSlice(s1, s2 []user) bool {
	b := true
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return b
}
