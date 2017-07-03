package main 

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	// func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}


type sortbyNbClicks struct {
	tracks []*Track
	less func(t1, t2 *Track) bool
}

func (x sortbyNbClicks) Len() int { return len(x.tracks) }
func (x sortbyNbClicks) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x sortbyNbClicks) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }



func main() {
	// first string in recentClicks are the most clicked string
	recentClicks := []string{"Title", "Year", "Artist"}

	sort.Sort(sortbyNbClicks{tracks, func(t1, t2 *Track) bool {
		if len(recentClicks) == 0 {
			return t1.Title < t2.Title
		}

		for _, click := range recentClicks {
			switch click {
			case "Title":
				if t1.Title != t2.Title {
					return t1.Title < t2.Title
				}
			case "Year":
				if t1.Year != t2.Year {
					return t1.Year < t2.Year
				}
			case "Length":
				if t1.Length != t2.Length {
					return t1.Length < t2.Length
				}
			case "Artist":
				if t1.Artist != t2.Artist {
					return t1.Artist < t2.Artist
				}
			}
		}

		return false
	}})

	printTracks(tracks)
}