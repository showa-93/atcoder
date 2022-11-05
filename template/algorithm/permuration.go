// 順列
package algorithm

// 引数の配列をもとに次の順列を生成する
// 最初の呼び出しで渡す引数の配列は降順にソート済みであること
func NextPermutation(list []int) bool {
	if len(list) <= 2 {
		return false
	}

	var i int
	for i = len(list) - 2; i >= 0; i-- {
		if list[i] < list[i+1] {
			break
		}
	}

	if i < 0 {
		return false
	}

	var j int
	for j = len(list) - 1; j >= i; j-- {
		if list[i] < list[j] {
			break
		}
	}

	list[i], list[j] = list[j], list[i]

	for p, q := i+1, len(list)-1; p < q; p, q = p+1, q-1 {
		list[p], list[q] = list[q], list[p]
	}

	return true
}

func PrevPermutation(list []int) bool {
	if len(list) <= 2 {
		return false
	}

	i := len(list) - 2
	for ; i >= 0; i-- {
		if list[i] > list[i+1] {
			break
		}
	}

	if i < 0 {
		return false
	}

	j := len(list) - 1
	for ; j >= i; j-- {
		if list[i] > list[j] {
			break
		}
	}

	list[i], list[j] = list[j], list[i]

	for p, q := i+1, len(list)-1; p < q; p, q = p+1, q-1 {
		list[p], list[q] = list[q], list[p]
	}

	return true
}
