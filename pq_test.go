package pq

import "testing"

type MinIntItem struct {
	value int
}

func (item MinIntItem) Less(other Item) bool {
	return item.value > other.(MinIntItem).value
}

type MaxIntItem struct {
	value int
}

func (item MaxIntItem) Less(other Item) bool {
	return item.value < other.(MaxIntItem).value
}

func createMinPQ() *PriorityQueue {
	pq := NewPriorityQueue()
	pq.Push(MinIntItem{value: 3})
	pq.Push(MinIntItem{value: 9})
	pq.Push(MinIntItem{value: 1})
	pq.Push(MinIntItem{value: 7})
	pq.Push(MinIntItem{value: 5})

	return pq
}

func createMaxPQ() *PriorityQueue {
	pq := NewPriorityQueue()
	pq.Push(MaxIntItem{value: 2})
	pq.Push(MaxIntItem{value: 8})
	pq.Push(MaxIntItem{value: 1})
	pq.Push(MaxIntItem{value: 18})
	pq.Push(MaxIntItem{value: 6})
	return pq
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := createMinPQ()

	if pq.Len() != 5 {
		t.Fail()
	}

	root := pq.data[0].(MinIntItem)
	if root.value != 1 {
		t.Fail()
	}
}

func TestPriorityQueue_Push2(t *testing.T) {
	pq := createMaxPQ()

	if pq.Len() != 5 {
		t.Fail()
	}

	root := pq.data[0].(MaxIntItem)
	if root.value != 18 {
		t.Fail()
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := createMinPQ()

	item := pq.Pop().(MinIntItem)
	if item.value != 1 {
		t.Fail()
	}

	item = pq.Pop().(MinIntItem)
	if item.value != 3 {
		t.Fail()
	}
}

func TestPriorityQueue_Pop2(t *testing.T) {
	pq := createMaxPQ()

	item := pq.Pop().(MaxIntItem)
	if item.value != 18 {
		t.Fail()
	}

	item = pq.Pop().(MaxIntItem)
	if item.value != 8 {
		t.Fail()
	}
}

func TestPriorityQueue_Remove(t *testing.T) {
	pq := createMinPQ()

	pq.Remove(2)
	if pq.Len() != 4 {
		t.Fail()
	}

	pq.Pop()
	pq.Pop()
	item := pq.Pop().(MinIntItem)
	if item.value != 7 {
		t.Fail()
	}
}

func TestPriorityQueue_Remove2(t *testing.T) {
	pq := createMaxPQ()

	pq.Remove(2)
	if pq.Len() != 4 {
		t.Fail()
	}

	pq.Pop()
	pq.Pop()
	item := pq.Pop().(MaxIntItem)
	if item.value != 6 {
		t.Fail()
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := createMinPQ()

	if pq.Len() != 5 {
		t.Fail()
	}

	pq.Push(MinIntItem{value: 11})
	pq.Push(MinIntItem{value: 15})
	if pq.Len() != 7 {
		t.Fail()
	}

	pq.Pop()
	if pq.Len() != 6 {
		t.Fail()
	}
}
