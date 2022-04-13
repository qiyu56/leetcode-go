package problems

// 差分数组

type Difference struct {
	Diff []int
}

func New(nums []int) *Difference {
	d := &Difference{}
	diff := make([]int, len(nums))
	// 根据原数组构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	d.Diff = diff
	return d
}

func (d *Difference) Increment(i, j, val int) {
	d.Diff[i] += val
	if j+1 < len(d.Diff) {
		d.Diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	res := make([]int, len(d.Diff))
	// 根据差分数组构造结果数组
	res[0] = d.Diff[0]
	for i := 1; i < len(d.Diff); i++ {
		res[i] = res[i-1] + d.Diff[i]
	}
	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	df := New(nums)
	for _, e := range bookings {
		i := e[0] - 1
		j := e[1] - 1
		val := e[2]
		df.Increment(i, j, val)
	}
	return df.Result()
}

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001) // 1001个车站
	df := New(nums)
	for _, trip := range trips {
		from := trip[1]
		to := trip[2] - 1
		num := trip[0]
		df.Increment(from, to, num)
	}

	res := df.Result()
	for _, e := range res {
		if e > capacity {
			return false
		}
	}
	return true
}
