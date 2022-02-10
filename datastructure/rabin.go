package datastructure

const primeRK = 26
const mod uint64 = 1 << 32

func check(s string, n int) int {
	// Rabin-Karp search
	table := map[uint64]int{}
	var pow, sq uint64 = 1, primeRK
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
			pow %= mod
		}
		sq = sq * sq
		sq %= mod
	}
	var h uint64
	for i := 0; i < n; i++ {
		h = ((h*primeRK)%mod + uint64(s[i])) % mod
	}
	table[h]++
	for i := n; i < len(s); {
		h *= primeRK
		h %= mod
		h += uint64(s[i])
		h %= mod
		h -= (pow * uint64(s[i-n])) % mod
		h %= mod
		i++
		table[h]++
		if table[h] > 1 {
			return i - n
		}
	}
	return -1
}
