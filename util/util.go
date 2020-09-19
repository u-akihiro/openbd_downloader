package util

// ChunkSize 分割するサイズ
const ChunkSize = 10000

// Slice 渡されたスライスを指定された個数に分割する
func Slice(slice []string, chunkSize int) [][]string {
	chunkTotal := len(slice) / chunkSize

	if (len(slice) % chunkSize) != 0 {
		chunkTotal++
	}

	var sliced [][]string

	var start int
	var end int
	for i := 0; i < chunkTotal; i++ {
		start = i * chunkSize
		end = start + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		sliced = append(sliced, slice[start:end])
	}

	return sliced
}
