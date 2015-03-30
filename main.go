package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"unicode/utf8"
)

type Book struct {
	Author  string
	Title   string
}

type Catalog struct {
	Id    int
	Name  string
	Books []Book
}

type Author struct {
	Name        string
	InBloggers  int
	InBookBlogs int
	InJournal   int
}

var catalogs []Catalog

var authors []Author

var maxLengthOfAuthorName int

func collectDatas(csvFiles []string) {
	
	for i, csvFileName := range csvFiles {
		csvFile, err := os.Open("csv/" + csvFileName)
		
		if err != nil {
			fmt.Println(err)
            return 
        }
        
        tmpCatalog := Catalog{Id:i, Name:csvFileName}

		defer csvFile.Close()
		
		reader := csv.NewReader(csvFile)
		reader.Comma = ';'
		//reader.FieldsPerRecord = 2
		
		
		rawCSVdata, err := reader.ReadAll()

		if err != nil {
			 fmt.Println(err)
			 os.Exit(1)
		}

		// sanity check, display to standard output
		for _, each := range rawCSVdata {
			tmp := Book{each[0], each[1]}
			tmpCatalog.Books = append(tmpCatalog.Books, tmp)
			
			insertAuthor(each[0])
			if maxLengthOfAuthorName < utf8.RuneCountInString(each[0]) {
				maxLengthOfAuthorName = utf8.RuneCountInString(each[0])
			}
			fmt.Printf("CATALOG: %s; AUTHOR: %s; TITLE: %s\n", tmpCatalog.Name, tmp.Author, tmp.Title)
		}
		
		catalogs = append(catalogs, tmpCatalog)
    }
}

func insertAuthor(name string) {
	
	for i, author := range authors {
		if name < author.Name {
			tmp := []Author{}
			tmp = append(tmp, Author{Name: name})
			
			authors = append(authors[:i], append(tmp, authors[i:]...)...)
			fmt.Printf("Author: %s\n", name)
			
			return
		}
		
		if name == author.Name {
			
			return
		}
	}
	
	authors = append(authors, Author{Name:name})
	fmt.Printf("Author: %s\n", name)
}

func addCriticToAuthor(name string, index int) {
	var toBeSearchedAuthor *Author
	for indexOfAuthor, author := range authors {
		if name == author.Name {
			
			toBeSearchedAuthor = &authors[indexOfAuthor]
		}
	}
	
	// újságok
	if 5 >= index {
		toBeSearchedAuthor.InJournal++
	}
	
	// blogok
	if 8 >= index && 6 <= index {
		toBeSearchedAuthor.InBookBlogs++
	}
	
	// bloggerek
	if 8 < index {
		toBeSearchedAuthor.InBloggers++
	}
}


func collectAuthorCritics() {
	for i, catalog := range catalogs {
	
		for _, book := range catalog.Books {
			addCriticToAuthor(book.Author, i)
			
			authorName := book.Author
		
			for i := 0; i < maxLengthOfAuthorName - utf8.RuneCountInString(book.Author); i++ {
				authorName += " "
			}
			
			fmt.Printf("%s szerepel a %s katalugusban (könyv: %s)\n", authorName, catalog.Name, book.Title)
			
		}
		
	}
}

 func main() {

	csvFiles := []string{
		"és.csv",
		"magyarnarancs.csv",
		"jelenkor.csv",
		"holmi.csv",
		"kaligram.csv",
		"kortars.csv",
		"könyvesblog.csv",
		"kötve-füzve.csv",
		"litera.csv",
		"amadea.csv",
		"dóri.csv",
		"katacita.csv",
		"könygalaxis.csv",
		"miamona.csv",
		"nikkincs.csv",
		"olvasónapló.csv",
		"pupilla.csv",
		"szilvamag.csv",
		"zakkant.csv",
	}
	
	collectDatas(csvFiles)
	
	fmt.Printf("-----------------------------------\n")
	
	collectAuthorCritics()
	
	fmt.Printf("----------------------------------------------------\n")
	
	c := 0
	for _, author := range authors {
		c++
		authorName := author.Name
		
		for i := 0; i < maxLengthOfAuthorName - utf8.RuneCountInString(author.Name); i++ {
			authorName += " "
		}
		
		fmt.Printf("%s szerepel %d folyóiratban, %d könyvesblogban, %d blogban.\n", authorName, author.InJournal, author.InBookBlogs, author.InBloggers)
	}
	fmt.Printf("Összesen:%d\n", c)
	
	fmt.Printf("----------------------------------------------------\n")
	fmt.Printf("Metszet:\n")
	c = 0
	for _, author := range authors {
		in := 0
		if author.InJournal != 0 {
			in++
		}
		if author.InBookBlogs != 0 {
			in++
		}
		if author.InBloggers != 0 {
			in++
		}
		
		if in < 2 {
			continue
		}
		
		c++
		authorName := author.Name
		
		for i := 0; i < maxLengthOfAuthorName - utf8.RuneCountInString(author.Name); i++ {
			authorName += " "
		}
		
		fmt.Printf("%s szerepel %d folyóiratban, %d könyvesblogban, %d blogban.\n", authorName, author.InJournal, author.InBookBlogs, author.InBloggers)
	}
	fmt.Printf("Összesen:%d\n", c)
	
	
	fmt.Printf("----------------------------------------------------\n")
	fmt.Printf("Legtöbb:\n")
	max1 := Author{}
	max2 := Author{}
	max3 := Author{}
	for _, author := range authors {
		
		if "" == author.Name {
			continue
		}
		
		if author.InJournal > max1.InJournal {
			max1 = author
		}
		
		if author.InBookBlogs > max2.InBookBlogs {
			max2 = author
		}
		
		if author.InBloggers > max3.InBloggers {
			max3 = author
		}
	}
	fmt.Printf("Folyóiratok között: %s %d említés.\n", max1.Name, max1.InJournal)
	fmt.Printf("Könyves blogok között: %s %d említés.\n", max2.Name, max2.InBookBlogs)
	fmt.Printf("Bloggerek között: %s %d említés.\n", max3.Name, max3.InBloggers)
	
 }
