package secondlargestdigitinastring

func secondHighest(s string) int {
	fst := -1
	snd := -1
	for _, c := range s {
		d := int(c - '0')
		if d < 0 || d > 9 {
			continue
		}
		if d > snd {
			if d > fst {
				snd = fst
				fst = d
			} else if d != fst {
				snd = d
			}
		}
	}

	if snd == fst {
		return -1
	}

	return snd
}
