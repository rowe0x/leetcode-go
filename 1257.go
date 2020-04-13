package leetcode

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	/*
	1) child -> parent map
	2) parent chain
	 */
	m := make(map[string]string)
	for _, region := range regions {
		for _, r := range region[1:] {
			m[r] = region[0]
		}
	}
	region1Parent := make(map[string]struct{})
	region := region1
	for {
		if p, ok := m[region]; ok {
			region1Parent[p] = struct{}{}
			region = p
		} else {
			break
		}
	}
	region = region2
	for {
		if _, ok := region1Parent[region]; ok {
			return region
		}
		if p, ok := m[region]; ok {
			region = p
		} else {
			break
		}
	}
	return ""
}