package main

import (
	"fmt"
	"net/http"
	"time"
)

//Result exported
type Result struct {
	userName string
	title    string
	likes    string
}

func (r Result) String() string {
	return fmt.Sprint(r.userName, " - ", r.title, " - ", r.likes, " claps")
}

func main() {
	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}
	var ini time.Time
	fmt.Println("Without go routine:")
	ini = time.Now()
	for _, url := range urlToProcess {
		r, err := http.Get(url)
		if err != nil {
			fmt.Println(http.StatusFound)
		}
		fmt.Println)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

}
