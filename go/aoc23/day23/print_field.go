package day23

import "fmt"

func printField(parseData []cell, w, h int) {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			idx := row*w + col
			c := parseData[idx]
			var char string
			switch c {
			case 0: // No connections
				char = "#"
			case dirExitTop:
				char = "╵"
			case dirExitBottom:
				char = "╷"
			case dirExitLeft:
				char = "╴"
			case dirExitRight:
				char = "╶"
			case dirExitTop | dirExitBottom:
				char = "│"
			case dirExitLeft | dirExitRight:
				char = "─"
			case dirExitTop | dirExitRight:
				char = "└"
			case dirExitTop | dirExitLeft:
				char = "┘"
			case dirExitBottom | dirExitRight:
				char = "┌"
			case dirExitBottom | dirExitLeft:
				char = "┐"
			case dirExitTop | dirExitBottom | dirExitRight:
				char = "├"
			case dirExitTop | dirExitBottom | dirExitLeft:
				char = "┤"
			case dirExitTop | dirExitLeft | dirExitRight:
				char = "┴"
			case dirExitBottom | dirExitLeft | dirExitRight:
				char = "┬"
			case dirExitAll:
				char = "┼"
			default:
				char = "?" // Should not happen
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}
