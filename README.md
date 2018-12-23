# go_webWithGin
golang web &amp; microservices project using gin framework, from :

https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin

see notes file for addl details

Pre-REQUISITES:
- properly set up Go environment

Only uses 1 dependency :
- go get -u github.com/gin-gonic/gin

- after you git clone and install above dependency, from command line
  - go build -o main and then access from localhost:8080
  - or, "go run ..." and list ALL the go files

- invocation:
 - /a/view/xx where 'xx' is the number of an article displayed
[GIN-debug] GET    /                         --> main.showIndexPage (4 handlers)
[GIN-debug] GET    /u/login                  --> main.showLoginPage (5 handlers)
[GIN-debug] POST   /u/login                  --> main.performLogin (5 handlers)
[GIN-debug] GET    /u/logout                 --> main.logout (5 handlers)
[GIN-debug] GET    /u/register               --> main.showRegistrationPage (5 handlers)
[GIN-debug] POST   /u/register               --> main.register (5 handlers)
[GIN-debug] GET    /a/view/:article_id       --> main.getArticle (4 handlers)
[GIN-debug] GET    /a/create                 --> main.showArticleCreationPage (5 handlers)
[GIN-debug] POST   /a/create                 --> main.createArticle (5 handlers)

