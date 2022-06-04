/* 

Read a Bhagavad Gita verse randomly

Author: HareKrsna16

What I learnt:
	http requests
	html parsing
	golang basics

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"bufio"
	"strings"
)

func chap_verse_gen() (int, int) {

	bg_verse_num := [18]int{46, 72, 43, 42, 29, 47, 30, 28, 34, 42, 55, 20, 35, 27, 20, 24, 28, 78}

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	chapter := rand.Intn(18)+1
	verse := rand.Intn(bg_verse_num[chapter-1])+1

	return chapter, verse

}

func content(url string) string {

	resp, err := http.Get(url)
	if err != nil {
        panic(err)
    }
	defer resp.Body.Close()

	html_code := ""

	if strings.Contains(resp.Status, "404") {
		return  "0"
	}

	sc := bufio.NewScanner(resp.Body)
	for i := 0; sc.Scan(); i++ {
		html_part := sc.Text()
		if strings.Contains(html_part, "<strong>"){
			html_code += sc.Text()
		}
	}

	if err := sc.Err(); err != nil {
        panic(err)
    }

	return strings.TrimSpace(html_code)

}

func cleanup(html_part string) string {
	return strings.Split(strings.Split(html_part, "<")[3], ">")[1]
}

func main() {

	chapter, verse := chap_verse_gen()

	url := fmt.Sprintf("https://vedabase.io/en/library/bg/%d/%d", chapter, verse)

	html_code := content(url)
	start_verse := verse

	if html_code == "0" {
		for html_code == "0" {
			start_verse -= 1
			url := fmt.Sprintf("https://vedabase.io/en/library/bg/%d/%d", chapter, start_verse)
			html_code = content(url)
		}
		html_code = "0"
		start_verse += 1
		end_verse := start_verse
		fmt.Println(start_verse)
		for html_code == "0" {
			end_verse += 1
			url := fmt.Sprintf("https://vedabase.io/en/library/bg/%d/%d-%d", chapter, start_verse, end_verse)
			html_code = content(url)
		}

		url = fmt.Sprintf("https://vedabase.io/en/library/bg/%d/%d-%d", chapter, start_verse, end_verse)
		fmt.Printf("Bg. %d.%d-%d\n\n", chapter, start_verse, end_verse)
	} else {
		fmt.Printf("Bg. %d.%d\n\n", chapter, verse)
	}
	
	fmt.Println(cleanup(content(url)), "\n")
	fmt.Println(url)

}