class Solution:
    def findEvenNumbers(self, digits: List[int]) -> List[int]:
        all_elements = [i for i in range(100, 1000) if i%2==0]
        result = []
        
        for i in all_elements:
            current_i = [int(x) for x in str(i)]
            digits_copy = digits.copy()
            
            for x in range(3):
                if current_i[x] not in digits_copy:
                    break
                else:
                    digits_copy.remove(current_i[x])
                    
                if x == 2:
                    result.append(i) 
        
        
        return sorted(result)
