package util

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")

	for tmp := list; tmp != nil; tmp = tmp.Next {
		buf.WriteString(strconv.Itoa(tmp.Val))
		buf.WriteString(",")
	}

	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteString("]")
	return buf.String()
}

func ParseAsList(str string) *ListNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	nodes := strings.Split(str, ",")

	if len(nodes) == 0 {
		return nil
	}

	head := &ListNode{Val: parseNum(nodes[0])}
	prev := head
	for i := 1; i < len(nodes); i++ {
		cur := &ListNode{Val: parseNum(nodes[i])}
		prev.Next = cur
		prev = cur
	}

	return head
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%v,%v", tn.Val, tn.Left, tn.Right)
}

func ParseAsTree(str string) *TreeNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)

	var currLevel []*TreeNode
	var nextLevel []*TreeNode
	nodes := strings.Split(str, ",")

	root := parseNode(nodes[0])
	currLevel = append(currLevel, root)

	pt := 0
	ct := 0
	for i := 1; i < len(nodes); i++ {
		parent := currLevel[pt]
		node := parseNode(nodes[i])
		if node != nil {
			nextLevel = append(nextLevel, node)
		}
		if ct == 0 {
			parent.Left = node
			ct++
		} else {
			parent.Right = node
			pt++
			ct = 0
		}
		if pt == len(currLevel) {
			currLevel = nextLevel
			nextLevel = make([]*TreeNode, 0)
			pt = 0
		}
	}
	return root
}

func parseNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	return &TreeNode{parseNum(str), nil, nil}
}
func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func SprintTree(root *TreeNode) string {
	if root == nil {
		return "[null]"
	}
	var buf bytes.Buffer

	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node == nil {
				buf.WriteString("null,")
			} else {
				buf.WriteString(strconv.Itoa(node.Val))
				buf.WriteString(",")
				nextLevel = append(nextLevel, node.Left)
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	return buf.String()
}
