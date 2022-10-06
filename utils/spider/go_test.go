package spider

import (
	"fmt"
	"testing"
)

func TestDownloader(t *testing.T) {
	downloader, _ := Downloader("https://tieba.baidu.com/f?kw=steam")
	fmt.Println(downloader.Html())
}
