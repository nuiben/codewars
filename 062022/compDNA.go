package kata

func DNAStrand(dna string) string {
	comp := ""
	for _, i := range dna {
		if string(i) == "A" {
			comp += "T"
		} else if string(i) == "T" {
			comp += "A"
		} else if string(i) == "G" {
			comp += "C"
		} else {
			comp += "G"
		}
	}
	return comp
}
