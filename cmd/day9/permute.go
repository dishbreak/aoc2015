package main

func Permute(target []string) [][]string {
	acc := make([][]string, 0)
	acc = permute(target, 0, acc)
	return acc
}

func permute(target []string, j int, acc [][]string) [][]string {
	if j == len(target)-1 {
		result := make([]string, len(target))
		copy(result, target)
		return append(acc, result)
	}

	for i := j; i < len(target); i++ {
		target[j], target[i] = target[i], target[j]
		acc = permute(target, j+1, acc)
		target[j], target[i] = target[i], target[j]
	}
	return acc
}
