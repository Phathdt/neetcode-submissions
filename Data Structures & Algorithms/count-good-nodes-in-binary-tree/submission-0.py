# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        ans = 0

        def dfs(node, currMax):
            nonlocal ans

            if not node:
                return

            if node.val >= currMax:
                ans += 1

            dfs(node.left, max(currMax, node.val))
            dfs(node.right, max(currMax, node.val))

        dfs(root, float("-inf"))
        return ans
