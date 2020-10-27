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

	songs := []song{
		{title: "wonderwall", artist: "oasis"},
		{title: "super sonic", artist: "oasis"},
	}

	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}

	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	test := rock.songs[0]
	fmt.Printf("%+v\n", test)

	test.title = "Test title"
	fmt.Printf("Original songs : %+v\n", songs[0])
	fmt.Printf("Test 2 : %+v\n", test)

}
