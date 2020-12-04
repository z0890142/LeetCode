package WordLadder

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var result int
	wordMap := make(map[string]struct{})

	for _, v := range wordList {
		wordMap[v] = struct{}{}
	}

	if _, exists := wordMap[endWord]; !exists {
		return 0
	}

	queue := list.New()
	queue.PushBack(beginWord)
	delete(wordMap, beginWord)

	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			word := queue.Remove(queue.Front()).(string)
			chars := []byte(word)
			for i := range chars {
				b := chars[i]

				for j := 0; j < 26; j++ {
					chars[i] = byte('a' + j)
					if string(chars) == endWord {
						return result + 1
					}

					if _, exists := wordMap[string(chars)]; exists {
						queue.PushBack(string(chars))
						delete(wordMap, string(chars))
					}
				}
				chars[i] = b
			}
		}
		result++
	}
	return 0

}
