package decryptstringfromalphabettointegermapping

func freqAlphabets(s string) string {
	out := []byte{}
	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			out = append(out, 'j'+(s[i]-'1')*10+s[i+1]-'0')
			i += 3
		} else {
			out = append(out, 'a'+s[i]-'1')
			i++
		}
	}

	return string(out)
}
