package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Println()
	if runtime.GOOS == "windows" {
		fmt.Println("¯\\_(ツ)_//¯  (ง^ᗜ^)ง  ʕ•̫͡•ʕ•̫͡•ʔ•̫͡•ʔ ")

		fmt.Println()
		fmt.Println("heyyy!  :) Im Nitish Sai Kommaraju Hope U having Great Day ༼ つ ◕_◕ ༽つ━☆ﾟ.*･｡ﾟ ")
		fmt.Println()
		fmt.Println("{\\__//} \n (●_●)\n ( > )  Windows os Detected with", runtime.GOARCH, time.Now().Format("Monday, 2 Jan 2006"))

	}
	fmt.Println()

	if runtime.GOOS == "Linux" {
		fmt.Println("¯\\_(ツ)_//¯  (ง^ᗜ^)ง  ʕ•̫͡•ʕ•̫͡•ʔ•̫͡•ʔ ")

		fmt.Println()
		fmt.Println("heyyy!  :) Im Nitish Sai Kommaraju Hope U having Great Day ༼ つ ◕_◕ ༽つ━☆ﾟ.*･｡ﾟ ")
		fmt.Println()
		fmt.Println("{\\__//} \n (●_●)\n ( > )  Linux os Detected with", runtime.GOARCH, time.Now().Format("Monday, 2 Jan 2006"))

	}

	for true {

		var URL string

		fmt.Println("Please Enter the Url in format { https//url.com/***.file_extension }")

		fmt.Scanln(&URL)
		z := strings.Split(URL, ":")
		test := z[0]

		//split the string to get the  last name
		if URL != "" && test == "https" || test == "http" {

			link := strings.Split(URL, "/")

			Name := link[len(link)-1]

			fmt.Println("Downloading", URL, "to File Name", Name)
			output, err := os.Create(Name)
			if err != nil {
				fmt.Println("ooops Thre is an Error While Downloading", Name, "-", err)

			}
			defer output.Close()

			response, err := http.Get(URL)
			if err != nil {
				fmt.Println("ooops Thre is an Error While Downloading", URL, "-", err)

			}
			defer response.Body.Close()

			n, err := io.Copy(output, response.Body)

			if err != nil {
				fmt.Println("ooops Thre is an Error While Downloading", URL, "-", err)

			}

			fmt.Println(n, "bytes downloaded.", time.Now().Format("Monday, 2 Jan 2006"))

		} else {
			fmt.Println("Enter Url in https format or http format with an file extention")

		}
		fmt.Println()
	}

}
