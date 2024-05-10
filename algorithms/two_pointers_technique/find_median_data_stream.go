package two_pointers_technique

import "fmt"

type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.data)
	middleIdx := (n - 1) / 2
	if n%2 != 0 {
		return float64(this.data[middleIdx])
	}
	if middleIdx+1 < n {
		return float64(this.data[middleIdx]+this.data[middleIdx+1]) / 2
	}
	return float64(this.data[middleIdx])
}

func TestFindMedian() {
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.data)
	obj.AddNum(2)
	fmt.Println(obj.data)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.data)
	fmt.Println(obj.FindMedian())

}
