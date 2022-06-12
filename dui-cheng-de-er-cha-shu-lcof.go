package main

func check(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}

	if A == nil || B == nil {
		return false
	}

	return A.Val == B.Val && check(A.Left, B.Right) && check(A.Right, B.Left)
}


func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
