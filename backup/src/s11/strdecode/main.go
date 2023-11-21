// How to decode string
package main

import "fmt"

func main() {

	text := `
	Đường lên xứ Lạng bao xa
	Cách một trái núi với ba quãng đồng
	Ai ơi! đứng lại mà trông
	Kìa núi Thành Lạc, kìa sông Tam Cờ.
	Em chớ thấy anh lắm bạn mà ngờ
	Bụng anh vẫn thẳng như tờ giấy phong...
	`
	/*
		for i := 0; i < len(text); {
			r, size := utf8.DecodeRuneInString(text[i:])
			fmt.Printf("%c", r)

			i += size
		}
		fmt.Println()
	*/

	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}
