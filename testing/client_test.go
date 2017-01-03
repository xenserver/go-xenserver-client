package testing


import (
	"testing"
	"flag" 
)

func init(){
	flag.StringVar(&testUrl, "url", "","")
	flag.StringVar(&testUsername, "user", "","")
	flag.StringVar(&testPassword, "pass", "","")
	flag.Parse()
}

func TestClient_StartVM (t *testing.T) {
	env := newTestEnvironment(t)
	defer env.Close()


}

