1418. Display Table of Food Orders in a Restaurant


Medium


Given the array orders, which represents the orders that customers have done in a restaurant. More specifically orders[i]=[customerNamei,tableNumberi,foodItemi] where customerNamei is the name of the customer, tableNumberi is the table customer sit at, and foodItemi is the item customer orders.

Return the restaurant's “display table”. The “display table” is a table whose row entries denote how many of each food item each table ordered. The first column is the table number and the remaining columns correspond to each food item in alphabetical order. The first row should be a header whose first column is “Table”, followed by the names of the food items. Note that the customer names are not part of the table. Additionally, the rows should be sorted in numerically increasing order.

 

Example 1:

```
Input: orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
Output: [["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]] 
Explanation:
The displaying table looks like:
Table,Beef Burrito,Ceviche,Fried Chicken,Water
3    ,0           ,2      ,1            ,0
5    ,0           ,1      ,0            ,1
10   ,1           ,0      ,0            ,0
For the table 3: David orders "Ceviche" and "Fried Chicken", and Rous orders "Ceviche".
For the table 5: Carla orders "Water" and "Ceviche".
For the table 10: Corina orders "Beef Burrito". 
```

Example 2:

```
Input: orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]]
Output: [["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]] 
Explanation: 
For the table 1: Adam and Brianna order "Canadian Waffles".
For the table 12: James, Ratesh and Amadeus order "Fried Chicken".
```

Example 3:

```
Input: orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melissa","2","Soda"]]
Output: [["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]
```

Constraints:

1 <= orders.length <= 5 * 10^4   
orders[i].length == 3  
1 <= customerNamei.length, foodItemi.length <= 20  
customerNamei and foodItemi consist of lowercase and uppercase English letters and the space character.  
tableNumberi is a valid integer between 1 and 500.


## 方法


```go
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
	for k,_:=range fmp{
		foods=append(foods,k)
	}
	for k,_:=range tabs{
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
```


```python
class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        data = defaultdict(lambda: defaultdict(int))
        food_items = set()
        for order in orders:
            table, food = int(order[1]), order[2]
            data[table][food] += 1
            food_items.add(food)
        sorted_food_items = sorted(list(food_items))
        result = [["Table"]+sorted_food_items]
        for table in sorted(data.keys()):
            result.append([str(table)]+[str(data[table][f]) for f in sorted_food_items])
        return result
```