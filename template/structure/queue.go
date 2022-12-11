package structure

type Item struct {
	priority int
	value    int
	index    int
}

type PriorityQueue struct {
	count int
	items []*Item
	less  func([]*Item, int, int) bool
}

func NewPriorityQueue(items []*Item, less func([]*Item, int, int) bool) *PriorityQueue {
	return &PriorityQueue{
		count: len(items),
		items: items,
		less:  less,
	}
}

func (pq PriorityQueue) Len() int { return pq.count }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.less(pq.items, i, j)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.count++
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.count--
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	pq.items = old[0 : n-1]
	return item
}

func Desc(items []*Item, i, j int) bool {
	return items[i].priority > items[j].priority
}

func Asc(items []*Item, i, j int) bool {
	return items[i].priority < items[j].priority
}
