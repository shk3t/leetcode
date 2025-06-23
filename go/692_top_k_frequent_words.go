package main

import (
	"container/heap"
)

type WordFrequencyEntry struct {
	word string
	freq int
}

type WFHeap []WordFrequencyEntry

func (h WFHeap) Len() int {
	return len(h)
}

func (h WFHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq || h[i].freq == h[j].freq && h[i].word > h[j].word
}

func (h WFHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WFHeap) Push(x any) {
	*h = append(*h, x.(WordFrequencyEntry))
}

func (h *WFHeap) Pop() any {
	n := len(*h) - 1
	val := (*h)[n]
	*h = (*h)[:n]
	return val
}

func (h *WFHeap) PeekMin() WordFrequencyEntry {
	return (*h)[0]
}

func topKFrequent(words []string, k int) []string {
	wordFrequencyMap := map[string]int{}

	frequentWords := new(WFHeap)

	for _, word := range words {
		if _, ok := wordFrequencyMap[word]; ok {
			wordFrequencyMap[word]++
		} else {
			wordFrequencyMap[word] = 1
		}
	}

	for word, freq := range wordFrequencyMap {
		if frequentWords.Len() < k {
			heap.Push(frequentWords, WordFrequencyEntry{word, freq})
		} else if minEntry := frequentWords.PeekMin(); freq > minEntry.freq ||
			freq == minEntry.freq && word < minEntry.word {
			heap.Pop(frequentWords)
			heap.Push(frequentWords, WordFrequencyEntry{word, freq})
		}
	}

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(frequentWords).(WordFrequencyEntry).word
	}
	return result
}

// Tip: Sometimes it's better to process data sequentially
