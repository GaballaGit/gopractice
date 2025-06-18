package intermediate 

import "fmt"

func main(){
	add := adding()
	fmt.Println("The value of i is:", add())
	fmt.Println("The value of i is:", add())
	fmt.Println("The value of i is:", add())
	fmt.Println("The value of i is:", add())
	fmt.Println("The value of i is:", add())

	subtractor := func() func(int) int {
		j := 100
		fmt.Println("The value of j is:", j)
		return func(x int) int {
			j -= x
			return j
		}
	}()

	fmt.Println("j subtracted by 5 is:", subtractor(5))
	fmt.Println("j subtracted by 3 is:", subtractor(3))
	fmt.Println("j subtracted by 8 is:", subtractor(8))
	fmt.Println("j subtracted by 1 is:", subtractor(1))

}


func adding() func() int{
	i := 0
	fmt.Println("The value of i is:", i)
	return func() int{
		i++
		return i
	}
}
