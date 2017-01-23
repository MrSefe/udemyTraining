package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Blank identifiers used to deter error checking
	res, _ := http.Get("http://jarjestot.uta.fi/typo/index.php")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

}
