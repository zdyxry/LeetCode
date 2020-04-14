
class Solution:
    def entityParser(self, text: str) -> str:
        dat={"&quot;":"\"",
            "&apos;":"'",
            "&amp;":"&",
            "&gt;":">",
            "&lt;":"<",
            "&frasl;":"/",
             }
        txt=''
        amp_idx,sem_idx=None,None
        for i,e in enumerate(text):
	    	# if & we update the amp_idx 
            if e=="&":
                amp_idx=i
			# if ; we update sem idx	
            if e==";":
                sem_idx=i
		   # if we don't have any amp_idx yet to be tracked just add the curr char from text
            if amp_idx==None:
                txt+=e
			# if we have amp_idx and sem_idx, means we have a contiguous block to compare in dat dictonary
            if amp_idx!=None  and sem_idx!=None:
                key = text[amp_idx:sem_idx+1] # we get that block to compare from text
				# if key in dat then we add the replacement in txt. e.g: &gt replace with >
                if key in dat.keys():
                    txt+=dat[key]
				# but what if we don't have that key in dat? e.g: &ambassador;. so we just add the full string aka key	
                else:
                    txt+=key
				# assign the idx tracker to None to track next blocks
                amp_idx,sem_idx=None,None

        return txt


text = "&amp; is an HTML entity but &ambassador; is not."
res = Solution().entityParser(text)
print(res)