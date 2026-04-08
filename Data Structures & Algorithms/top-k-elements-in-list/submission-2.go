import "slices"

func topKFrequent(nums []int, k int) []int {

    countMap := make(map[int]int)

    for _, n := range nums{
        countMap[n]++ 
    }

    type pair struct{
        num int
        count int
    }

    pairs := make([] pair, 0, len(countMap))

    for num, count := range countMap {
        pairs = append(pairs, pair{num, count})
    }

    slices.SortFunc(pairs, func(a, b pair) int {
        return b.count - a.count
    })

    result := make([]int, k)
    for i:=0 ; i < k ; i++ { 
        result[i] = pairs[i].num
    }
    return result
}
