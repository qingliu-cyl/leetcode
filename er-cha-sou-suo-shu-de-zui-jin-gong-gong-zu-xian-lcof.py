
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.res = None
        self.p = self.q = 0

    def search(self, node: 'TreeNode'):
        if node is None:
            return
        if node.val > self.q:
            self.search(node.left)
        elif node.val < self.p:
            self.search(node.right)
        else:
            self.res = node

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if q is None:
            return p

        if p is None:
            return q

        if root is None:
            return None

        if p.val < q.val:
            self.p = p.val
            self.q = q.val
        else:
            self.p = q.val
            self.q = p.val


        self.search(root)

        return self.res

if  __name__ == "__main__":
    node6 = TreeNode(6)
    node2 = TreeNode(2)
    node8 = TreeNode(8)
    node0 = TreeNode(0)
    node4 = TreeNode(4)
    node7 = TreeNode(7)
    node3 = TreeNode(3)
    node9 = TreeNode(9)
    node5 = TreeNode(5)

    node6.left = node2
    node6.right = node8

    node2.left = node0
    node2.right = node4

    node8.left = node7
    node8.right = node9

    node4.left = node3
    node4.right = node5

    s = Solution()
    s.lowestCommonAncestor(node6, TreeNode(3), TreeNode(5))


