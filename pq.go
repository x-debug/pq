package pq

//PriorityQueue is an implementation of the priority queue, it can build a min/max order queue
type PriorityQueue struct {
	len  int
	data []Item
}

//Item is all elements must implement an interface
type Item interface {
	Less(item Item) bool
}

//NewPriorityQueue construct a new priority queue
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	pq.len = 0
	pq.data = make([]Item, 0)
	return pq
}

//Push an element to PQ, it adjusts orders of all items automatically
func (pq *PriorityQueue) Push(ele Item) {
	pq.data = append(pq.data, ele)
	pq.shiftUp(pq.len)
	pq.len += 1
}

//Len return length of PQ
func (pq *PriorityQueue) Len() int {
	return pq.len
}

//Pop removes an element, return an element to the caller, it adjusts orders of all items automatically
func (pq *PriorityQueue) Pop() Item {
	if pq.len == 0 {
		return nil
	}

	item := pq.data[0]
	pq.data[0] = pq.data[pq.len-1]
	pq.data = pq.data[0 : pq.len-1]
	pq.len -= 1
	pq.shiftDown(0)
	return item
}

//Remove also see Pop method, but it can select anywhere
func (pq *PriorityQueue) Remove(k int) Item {
	if k < 0 || k >= pq.len {
		return nil
	}

	item := pq.data[k]
	pq.len -= 1
	pq.data[k] = pq.data[pq.len]
	pq.shiftUp(k)
	pq.shiftDown(k)
	return item
}

func (pq *PriorityQueue) swap(p int, q int) {
	tmp := pq.data[p]
	pq.data[p] = pq.data[q]
	pq.data[q] = tmp
}

func (pq *PriorityQueue) shiftUp(k int) {
	for {
		p := (k - 1) / 2

		if k == 0 {
			return
		}
		ele := pq.data[k]
		if ele.Less(pq.data[p]) {
			return
		}

		pq.swap(k, p)
		k = p
	}
}

func (pq *PriorityQueue) shiftDown(k int) {
	for {
		lc := 2*k + 1 //left child
		rc := 2*k + 2 //right child

		s := k
		if lc < pq.len {
			if pq.data[s].Less(pq.data[lc]) {
				s = lc
			}
		}

		if rc < pq.len {
			if pq.data[s].Less(pq.data[rc]) {
				s = rc
			}
		}

		if s == k {
			return
		}

		pq.swap(k, s)
		k = s
	}
}
