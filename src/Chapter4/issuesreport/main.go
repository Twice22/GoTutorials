package main 

import (
	"log"
	"os"
	"text/template"
	"time"

	"Chapter4/github"
)

// 'text/template' allows to substitute the values of variables into a text
// {{...}} are actions
// the current value is referred to as "dot". The "dot" initially refers to the template's
// parameter. Here it's a github.IssuesSearchResult

// {{range .Items}} and {{end}} create a loop
// the | notation within an action makes the result of one operation the argument of another.
// printf is a built-in synonym for fmt.Sprintf
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Parse the template named templ defined above
// Funcs add daysAgo to the set of fcts accessible within the template
/*report, err := template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ)
if err != nil {
	log.Fatal(err)
}*/

// we use template.Must (see Chapter 5.9) helper fct that makes error
// handling more convenient
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

// use go build Chapter4/issuesreport
// ./issuesreport.exe repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}