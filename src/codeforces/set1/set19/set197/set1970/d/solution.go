package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	play := func(f func(int) []int) {
		q := readNum(reader)

		for i := 0; i < q; i++ {
			p := readNum(reader)
			res := f(p)
			fmt.Printf("%d %d\n", res[0], res[1])
		}
	}

	solve(n, func(sets []string) {
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteString(sets[i])
			buf.WriteByte('\n')
		}
		fmt.Print(buf.String())
	}, play)

}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

var p = []int{1, 2, 6, 8, 12, 16, 20, 24, 30, 36, 40, 44, 48, 56, 60, 65, 72, 81, 88, 100, 107, 116, 131, 140, 144, 156, 160, 162, 171, 176, 192, 200, 210, 212, 221, 225, 242, 248, 254, 261, 276, 280, 292, 300, 309, 336, 351, 357, 362, 371, 378, 380, 396, 410, 417, 456, 465, 477, 501, 507, 522, 540, 546, 550, 552, 556, 597, 611, 651, 657, 672, 696, 707, 722, 732, 756, 760, 771, 786, 800, 852, 857, 861, 875, 877, 885, 920, 945, 950, 972, 976, 980, 996, 1002, 1027, 1037, 1041, 1060, 1080, 1092, 1124, 1130, 1152, 1173, 1185, 1200, 1212, 1227, 1236, 1260, 1271, 1275, 1277, 1281, 1285, 1338, 1344, 1365, 1380, 1386, 1390, 1407, 1431, 1437, 1452, 1485, 1501, 1520, 1530, 1536, 1542, 1561, 1565, 1610, 1617, 1635, 1653, 1704, 1712, 1722, 1731, 1736, 1821, 1850, 1872, 1887, 1900, 1905, 1911, 1917, 1932, 1940, 1948, 1981, 1992, 1997, 2016, 2025, 2037, 2096, 2100, 2104, 2110, 2120, 2135, 2178, 2205, 2216, 2220, 2227, 2241, 2332, 2337, 2346, 2370, 2380, 2400, 2432, 2445, 2452, 2467, 2472, 2488, 2516, 2520, 2530, 2554, 2578, 2592, 2600, 2632, 2656, 2688, 2706, 2710, 2720, 2730, 2740, 2795, 2802, 2811, 2824, 2841, 2856, 2892, 2901, 2915, 2955, 2961, 2976, 2984, 2997, 3021, 3060, 3082, 3096, 3098, 3104, 3112, 3117, 3140, 3165, 3173, 3186, 3192, 3201, 3216, 3252, 3257, 3306, 3317, 3321, 3340, 3360, 3370, 3385, 3412, 3462, 3480, 3492, 3516, 3531, 3550, 3577, 3591, 3600, 3612, 3621, 3642, 3675, 3700, 3717, 3732, 3737, 3822, 3851, 3861, 3880, 3896, 3930, 3941, 3945, 3956, 3961, 4005, 4020, 4026, 4032, 4050, 4103, 4125, 4131, 4137, 4176, 4196, 4224, 4240, 4251, 4257, 4280, 4347, 4356, 4371, 4396, 4404, 4424, 4440, 4491, 4531, 4545, 4560, 4567, 4585, 4602, 4620, 4650, 4720, 4722, 4785, 4797, 4820, 4875, 4901, 4917, 4929, 4944, 4972, 5005, 5007, 5037, 5076, 5085, 5096, 5100, 5112, 5145, 5226, 5232, 5292, 5296, 5307, 5340, 5365, 5384, 5397, 5411, 5435, 5440, 5445, 5456, 5491, 5501, 5505, 5517, 5571, 5580, 5636, 5645, 5667, 5690, 5700, 5748, 5799, 5826, 5852, 5860, 5865, 5877, 5922, 5937, 6000, 6007, 6036, 6072, 6085, 6120, 6124, 6132, 6156, 6196, 6216, 6315, 6330, 6357, 6372, 6425, 6427, 6443, 6450, 6457, 6471, 6481, 6501, 6580, 6597, 6657, 6700, 6716, 6736, 6780, 6796, 6812, 6840, 6861, 6882, 6897, 6915, 6917, 6947, 6960, 6981, 6996, 7010, 7032, 7056, 7076, 7080, 7130, 7140, 7147, 7165, 7182, 7212, 7228, 7242, 7261, 7321, 7331, 7350, 7380, 7382, 7392, 7413, 7491, 7497, 7506, 7515, 7521, 7527, 7536, 7557, 7576, 7587, 7605, 7641, 7690, 7692, 7700, 7722, 7752, 7816, 7821, 7832, 7856, 7890, 7917, 7932, 7947, 7976, 8037, 8052, 8060, 8091, 8100, 8145, 8156, 8196, 8220, 8236, 8246, 8284, 8313, 8330, 8336, 8346, 8352, 8385, 8421, 8432, 8460, 8484, 8514, 8541, 8565, 8576, 8592, 8652, 8676, 8712, 8725, 8752, 8781, 8842, 8892, 8960, 8964, 9012, 9027, 9032, 9036, 9057, 9066, 9090, 9132, 9146, 9170, 9180, 9186, 9192, 9225, 9234, 9240, 9276, 9284, 9321, 9357, 9402, 9412, 9435, 9440, 9464, 9472, 9501, 9570, 9585, 9625, 9660, 9675, 9696, 9720, 9741, 9747, 9765, 9780, 9816, 9849, 9857, 9880, 9885, 9900, 9906, 9937, 9972, 9984, 9990, 9996, 10000, 10005, 10010, 10040, 10067, 10101, 10117, 10161, 10192, 10200, 10252, 10272, 10290, 10332, 10336, 10340, 10356, 10416, 10452, 10531, 10557, 10565, 10605, 10650, 10665, 10685, 10707, 10747, 10760, 10836, 10851, 10860, 10909, 10956, 11001, 11056, 11066, 11112, 11125, 11137, 11160, 11172, 11180, 11211, 11272, 11277, 11300, 11316, 11330, 11377, 11400, 11421, 11427, 11442, 11472, 11481, 11487, 11496, 11547, 11556, 11586, 11596, 11620, 11676, 11697, 11704, 11720, 11767, 11805, 11811, 11832, 11865, 11937, 11991, 12041, 12072, 12084, 12105, 12114, 12128, 12156, 12162, 12195, 12200, 12225, 12240, 12261, 12300, 12345, 12400, 12420, 12432, 12456, 12471, 12485, 12492, 12512, 12525, 12537, 12546, 12552, 12572, 12600, 12645, 12651, 12656, 12685, 12720, 12756, 12761, 12807, 12827, 12840, 12882, 12891, 12897, 12960, 12981, 13002, 13041, 13056, 13092, 13101, 13125, 13140, 13145, 13181, 13220, 13227, 13257, 13308, 13344, 13356, 13370, 13380, 13387, 13416, 13420, 13436, 13517, 13521, 13612, 13680, 13713, 13748, 13827, 13875, 13920, 13932, 13940, 13962, 13996, 14050, 14080, 14085, 14106, 14112, 14141, 14147, 14157, 14186, 14205, 14224, 14226, 14241, 14256, 14260, 14322, 14345, 14372, 14385, 14399, 14421, 14436, 14466, 14504, 14526, 14561, 14592, 14596, 14616, 14620, 14682, 14690, 14697, 14721, 14736, 14820, 14865, 14886, 14890, 14896, 14900, 14905, 15012, 15041, 15067, 15072, 15080, 15096, 15138, 15186, 15226, 15240, 15267, 15276, 15296, 15330, 15360, 15372, 15396, 15444, 15472, 15627, 15642, 15677, 15681, 15732, 15752, 15756, 15764, 15785, 15796, 15800, 15827, 15844, 15852, 15857, 15864, 15936, 16002, 16010, 16032, 16051, 16060, 16116, 16137, 16160, 16177, 16220, 16227, 16320, 16332, 16362, 16367, 16380, 16397, 16411, 16416, 16485, 16492, 16503, 16545, 16572, 16580, 16592, 16626, 16634, 16647, 16656, 16660, 16732, 16802, 16841, 16856, 16920, 17020, 17025, 17040, 17052, 17076, 17160, 17252, 17256, 17280, 17292, 17320, 17325, 17347, 17376, 17385, 17391, 17397, 17437, 17461, 17481, 17517, 17562, 17607, 17682, 17696, 17725, 17732, 17805, 17817, 17857, 17880, 17901, 17941, 18012, 18020, 18085, 18090, 18096, 18132, 18180, 18216, 18252, 18272, 18285, 18300, 18312, 18345, 18396, 18404, 18408, 18440, 18476, 18480, 18488, 18501, 18552, 18585, 18595, 18636, 18690, 18741, 18752, 18792, 18837, 18915, 18932, 18957, 19035, 19057, 19077, 19100, 19205, 19210, 19236, 19266, 19300, 19320, 19377, 19440, 19476, 19482, 19516, 19540, 19581, 19605, 19656, 19747, 19770, 19800, 19812, 19840, 19856, 19860, 20010, 20076, 20121, 20151, 20181, 20196, 20202, 20281, 20322, 20336, 20352, 20475, 20485, 20532, 20541, 20605, 20610, 20625, 20640, 20661, 20727, 20764, 20787, 20850, 20856, 20871, 20917, 20946, 20972, 21000, 21029, 21056, 21060, 21091, 21177, 21188, 21192, 21210, 21225, 21276, 21312, 21324, 21380, 21465, 21476, 21490, 21505, 21561, 21576, 21597, 21660, 21672, 21732, 21765, 21770, 21781, 21816, 21856, 22002, 22092, 22107, 22297, 22341, 22392, 22396, 22500, 22512, 22572, 22611, 22620, 22640, 22717, 22737, 22747, 22785, 22810, 22842, 22936, 22941, 22962, 22971, 22991, 23005, 23025, 23037, 23052, 23061, 23157, 23162, 23172, 23184, 23226, 23236, 23284, 23296, 23313, 23332, 23370, 23377, 23426, 23430, 23485, 23556, 23565, 23601, 23622, 23640, 23653, 23661, 23676, 23740, 23751, 23760, 23766, 23772}

func solve(n int, initilaizer func([]string), play func(func(int) []int)) {
	sets := make([]string, n)
	pref := "XO"
	prev := 0
	for i := 0; i < n; i++ {
		tmp := pref
		for j := prev; j < p[i]; j++ {
			tmp += "X"
		}
		sets[i] = tmp
		pref = tmp
		prev = p[i]
	}

	calc := func(x, y int) int {
		var ans int

		ans += max(x+1, y)
		// 包含一个o
		// 1/x, x/y-1
		ans += (1+1)*(x+2) + (x+2)*(y+1) - min(1+1, x+2)*min(x+2, y+1)
		// // 左1右x+1,左x+1右y
		// 包含两个o
		ans += (1 + 1) * (y + 1)

		return ans
	}

	ans := make(map[int]Pair)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans[calc(p[i], p[j])] = Pair{i + 1, j + 1}
		}
	}

	initilaizer(sets)

	play(func(power int) []int {
		res := ans[power]
		return []int{res.first, res.second}
	})
}

type Pair struct {
	first  int
	second int
}