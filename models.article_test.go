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

// Test the function that fetche an Article by its ID
func TestGetArticleByID(t *testing.T) {
	a, err := getArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}

// Test the function that creates a new article
func TestCreateNewArticle(t *testing.T) {
	// get the original count of articles
	originalLength := len(getAllArticles())

	// add another article
	a, err := createNewArticle("New test title", "New test content")

	// get the new count of articles
	allArticles := getAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
