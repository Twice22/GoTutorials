package main 

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type sortbyNbClicks struct {
	tracks []*Track
	less func(t1, t2 *Track) bool
}

func (x sortbyNbClicks) Len() int { return len(x.tracks) }
func (x sortbyNbClicks) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x sortbyNbClicks) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

// create slice containing pointer to track struct elements
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}


var table = template.Must(template.New("table").Parse(`
	<html>
		<head>
			<title>Sorted Table</title>
		</head>

		<body>
			<table>
				<thead>
					<tr>
						<th><a href="/?sort=Title">Title</a></th>
						<th><a href="/?sort=Artist">Artist</a></th>
						<th><a href="/?sort=Album">Album</a></th>
						<th><a href="/?sort=Year">Year</a></th>
						<th><a href="/?sort=Length">Length</a></th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{ .Title}}</td>
						<td>{{ .Artist}}</td>
						<td>{{ .Album}}</td>
						<td>{{ .Year}}</td>
						<td>{{ .Length}}</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</body>
	</html>
	`))

func printTracks(writer io.Writer, tracks []*Track) {
	if err := table.Execute(writer, tracks); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	sortBy := request.FormValue("sort")

	sort.Sort(sortbyNbClicks{tracks, func(t1, t2 *Track) bool {
		switch sortBy {
		case "Title":
			return t1.Title < t2.Title
		case "Year":
			return t1.Year < t2.Year
		case "Length":
			return t1.Length < t2.Length
		case "Artist":
			return t1.Artist < t2.Artist
		case "Album":
			return t1.Album < t2.Album
		}
		return false
	}})
	printTracks(responseWriter, tracks)
}