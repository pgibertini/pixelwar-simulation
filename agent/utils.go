package agent

import "math/rand"

func ContainsHobby(hobbies []string, hobby string) bool {
	for _, v := range hobbies {
		if v == hobby {
			return true
		}
	}
	return false
}

func exists(s []*AgentWorker, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func GetManagerIndex(s []*AgentManager, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func remove[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func MakeRandomSliceOfHobbies(hobbies []string) (result []string) {
	result = make([]string, 0)
	for i := 0; i < 1; i++ {
		k := rand.Intn(len(hobbies))
		result = append(result, hobbies[k])
	}
	return
}
