package array

type dynamicArray struct {
	items    []interface{} // holds the elements
	size     int           // stores current number of elements
	capacity int           // store current capacity
}

func NewDynamicArray(initialCapacity int) *dynamicArray {

	if initialCapacity < 1 {
		initialCapacity = 1
	}

	return &dynamicArray{
		items:    make([]interface{}, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

func (d *dynamicArray) Resize(capacity int) {
	newItems := make([]interface{}, capacity)
	copy(newItems, d.items)
	d.items = newItems
}

func (d *dynamicArray) Add(item interface{}) {
	if d.size == d.capacity {
		//d.Resize(d.items)
	}

	//d[d.size] = item
}

func (d *dynamicArray) Get(item interface{}) {

}
