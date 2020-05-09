package support

import (
	"encoding/json"
	"testing"
)

func TestJSONIndent(t *testing.T) {
	data := map[string]interface{}{
		"hello": 1,
		"world": 88,
	}

	normalWant := `{"hello":1,"world":88}`

	indentWant := `{
    "hello": 1,
    "world": 88
}`

	m1, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	if normalGot := string(m1); normalGot != normalWant {
		t.Errorf("data={%v}, want %s, normalGot %s", data, normalWant, normalGot)
	}

	if indentGot := string(JSONIndent(data)); indentGot != indentWant {
		t.Errorf("data={%v}, want %s, indentGot %s", data, indentWant, indentGot)
	}
}
