package hello


import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, test!"
	got := Say("Test")

	if want != got {
		t.Errorf("wanted %s, got %s, want, got")
	}
}