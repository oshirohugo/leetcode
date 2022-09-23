package design

type OrderedStream struct {
	stream map[int]string
	index  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make(map[int]string, n),
		index:  0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	key := idKey - 1
	this.stream[key] = value

	s := []string{}
	if this.index == key {
		for {
			if this.stream[this.index] != "" {
				s = append(s, this.stream[this.index])
			}
			if _, ok := this.stream[this.index]; ok {
				this.index++
				continue
			}
			break
		}

	}

	return s
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
