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
	cookie1 := "ec=oD7RMtJa-1601675640868-31f934725df822056840881; FSSBBIl1UgzbN7NO=5k.cmVaqduTc0TmSTrybb2x2yqIEd38IJfsc7zpQiLFL.h9zkUfbbsTsXaT2bdaCe64NrkAVsy6thg6awpe1iTa; sid=C36ufrYa5xi0jFnXza08; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601675713,1601865200; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601870125; FSSBBIl1UgzbN7NP=5U6jRhYe_b1aqqqm0eXRf3GsApgM5iGhYzwCIvbkmieRE0cDSpyQHXagVoUKgnFfr.1mDQ7uCCpEQJ5HDzcDAS_lA54QYFgGAofFsb19xbtOBhZXnSu6JW2nG7996BqMEQwrySXZI8O3unPApRaf.vdyLNPDboQCbUoQjLoJ4.RdUpj2q1K40CQk4NWysJ8ASVm4su1zy.7UU0WV1WDS.HFWmX.3kylqDYj2UzSDHjeXUcHyRXpiqNDbpSQx4BTK93; _efmdata=gpBuhgj%2BzwwMv8FDm38ry9QE0TR2L1iKRwSkT%2FWw6ofnbqw%2FKO6AmcgpLURa1hp0QVXV%2FXSjXJshrmaNhTEcOJxO8MxUlNaiBH1v2a644kk%3D; _exid=sf1N%2FJNfd988sww0qtEL1IPxvTSWNT7jrQEmtfWYktMEc%2BW1qvQmTAX4bkrUV%2BsbAb%2FVFoDYfMpLTd3o2ug4tw%3D%3D"
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
