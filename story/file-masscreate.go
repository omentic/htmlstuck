package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := []int{3297, 3321, 3438, 3520, 3679, 3695, 3696, 3717, 3718, 3725, 3730, 3744, 3760, 3851, 3860, 3952, 4085, 4109, 4111, 4113, 4229, 4275, 4282, 4373, 4390, 4486, 4541, 4572, 4617, 4665, 4815, 4820, 4825, 4827, 4942, 4944, 5027, 5238, 5252, 5261, 5263, 5308, 5373, 5374, 5375, 5376, 5377, 5398, 5426, 5427, 5438, 5440, 5471, 5472, 5473, 5474, 5485, 5494, 5512, 5655, 5711, 5712, 5714, 5726, 5740, 5759, 5763, 5764, 5776, 5777, 5787, 5981, 5997, 6066, 6073, 6231, 6243, 6278, 6400, 6552, 6652, 6901, 6943, 7086, 7098, 7100, 7101, 7405, 7409, 7448, 7449, 7635, 7823, 7928, 7959, 8087, 8127, 8129, 8130}
	for i := 0; i < len(numbers); i++ {
		file, err := os.Create(strconv.Itoa(numbers[i]) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		var padded string = fmt.Sprintf("%05d", numbers[i]-3)
		fmt.Fprintln(file, "---")
		fmt.Fprintln(file, "layout: flash")
		fmt.Fprintln(file, "---")
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "<embed src=\"https://www.homestuck.com/flash/hs2/"+padded+"/"+padded+".swf\" quality=\"high\" width=\"100%\" height=\"100%\" allowScriptAccess=\"sameDomain\" allowFullScreen=\"false\" type=\"application/x-shockwave-flash\" pluginspage=\"http://www.macromedia.com/go/getflashplayer\" />")
	}
}
