package p921

func minAddToMakeValid(S string) int {
	var ans int
	var cnt int

	n := len(S)
	for i := 0; i < n; i++ {
		if S[i] == '(' {
			cnt++
		} else if S[i] == ')' {
			cnt--
		}
		if cnt < 0 {
			ans++
			cnt = 0
		}
	}

	return ans + cnt
}

func minAddToMakeValid1(S string) int {
	n := len(S)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	// dp[i] = j, means the far pos that make [i..j] balance

	for i := 0; i < n; i++ {
		if S[i] == ')' {
			continue
		}
		var level int
		for j := i; j < n; j++ {
			if S[j] == '(' {
				level++
			} else if S[j] == ')' {
				level--
			}
			if level == 0 {
				dp[i] = j
				break
			}
		}
		// skip the balance pairs
		if dp[i] > i {
			i = dp[i]
		}
	}

	var ans int
	var pos int
	for pos < n {
		if dp[pos] < 0 {
			ans++
			pos++
		} else {
			pos = dp[pos] + 1
		}
	}
	return ans
}
