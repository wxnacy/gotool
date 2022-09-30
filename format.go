package gotool

import (
	"fmt"
	"math"
)

// format size to string
func FormatSize(size int64) string {
	suffixs := []string{"B", "K", "M", "G", "T"}
	scale := float64(1024)
	for i, suffix := range suffixs {
		suffix_size := math.Pow(scale, float64(i+1))
		if size < int64(suffix_size) {
			beichu := math.Pow(scale, float64(i))
			if beichu == 0 {
				beichu = 1
			}
			res := float64(size) / beichu
			if float64(int64(res)) == res {
				return fmt.Sprintf("%d%s", int64(res), suffix)
			} else {
				return fmt.Sprintf("%.2f%s", res, suffix)
			}
		}
	}
	return fmt.Sprintf("%dB", size)
}
