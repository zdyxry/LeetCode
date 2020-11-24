package main 

import "sort"

type item struct {
    v int
    l int
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    if len(values) == 0 || num_wanted == 0 {
        return 0
    }
    
    items := make([]item, len(values))
    for i := 0; i < len(values); i ++ {
        items[i] = item{values[i], labels[i]}
    }
    
    sort.Slice(items, func(i, j int)bool{
        return items[i].v > items[j].v
    })
    
    res := 0
    cnt := make(map[int]int)
    total := 0
    for i := 0; i < len(items); i++ {
        
        if cnt[items[i].l] == use_limit {
            continue
        }
        
        res += items[i].v
        cnt[items[i].l]++
        total++
        
        if total == num_wanted {
            break
        }
    }
    
    return res
}