package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"

	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	newUrl := strings.Replace(url, "http://", "https://", 1)
	client := http.Client{}
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		fmt.Printf("wrong http request: %s", err.Error())
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	cookie1 := "ec=oD7RMtJa-1601675640868-31f934725df822056840881; FSSBBIl1UgzbN7NO=5k.cmVaqduTc0TmSTrybb2x2yqIEd38IJfsc7zpQiLFL.h9zkUfbbsTsXaT2bdaCe64NrkAVsy6thg6awpe1iTa; sid=C36ufrYa5xi0jFnXza08; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601675713,1601865200; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1602046505; FSSBBIl1UgzbN7NP=5U6TttYee1PGqqqmTqKbeyAmHTsQqSfCaVx.IAPQmca_2zn4nyI_yLCvcR9_6dzMGrnOpV3U8lowFedud_IN_O6FSu0A7qoQxHlhKFn45OQwAA5H4fY77wUpBZ8JQdYUTw0B2bwIgtrsx5UhP6ckXJ1EhW_gYR6cSqCRv_EVkJSP6fg0oKaoPKHFVaub4._w1qypSUduvMrgVDi134olj.wYAmTMN7lj5cA4Z2cGC5urtPQCdwx52KWYxDy3oBaQWUlY38K.UgIgPCT0f0ymIpi; _efmdata=gpBuhgj%2BzwwMv8FDm38ry9QE0TR2L1iKRwSkT%2FWw6ofnbqw%2FKO6AmcgpLURa1hp03Zh0Qh3n8raEk2mX6C%2BH83riVvQgKrmPHitCN2gVfMY%3D; _exid=DctzMeTN0GQyS3QulcxAn06%2Bs3y0xpmLR5ESPNrBapmTLs%2BTYQRZhBp5vv07wNj6NPgYvq4YYw20pgjzCLSFvA%3D%3D"
	req.Header.Add("cookie", cookie1)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Print("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
