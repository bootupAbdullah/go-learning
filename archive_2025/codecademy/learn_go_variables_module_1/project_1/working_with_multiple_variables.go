package main

import "fmt"

func main () {

var publisher string
var writer string
var artist string
var title string 
var year int32
var pageNumber int32
var grade float32

publisher = "DizzyBooks Publishing Inc."
writer = "Tracey Hatchet"
artist = "Jewel Thompson"
title = "Mr. GoToSleep"
year = 1997
pageNumber = 14
grade = 6.5

fmt.Println(title, "written by", writer, "drawn by", artist + ".", "Condition :", grade, "number of pages: ", pageNumber, publisher, year)

title = "Epic Vol. 1"
writer = "Ryan N. Shawn"
artist = "Phoebe Paperclips"
publisher = "DizzyBooks Publishing Inc."
year = 2013
pageNumber = 160
grade = 9.0




fmt.Println(title, "written by", writer, "drawn by", artist + ".", "Condition :", grade, "number of pages: ", pageNumber, publisher, year)

// fmt.Println(publisher)
// fmt.Println(publisher, writer, artist, title)