// models.article_test.go

package main

import "testing"

// test the function that fetches all the articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check that the length of the list of articles returned is the
	// same as the length of the global variable holdingthe list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check that each member is identical
	for i, v := range alist {
		if v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title ||
			v.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}
}
