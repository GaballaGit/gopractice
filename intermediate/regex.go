package intermediate 

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9.+%_-]+@[a-zA-Z0-9._-]+\.[a-zA-Z]{2,}`)

	email1 := "example@email.com"

	email2 := "adofan42_3@gmail.c"

	fmt.Println(re.MatchString(email1))
	fmt.Println(re.MatchString(email2))

	re = regexp.MustCompile(`(\d{2})-(\d{2})-(\d{4})`)

	date1 := "21-05-2024"
	date2 := "01-12-1998"

	d1 := re.FindStringSubmatch(date1)
	d2 := re.FindStringSubmatch(date2)


	fmt.Println(d2[3])
	fmt.Println(d1[3])
}


