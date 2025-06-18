package intermediate

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: Can not calculate a negative square root")
	}
	res := math.Sqrt(x)
	return res, nil
}

func main() {
	res, err := sqrt(64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	res, err = sqrt(-64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
