package main

type DynamicArray struct{ 
	size int
	capacity int
	data []int
}

func (d *DynamicArray)get(index int) int {
	if 0 <= index && index < d.size { 
		return d.data[index];
	} else {
		panic("wrong index")
	}
}

func (d *DynamicArray)set(index int, val int) {
	if 0 <= index && index < d.size { 
		d.data[index] = val
	} else {
		panic("wrong index")
	}
}

func (d *DynamicArray)push(val int)  {
	if d.data != nil {
		if d.size == d.capacity {
			d.resize()
			d.data[d.size] = val
			d.size++
		} else{
			d.data[d.size] = val
			d.size++
		}
	} else {
		d.data = make([]int, 4)
		d.data[d.size] = val
	}
}

func(d *DynamicArray) pop() int {
	d.size--
	val := d.data[d.size]
	return val
}

func (d *DynamicArray)insertAt(index int, val int)  {
	if 0 <= index && index <= d.size {
		if d.size == d.capacity {
			d.resize()
		}
		for i := d.size; i > index; i-- {
			d.data[i] = d.data[i-1]
		}
		d.data[index] = val
		d.size++
	} else {
		panic("wrong index")
	}
}

func (d *DynamicArray)removeAt(index int)  {
	if 0 <= index && index < d.size {
		for i := index + 1; i < d.size; i++ {
			d.data[i-1] = d.data[i]
		}
		d.size--
	} else {
		panic("wrong index")
	}
}

func (d *DynamicArray)Size() int {
	return d.size;
}

func (d *DynamicArray)Capacity() int {
	return d.capacity;
}

func (d *DynamicArray)resize()  {
	newArr := make([]int , d.capacity * 2)
	for i := 0; i < d.capacity; i++ {
		newArr[i] = d.data[i]
	}
	d.data = newArr
	d.capacity *= 2
}

func main()  {
	dynamicArray := &DynamicArray{
		size: 0,
		capacity: 4,
		data: make([]int, 4),
	}
	dynamicArray.push(1)
	dynamicArray.push(2)
	dynamicArray.push(3)
	dynamicArray.push(4)
	dynamicArray.push(5)
	// for i := 0; i < dynamicArray.Size(); i++ {
	// 	println(dynamicArray.get(i))	
	// }
	println("capacity ", dynamicArray.Capacity())
	println("size ", dynamicArray.Size())
	dynamicArray.insertAt(3 , 7)
	for i := 0; i < dynamicArray.Size(); i++ {
		println(dynamicArray.get(i))	
	}
	dynamicArray.removeAt(3)
	for i := 0; i < dynamicArray.Size(); i++ {
		println(dynamicArray.get(i))	
	}
	dynamicArray.pop()
	for i := 0; i < dynamicArray.Size(); i++ {
		println(dynamicArray.get(i))	
	}
}