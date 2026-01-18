package main

// The Bad: Uses a boolean flag to do two different things in one function.
// func RenderPage(content string, isMobile bool) {
// 	if isMobile {
// 		// logic for mobile layout
// 	} else {
// 		// logic for desktop layout
// 	}
// }

// The Good: Splits responsibilities into two specific, clear functions.
func RenderMobilePage(content string) {
	// ...
}

func RenderDesktopPage(content string) {
	// ...
}
