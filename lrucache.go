type Queue struct {
	El []int
}

func (s *Queue) Add(val int) {
	s.El = append(s.El, val)
}

func (s *Queue) Pop() int {
	var temp = s.El[0]
	s.El = s.El[1:]
	return temp
}

func (s *Queue) PopLast() int {
	var l = len(s.El)
	var temp = s.El[l-1]
	s.El = s.El[:l-1]
	return temp
}

func (s *Queue) PopValue(val int) {
	var target = 0
	for i, v := range s.El {
		if v == val {
			target = i
			break
		}
	}
	s.El = append(s.El[:target], s.El[target+1:]...)
}

func (s *Queue) AddFront(val int) {
	if len(s.El) <= 1 {
		s.Add(val)
	} else {
		s.El = append([]int{val}, s.El[1:]...)
	}
}

type LRUCache struct {
	Capacity int
	HasFull  bool
	Index    int
	Caches   map[int]int
	Queue
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Caches:   map[int]int{},
		Queue:    Queue{El: []int{}},
	}
}

func (this *LRUCache) Get(key int) int {
	if v, f := this.Caches[key]; f {
		this.Queue.PopValue(key)
		this.Queue.Add(key)
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, f := this.Caches[key]; f {
		this.Queue.PopValue(key)
		this.Queue.Add(key)
	} else {
		if !(len(this.Queue.El) != this.Capacity) {
			delete(this.Caches, this.Queue.Pop())
		}
		this.Queue.Add(key)
	}

	this.Caches[key] = value
}