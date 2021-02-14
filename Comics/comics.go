/*Comic Mischief
PROJECT
Become the software administrator of the comic book store you own and operate, Comic Mischief.
Use variables to store data about your shop and update those variables to maintain an internal
state of what comics are being added to the store!*/

package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32

	fmt.Print("\n")

	//"Mr. GoToSleep", "Tracey Hatchet", "Jewel Tampson", "DizzyBooks Publishing Inc."
	title, writer, artist, publisher, year, pageNumber, grade = "Mr. GoToSleep", "Tracey Hatchet", "Jewel Tampson", "DizzyBooks Publishing Inc.", 1997, 14, 6.5

	fmt.Println("Title:", title, "\nWriter:", writer, "\nArtist:", artist, "\nPublisher:", publisher, "\nYear of publication:", year, "\nPages:", pageNumber, "\nCondition:", grade, "/10")

	fmt.Print("\n")

	//"Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Publishing Inc."
	title, writer, artist, publisher, year, pageNumber, grade = "Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Publishing Inc.", 2013, 160, 9.0

	fmt.Println("Title:", title, "\nWriter:", writer, "\nArtist:", artist, "\nPublisher:", publisher, "\nYear of publication:", year, "\nPages:", pageNumber, "\nCondition:", grade, "/10")
}
