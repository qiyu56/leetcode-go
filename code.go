package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ret := [][]string{{""}, {"()"}}
	for i := 2; i <= n; i++ {
		var tmp []string
		for j := 0; j < i; j++ {
			for _, s1 := range ret[j] {
				for _, s2 := range ret[i-j-1] {
					s := "(" + s1 + ")" + s2
					tmp = append(tmp, s)
				}
			}
		}
		ret = append(ret, tmp)
	}
	return ret[n]
}

type Interval struct {
	Start, End int
}

func merge(intervals []*Interval) []*Interval {
	// write your code here
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	var ret []*Interval
	tmp := intervals[0]
	for i, item := range intervals {
		if i == 0 {
			continue
		}
		if item.Start >= tmp.Start && item.End <= tmp.End {
			continue
		} else if item.Start <= tmp.End && item.End >= tmp.End {
			tmp.End = item.End
		} else {
			ret = append(ret, tmp)
			tmp = item
		}
	}
	ret = append(ret, tmp)
	return ret
}

func firstUniqChar(s string) int {
	// write your code here
	countMap := make(map[int32]int)
	for _, c := range s {
		countMap[c] += 1
	}
	for i, c := range s {
		if countMap[c] == 1 {
			return i
		}
	}
	return -1
}

func isNStraightHand(hand []int, groupSize int) bool {
	counts := make(map[int]int)
	h := &IntHeap{}
	for _, val := range hand {
		counts[val] += 1
		heap.Push(h, val)
	}
	for h.Len() > 0 {
		t := heap.Pop(h).(int)
		if counts[t] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			cnt := counts[t+i]
			if cnt == 0 {
				return false
			}
			counts[t+i] = cnt - 1
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: 0, Next: nil}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum / 10
		newNode := ListNode{Val: sum % 10, Next: nil}
		cur.Next = &newNode
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		newNode := ListNode{Val: carry, Next: nil}
		cur.Next = &newNode
	}
	return pre.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	head := list1
	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	pre := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else if list2 != nil {
		pre.Next = list2
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := head
	right := head.Next
	node := swapPairs(right.Next)
	right.Next = left
	left.Next = node
	return right
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	var lastNode *ListNode
	for node := head; node != nil; node = node.Next {
		count++
		lastNode = node
	}
	n := k % count
	if n == 0 {
		return head
	}
	fastNode := head
	slowNode := head
	for i := 0; i < n; i++ {
		fastNode = fastNode.Next
	}
	for fastNode.Next != nil {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	lastNode.Next = slowNode
	head = slowNode.Next
	slowNode.Next = nil
	return head
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	cur := head
// 	node := cur.Next
// 	for node != nil {
// 		if cur.Val != node.Val {
// 			cur.Next = node
// 			cur = node
// 			node = node.Next
// 		} else {
// 			node = node.Next
// 		}
// 	}
// 	cur.Next = node
// 	return head
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &ListNode{Val: 0, Next: nil}
	cur := h
	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val != node.Val {
			cur.Next = node
			cur = cur.Next
			node = node.Next
		} else {
			for node.Next != nil && node.Val == node.Next.Val {
				node = node.Next
			}
			node = node.Next
		}
	}
	cur.Next = node
	return h.Next
}

func pancakeSort(arr []int) []int {
	n := len(arr)
	idxs := make([]int, n+1)
	for i := 0; i < n; i++ {
		idxs[arr[i]] = i
	}
	var ans []int
	for i := n; i >= 1; i-- {
		idx := idxs[i]
		// 第i大的数就在第i个位置上
		if idx == i-1 {
			continue
		}
		// 不在第一个，翻转到第一个
		if idx != 0 {
			ans = append(ans, idx+1)
			reverse(arr, 0, idx, idxs)
		}
		// 再将第一个翻转到第i个
		ans = append(ans, i)
		reverse(arr, 0, i-1, idxs)
	}
	return ans
}

func reverse(arr []int, i int, j int, idxs []int) {
	for i < j {
		idxs[arr[i]] = j
		idxs[arr[j]] = i
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fs, fb, s, b *ListNode
	cur := head
	for cur != nil {
		if cur.Val < x {
			if s != nil {
				s.Next = cur
				s = s.Next
			} else {
				s = cur
				fs = s
			}
		} else {
			if b != nil {
				b.Next = cur
				b = b.Next
			} else {
				b = cur
				fb = b
			}
		}
		cur = cur.Next
	}
	if b != nil {
		b.Next = nil
	}
	if s != nil {
		s.Next = fb
	}
	return fs
}

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	idx := 0
	for idx < n-1 {
		if bits[idx] == 0 {
			idx++
		} else {
			idx += 2
		}
	}
	return idx == n-1
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	//  Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
}

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	mask := 1 << n // 2的n次方种子集
	max, ans := 0, 0
	for s := 0; s < mask; s++ {
		cur := 0
		for i := 0; i < n; i++ {
			if (s>>i)&1 == 1 { // 该子集中是否存在nums[i]
				cur |= nums[i]
			}
		}
		if cur > max {
			max = cur
			ans = 1
		} else if cur == max {
			ans += 1
		}
	}
	return ans
}

func countMaxOrSubsets1(nums []int) int {
	n := len(nums)
	max, ans := 0, 0
	var dfs func(int, int)
	dfs = func(idx int, val int) {
		if idx == n {
			if val > max {
				max = val
				ans = 1
			} else if val == max {
				ans += 1
			}
			return
		}
		dfs(idx+1, val)
		dfs(idx+1, val|nums[idx])
	}
	dfs(0, 0)
	return ans
}

// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	total := 0
	for i := 1; i < n-1; i++ {
		lMax := IntSliceMax(height[:i])
		rMax := IntSliceMax(height[i+1 : n])
		if lMax > rMax && rMax > height[i] {
			total += rMax - height[i]
		} else if rMax >= lMax && lMax > height[i] {
			total += lMax - height[i]
		}
	}
	return total
}

func trapOpt(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = IntMax(height[i], leftMax[i-1])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = IntMax(height[i], rightMax[i+1])
	}
	total := 0
	for i := 1; i < n-1; i++ {
		total += IntMax(0, IntMin(leftMax[i], rightMax[i])-height[i])
	}
	return total
}

func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	} else if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{-1, nil}
	p := dummy
	h := &NodeHeap{}
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA, countB := calcCount(headA), calcCount(headB)
	shorter, longer := headA, headB
	if countA > countB {
		longer, shorter = headA, headB
	}
	diff := Abs(countA - countB)
	for i := 0; i < diff; i++ {
		longer = longer.Next
	}
	for shorter != nil && longer != nil {
		if shorter == longer {
			return shorter
		}
		shorter, longer = shorter.Next, longer.Next
	}
	return nil
}

func calcCount(head *ListNode) (count int) {
	for head != nil {
		head = head.Next
		count++
	}
	return
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var node *ListNode

// 翻转链表前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		node = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = node
	return last
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	n := right - left + 1
	if left == 1 {
		return reverseN(head, n)
	}
	p := head
	for i := 1; i < left-1; i++ {
		p = p.Next
	}
	// p指向left的前一个节点
	p.Next = reverseN(p.Next, n)
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseList2(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	cur, next := head, head.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(n1, n2 *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := n1, n1
	for cur != n2 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type FibFunc func(uint64, uint64) uint64

func buildLazyIntFibIterator(fibFunc FibFunc, initState1, initState2 uint64) func() uint64 {
	retValChan := make(chan uint64)
	loopFunc := func() {
		state1, state2 := initState1, initState2
		var retVal uint64
		for {
			retVal = fibFunc(state1, state2)
			retValChan <- state1
			state1 = state2
			state2 = retVal
		}
	}
	retFunc := func() uint64 {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func iteratorFibs() {
	fibFunc := func(a, b uint64) uint64 {
		return a + b
	}
	fibs := buildLazyIntFibIterator(fibFunc, 0, 1)
	for i := 0; i < 50; i++ {
		fmt.Printf("%dth fib: %v\n", i, fibs())
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in // 3
		if i%prime != 0 {
			out <- i
		}
	}
}

func sieve() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch // 2,3
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// 2,3,4,5,6,7,8,9,10,11
// 3,5,7,9,11
// 5,7,11
// 7,11
// 11

var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	// 相当于把链表节点压栈，来翻转链表，翻转后的和未反转的逐一比较
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}

func isPalindromeOpt(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { // 奇数个节点
		slow = slow.Next
	}

	left := head
	right := reverseListNodes(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseListNodes(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//type NumArray struct {
//	arr    []int
//	preSum []int
//}
//
//func Constructor(nums []int) NumArray {
//	preSum := make([]int, len(nums)+1)
//	for i := 1; i < len(preSum); i++ {
//		preSum[i] = preSum[i-1] + nums[i]
//	}
//	return NumArray{arr: nums, preSum: preSum}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	return this.preSum[right+1] - this.preSum[left]
//}

type NumMatrix struct {
	matrix [][]int
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	h, w := len(matrix), len(matrix[0])
	preSum := make([][]int, h+1)
	for i := 0; i < len(preSum); i++ {
		preSum[i] = make([]int, w+1)
	}
	for i := 1; i < len(preSum); i++ {
		for j := 1; j < len(preSum[0]); j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{matrix: matrix, preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	var res int
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				res++
			}
		}
	}
	return res
}

func subarraySumOpt(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int)
	preSum[0] = 1
	res, sum0_i := 0, 0
	for i := 0; i < n; i++ {
		sum0_i += nums[i]
		sum0_j := sum0_i - k // sum0_i - sum0_j = k
		if val, ok := preSum[sum0_j]; ok {
			res += val
		}
		if val, ok := preSum[sum0_i]; ok {
			preSum[sum0_i] = val + 1
		} else {
			preSum[sum0_i] = 1
		}
	}
	return res
}
