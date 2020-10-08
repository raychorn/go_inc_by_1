package main

func inc_by_1(items int[]) int[] {
	fmt.Println(items)
	return items
}

func main() {
	fmt.Println(inc_by_1([1, 0, 1]))
}