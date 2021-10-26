package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := []int{7408,7410,7411,7412,7413,7414,7415,7416,7417,7418,7419,7420,7421,7422,7423,7424,7425,7426,7427,7428,7429,7430,7431,7432,7433,7434,7435,7436,7437,7438,7439,7440,7441,7442,7443,7444,7445}
	for i := 0; i < len(numbers); i++ {
		file, err := os.Create(strconv.Itoa(numbers[i]) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		var padded string = fmt.Sprintf("%05d", numbers[i] - 3)
		fmt.Fprintln(file, "---")
		fmt.Fprintln(file, "layout: flash")
		fmt.Fprintln(file, "---")
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "<embed src=\"https://www.homestuck.com/flash/hs2/"+padded+"/"+padded+".swf\" quality=\"high\" width=\"100%\" height=\"100%\" allowScriptAccess=\"sameDomain\" allowFullScreen=\"false\" type=\"application/x-shockwave-flash\" pluginspage=\"http://www.macromedia.com/go/getflashplayer\" />")
	}
}
