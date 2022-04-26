package range_parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func Parse(size int, header string) ([]*Range, error) {
	index := strings.Index(header, "=")

	if index == -1 {
		return nil, fmt.Errorf("invalid range header")
	}

	arr := strings.Split(header[index+1:], ",")
	ranges := make([]*Range, 0, len(arr))

	for _, value := range arr {
		r := strings.Split(value, "-")
		start, startErr := strconv.Atoi(r[0])
		end, endErr := strconv.Atoi(r[1])

		if startErr != nil && endErr != nil {
			continue
		}

		// -nnn and nnn-
		if startErr != nil {
			start = size - end
			end = size - 1
		} else if endErr != nil {
			end = size - 1
		}

		if end >= size {
			end = size - 1
		}

		if start > end || start < 0 {
			continue
		}

		ranges = append(ranges, &Range{
			Start: start,
			End:   end,
		})
	}

	if len(ranges) == 0 {
		return nil, fmt.Errorf("unsatisifiable range header")
	}

	return ranges, nil
}
