package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"Title", "Post"})

	if len(feed.Items) != 1 {
		t.Errorf("Error in adding feed")
	}

}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"Title", "Post"})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Error in adding feed")
	}
}
