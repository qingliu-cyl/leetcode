class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def __init__(self):
        self.res = []

    def search(self, node: 'Node'):
        if node is None:
            return

        self.search(node.left)
        self.res.append(node)
        self.search(node.right)

    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if root is None:
            return None

        self.search(root)

        if len(self.res) == 1:
            self.res[0].left = self.res[0].right = self.res[0]
            return self.res[0]

        for i in range(len(self.res)):
            if i == 0:
                self.res[i].left = self.res[len(self.res) - 1]
                self.res[i].right = self.res[i + 1]
            elif i == len(self.res)-1:
                self.res[i].right = self.res[0]
                self.res[i].left = self.res[i - 1]
            else:
                self.res[i].left = self.res[i - 1]
                self.res[i].right = self.res[i + 1]

        return self.res[0]

if __name__ == "__main__":
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)
    node5 = Node(5)
    node4.left = node2
    node4.right = node5
    node2.left = node1
    node2.right = node3

    s = Solution()
    s.treeToDoublyList(node4)
