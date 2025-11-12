package day23

func countDegrees(parsedData []cell, idx, w, h int) (entries, exits int) {
	x := idx % w
	y := idx / w
	mask := parsedData[idx]
	if mask&dirExitTop == 0 {
		exits++
	} else if y > 0 && parsedData[idx-w]&dirExitBottom == 0 {
		entries++
	}
	if mask&dirExitBottom == 0 {
		exits++
	} else if y < h-1 && parsedData[idx+w]&dirExitTop == 0 {
		entries++
	}
	if mask&dirExitLeft == 0 {
		exits++
	} else if x > 0 && parsedData[idx-1]&dirExitRight == 0 {
		entries++
	}
	if mask&dirExitRight == 0 {
		exits++
	} else if x < w-1 && parsedData[idx+1]&dirExitLeft == 0 {
		entries++
	}
	return entries, exits
}
