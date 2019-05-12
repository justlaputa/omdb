package omdb

import "testing"

func TestByTitle(t *testing.T) {
	client := makeTestClient()

	m, err := client.ByTitle("The Mule", 2018)
	if err != nil {
		t.Fatalf("failed to search by title: %v", err)
	}

	t.Logf("%+v", m)
}

func makeTestClient() *Client {
	return New("<TEST_KEY>")
}
