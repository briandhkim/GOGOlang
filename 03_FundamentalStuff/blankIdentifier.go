//every variable must be used
//blank identifier allows user to tell comiler something is not being used

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://briandhkim.fun/")
	//the underscore is blank identifier
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
