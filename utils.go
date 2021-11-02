package bst

func getIndex(i indexer, key string) (index int, match bool) {
	sz := i.Len()
	if sz == 0 {
		return
	}

	start := 0
	end := sz - 1
	index = sz / 2

	for {
		ref := i.getKey(index)
		switch {
		case key == ref:
			match = true
			return
		case key < ref:
			end = index
		case key > ref:
			start = index
		}

		switch {
		case start == end:
			if key > ref {
				index++
			}

			return
		case end-start > 1:
			index = (start + end) / 2
		case start == index:
			start++
			index++
		case end == index:
			end--
			index--
		}
	}
}

type indexer interface {
	getKey(index int) string
	Len() int
}
