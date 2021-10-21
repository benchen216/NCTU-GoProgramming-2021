package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
	"gopl.io/ch4/github"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>Total {{.Items | len}} issues</h1>
<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range $i, $e := .Items}}
	<tr>
		<td><a href='/issues/{{$i}}'>{{.Number}}</td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='/issues/{{$i}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
</table>
`))

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

func logPrint(v interface{}) {
	if v != nil {
		log.Print(v)
	}
}

func (nis newIssues) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	if len(pathParts) < 3 || pathParts[2] == "" {
		logPrint(issueListTemplate.Execute(w, nis.IssuesSearchResult))
		return
	}
	id, _ := strconv.Atoi(pathParts[2])
	logPrint(issueTemplate.Execute(w, *nis.Items[id]))
}

func main() {
	queryString := []string{"repo:vuejs/vue", "is:open", "label:bug"}
	isr, err := github.SearchIssues(queryString)

	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", newIssues{*isr})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
