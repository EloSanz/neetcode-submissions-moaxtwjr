func isAnagram(s string, t string) bool {
    if len(t) != len(s){ 
        return false
    }

    var counter [26]int

    for i:=0 ; i < len(s) ; i++ {
        counter[s[i] - 'a'] ++
        counter[t[i] - 'a'] --
    }
    for val := range counter {
        if counter[val] != 0 {
            return false
        }
    }
    return true
}
