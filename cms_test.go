package differential_privacy

import (
	"fmt"
	"testing"
)

func TestCMS(t *testing.T) {
	k := 128
	hash := GenHash(k, 3)
	privacyBudget := 4
	client := &CMSClient{
		m:         3,
		kWiseHash: hash,
		prob:      calculateProb(privacyBudget),
	}
	server := &CMSServer{
		m:         3,
		c:         calculateC(float64(privacyBudget)),
		matrix:    GenMatric(k, 3),
		n:         0,
		kWiseHash: hash,
	}
	track := func(n int, data string) {
		for i := 0; i < n; i++ {
			e := client.Encode([]byte(data))
			server.Track(e)
		}
	}
	// we are tracking what view users are using in their project management tool.
	// 6k users are using list view.
	track(6000, "list")
	// 9k user using board view.
	track(9000, "board")
	// 2k user using calendar view.
	track(2000, "calendar")
	// Let's print all the estimate.
	fmt.Println("estimate for list", server.Estimate([]byte("list")))
	fmt.Println("estimate for board", server.Estimate([]byte("board")))
	fmt.Println("estimate for calendar", server.Estimate([]byte("calendar")))
	// output
	// estimate for list 6572.1029024055715
	// estimate for board 9154.186791339975
	// estimate for calendar 1157.801902649071
}
