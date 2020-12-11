// comparestruct -
package main

import "fmt"

// song struct
type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {

	rock := playlist{genre: "indie rock",
		songs: []song{
			{title: "winderwall", artist: "oasis"},
			{title: "super sonic", artist: "oasis"},
		},
	}

	// make a copy of the first song in rock playlist songs
	// song := rock.songs[0]

	// Chang the song title | this does not change the original songs slice
	// song.title = "Testing Song"
	// fmt.Printf("%+v\n%+v\n\n", song, rock.songs[0])

	// to change the original songs 1. title
	// rock.songs[0].title = "Live forever"

	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		// s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
