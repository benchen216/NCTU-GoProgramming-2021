package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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
		/*
			List issues (issueListTemplate) here
		*/

		return
	}

	/*
		Show issues (issueTemplate) here
	*/

}

func main() {
	queryString := []string{"repo:vuejs/vue", "is:open", "label:bug"}
	isr, err := github.SearchIssues(queryString)

	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", newIssues.ServeHTTP())

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	/*
		Hint: "isr" is "github.issuesSearchResult"
			http.Handle("/", ???)
			log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
