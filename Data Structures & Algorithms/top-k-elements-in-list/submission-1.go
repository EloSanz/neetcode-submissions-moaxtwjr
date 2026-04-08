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

    sort.Slice(pairs, func (i, j int) bool {

        return pairs[i].count > pairs[j].count
    })
    result := make([]int, k)
    for i:=0 ; i < k ; i++ { 
        result[i] = pairs[i].num
    }
    return result
}
