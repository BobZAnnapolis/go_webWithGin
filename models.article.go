// models.article.go

package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// for this demo, we're storing the article list inmemory
// in a real app, db or a list in static files
var articleList = []article{
	article{ID: 13, Title: "Article 13 Title", Content: "13 Yule Lads in Iceland"},
	article{ID: 666, Title: "Article 666 Title", Content: "Beelzebub Has A Devil Put Aside For Me"},
}

// return a list of all the articles
func getAllArticles() []article {
	return articleList
}
