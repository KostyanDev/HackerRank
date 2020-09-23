//https://www.hackerrank.com/challenges/designer-pdf-viewer/problem
func designerPdfViewer(h []int32, word string) int32 {

	var currentHeight, returnValue int32
	var maxHeight int32 = h[word[0] - 'a']

	for i := 1; i < len(word); i++ {
		currentHeight = h[word[i] - 'a']
		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}
	}
	returnValue = maxHeight * int32((len(word)))
	return returnValue
}
