package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
    fmp:=make(map[string]int)
	ods:=make(map[string]int)
	tabs:=make(map[int]int)
	var tables []int
	for _,v:=range orders{
		fmp[v[2]]++
		ts:=v[1]+" "+v[2]
		ods[ts]++
		tas,_:=strconv.Atoi(v[1])
		tabs[tas]++
	}
	var foods []string
	for k :=range fmp{
		foods=append(foods,k)
	}
	for k :=range tabs{
		tables=append(tables,k)
	}
	sort.Strings(foods)
	foods=append([]string{"Table"},foods...)
	//food的映射关系
	mp3:=make(map[int]string)
	for k,v:=range foods{
		mp3[k]=v
	}
	var res [][]string
	res=append(res,foods)
	sort.Ints(tables)
	for _,v:=range tables{
		arr:=make([]string,len(foods))
		arr[0]=strconv.Itoa(v)
		for j:=1;j<len(foods);j++{
			str:=strconv.Itoa(v)+" "+mp3[j]
			if v1,ok:=ods[str];ok{
				arr[j]=strconv.Itoa(v1)
			}else{
				arr[j]="0"
			}
		}
		res=append(res,arr)
	}
	return res

}