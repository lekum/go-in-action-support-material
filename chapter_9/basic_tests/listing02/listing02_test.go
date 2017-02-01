package listing02

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?=alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading content.")
	{
		for _, tc := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", tc.url, tc.statusCode)
			{
				resp, err := http.Get(tc.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", checkMark)
				defer resp.Body.Close()

				if resp.StatusCode == tc.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v", tc.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShoud receive a \"%d\" status. %v %v", tc.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}

}
