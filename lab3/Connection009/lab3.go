package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	//舊版go
	//get the package by command "$go get github.com/adonovan/gopl.io/ch4/github"
	//github "github.com/adonovan/gopl.io/ch4/github"

	//go 版本17 get the package by command
	//cd Lab4
	//go mod init whatever_you_like (在Lab4資料夾下開mod)
	//go mod tidy   (或 go get gopl.io/ch4/github)
	"gopl.io/ch4/github"
)

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>#{{.Number}} {{.Title}}</h1>
<dl>
	<dt>User:</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>State:</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type newIssues struct {
	github.IssuesSearchResult
}

// Call this function to print error logs
func logPrint(v interface{}) {
	if v != nil {
		log.Print(v)
	}
}

// Hint: use "logPrint(issueTemplate.Execute(w, ???))" to render html
func (nis newIssues) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	if len(pathParts) < 3 || pathParts[2] == "" {
		logPrint(issueListTemplate.Execute(w, nis))

		return
	}

	logPrint(issueTemplate.Execute(w, nis))
}

func main() {
	queryString := []string{"repo:vuejs/vue", "is:open", "label:bug"}
	isr, err := github.SearchIssues(queryString)

	if err != nil {
		log.Fatal(err)
	}

	ni := newIssues{ *isr }
	http.Handle("/", http.HandlerFunc(ni.ServeHTTP))
	logPrint(http.ListenAndServe(":8080", nil))
}
