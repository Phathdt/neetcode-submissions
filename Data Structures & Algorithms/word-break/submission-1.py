from functools import cache

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        words = set(wordDict)
        
        @cache
        def dfs(sub):
            if sub == "":
                return True

            for i in range(1, len(sub) + 1):
                prefix = sub[:i]

                if prefix in words:
                    if dfs(sub[i:]):
                        return True

            return False

        return dfs(s)
