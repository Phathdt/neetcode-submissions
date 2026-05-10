from collections import defaultdict 

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        result = defaultdict(list)

        for str in strs: 
            key = tuple(sorted(str))
            result[key].append(str)
        
        return list(result.values())