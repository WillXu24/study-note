package main

import "fmt"

func main() {
	bubbleSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	selectSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	insertSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	shellSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	heapSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	mergeSort0([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
	quickSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})
}

/**********************************************************************************************************************/
func bubbleSort0(a []int) {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				count++
			}
		}
	}
	fmt.Println("Bubble Sort 0:", a, "Swap Count:", count)
}

func bubbleSort1(a []int) {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				count++
			}
		}
	}
	fmt.Println("Bubble Sort 1:", a, "Swap Count:", count)
}

func bubbleSort(a []int) {
	count := 0
	flag := true
	for i := 0; i < len(a) && flag; i++ {
		flag = false
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				count++
				flag = true
			}
		}
	}
	fmt.Println("Bubble Sort:", a, "Swap Count:", count)
}

/**********************************************************************************************************************/
func selectSort(a []int) {
	count := 0
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
			count++
		}
	}
	fmt.Println("Select Sort:", a, "Swap Count:", count)
}

/**********************************************************************************************************************/
func insertSort(a []int) {
	count := 0
	for i := 1; i < len(a); i++ {
		temp := a[i]
		position := i
		for j := i; j > 0 && a[j-1] > temp; j-- {
			a[j] = a[j-1]
			count++
			position = j - 1
		}
		a[position] = temp
	}
	fmt.Println("Insert Sort:", a, "Swap Count:", count)
}

/**********************************************************************************************************************/
func shellSort(a []int) {
	count, gap := 0, 0
	for ; gap < len(a); gap = gap*2 + 1 {
	}
	for gap = (gap - 1) / 2; gap > 0; gap = (gap - 1) / 2 {
		for i := gap; i < len(a); i++ {
			temp := a[i]
			position := i
			for j := i;j-gap >= 0 && a[j-gap] > temp; j = j - gap {
				a[j] = a[j-gap]
				count++
				position = j - gap
			}
			a[position] = temp
		}
	}
	fmt.Println("Shell Sort:", a, "Swap Count:", count)
}

/**********************************************************************************************************************/
func heapSort(a []int) {
	count:=0
	for i:=len(a)/2;i>=0;i--{
		count += heapAdjust(a, i)
	}
	for i:=len(a)-1;i>0;i--{
		a[0],a[i]=a[i],a[0]
		count +=heapAdjust(a[:i],0)
	}
	fmt.Println("Heap Sort:", a, "Swap Count:",count)
}

func heapAdjust(a []int,root int) int {
	count:=0
	for i:=2*root+1;i<len(a);i=i*2+1{
		if i+1<len(a)&&a[i]<a[i+1]{
			i++
		}
		if a[root]>a[i]{
			break
		}
		a[root],a[i]=a[i],a[root]
		count++
		root=i
	}
	return count
}

/**********************************************************************************************************************/
func mergeSort0(a []int)  {
	res:=mSort(a)
	fmt.Println("Merge Sort:", res)
}

func mSort(a []int)[]int {
	if len(a)<2{
		return a
	}
	m:=len(a)/2
	l:=mSort(a[:m])
	r:=mSort(a[m:])
	return merge(l,r)
}

func merge(a, b []int)  []int  {
	var res []int
	i,j:=0,0
	for i<len(a)&&j<len(b){
		if a[i] <= b[j]{
			res=append(res,a[i])
			i++
		}else {
			res=append(res,b[j])
			j++
		}
	}
	res=append(res,a[i:]...)
	res=append(res,b[j:]...)
	return res
}

/**********************************************************************************************************************/
func quickSort(a []int) {
	qSort(a,0,len(a)-1)
	fmt.Println("Quick Sort:", a)
}
func qSort(a []int,left,right int)  {
	if left>=right{
		return
	}
	p:=partition(a,left,right)
	qSort(a,left,p-1)
	qSort(a,p+1,right)
}

func partition(a []int,left,right int) int {
	p:=a[left]
	for left!=right{
		for left<right&&p<a[right]{
			right--
		}
		a[left],a[right] = a[right],a[left]
		for left<right&&a[left]<p{
			left++
		}
		a[left],a[right] = a[right],a[left]
	}
	return left
}