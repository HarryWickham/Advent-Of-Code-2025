package main

import (
	"fmt"
	"strconv"
	"strings"
)

const example_input string = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

const real_input string = "16100064-16192119,2117697596-2117933551,1-21,9999936269-10000072423,1770-2452,389429-427594,46633-66991,877764826-877930156,880869-991984,18943-26512,7216-9427,825-1162,581490-647864,2736-3909,39327886-39455605,430759-454012,1178-1741,219779-244138,77641-97923,1975994465-1976192503,3486612-3602532,277-378,418-690,74704280-74781349,3915-5717,665312-740273,69386294-69487574,2176846-2268755,26-45,372340114-372408052,7996502103-7996658803,7762107-7787125,48-64,4432420-4462711,130854-178173,87-115,244511-360206,69-86"

type RangePair struct {
	start int
	end   int
}

func main() {
	ranges := parse_input(real_input)
	invalid_ids := []int{}

	for _, r := range ranges {
		invalid_ids = append(invalid_ids, calculate_repeats(r)...)

	}

	var total int
	for _, id := range invalid_ids {
		total += id
	}
	fmt.Println("Total invalid IDs sum:", total)
}

func calculate_repeats(r RangePair) []int {
	invalid_ids := []int{}
	for i := r.start; i <= r.end; i++ {
		if calculate(i) > 1 {
			invalid_ids = append(invalid_ids, i)
		}
	}

	return invalid_ids
}

func calculate(number int) int {
	str_i := strconv.Itoa(number)

	str_len := len(str_i)

	for sections := 2; sections <= str_len; sections++ {
		parts, ok := divide_into_sections(str_i, sections)
		if ok && check_all_equal(parts) {
			return sections
		}
	}

	return 0
}

func divide_into_sections(str string, numSections int) ([]string, bool) {
	if numSections <= 0 {
		return nil, false
	}

	strLen := len(str)
	if strLen%numSections != 0 {
		return nil, false
	}

	sectionLen := strLen / numSections
	sections := make([]string, numSections)

	for i := range numSections {
		start := i * sectionLen
		end := start + sectionLen
		sections[i] = str[start:end]
	}

	return sections, true
}

func check_all_equal(sections []string) bool {
	if len(sections) == 0 {
		return false
	}
	reference := sections[0]
	for _, sec := range sections[1:] {
		if sec != reference {
			return false
		}
	}
	return true
}

func parse_input(input string) []RangePair {
	ranges := strings.Split(input, ",")
	parsedRanges := []RangePair{}

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		if len(bounds) == 2 {

			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				fmt.Println("Error converting range", r)
				continue
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				fmt.Println("Error converting range", r)
				continue
			}
			parsedRanges = append(parsedRanges, RangePair{start, end})
		}
	}
	return parsedRanges
}
