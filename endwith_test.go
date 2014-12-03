package endswith

import "fmt"
import "testing"

func TestEndsWith(t *testing.T) {
	if EndsWith("Im a sentence", "sentence") {
		fmt.Println("Im a sentence ends with `sentence`")
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}

func TestEndsWithIn(t *testing.T) {
	if EndsWithIn("abc.css", []string{".css", ".js"}) {
		fmt.Println("abc.css ends with in ['.css','.js']")
		t.Log("Passed.")
	} else {
		t.Error("Failed.")
	}
}
