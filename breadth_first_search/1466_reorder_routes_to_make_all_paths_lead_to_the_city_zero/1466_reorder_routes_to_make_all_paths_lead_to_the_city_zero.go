package main 

func minReorder(n int, connections [][]int) int {
    from, to := 1, 2
    m := map[int]map[int]int{}
    for i := 0; i < n; i++ {
        m[i] = make(map[int]int)
    }
    for _, c := range connections {
        m[c[0]][c[1]] = from
        m[c[1]][c[0]] = to
    }

    modify := 0
    stack := []int{0}
    visited := map[int]bool{0: true}
    for i := 0; i < len(stack); i++ {
        for k, v := range m[stack[i]] {
            if visited[k] {
                continue
            }
            if v == from {
                modify++
            }
            stack = append(stack, k)
            visited[k] = true
        }
    }
    return modify
}