package output

import (
	"fmt"
	"os"
)

// ///////////////////////////////////////////FUNCTION OUTPUT//////////////////////////////////////////////////////////////////////
func OutPut(filename string, outputline []string) {
	f, e := os.Create(filename)
	if e != nil {
		fmt.Println("Please, write the flag \"--output=\", followed by name of the file to create")
		os.Exit(0)
	}
	defer f.Close()
	for _, v := range outputline {
		fmt.Fprintln(f, v)
	}

}
