package byteUtil


func ReadLines(bytes []byte) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; bytes[i] != 0 ; i++ {
		if bytes[i] == 10 {
			// println(string(bytes[start:i]))
			strSlice = append(strSlice, string(bytes[start:i]))
			start = i + 1
		}
	}
	return strSlice
}