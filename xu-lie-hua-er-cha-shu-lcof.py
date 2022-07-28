# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
import collections


class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if root is None:
            return ""
        res = []
        queue = collections.deque()
        queue.append(root)
        while queue:
            node = queue.popleft()
            if node is not None:
                res.append(str(node.val))
                queue.append(node.left)
                queue.append(node.right)
            else:
                res.append("#")
        return ",".join(res)


    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        if data == "":
            return None

        nodes = data.split(",")

        i = 0
        queue = collections.deque()
        root = TreeNode(int(nodes[i]))
        i += 1
        queue.append(root)
        while queue:
            node = queue.popleft()
            if nodes[i] != "#":
                node.left = TreeNode(int(nodes[i]))
                queue.append(node.left)
            i += 1
            if nodes[i] != "#":
                node.right = TreeNode(int(nodes[i]))
                queue.append(node.right)
            i += 1

        return root

# Your Codec object will be instantiated and called as such:
codec = Codec()
codec.deserialize(codec.serialize(None))