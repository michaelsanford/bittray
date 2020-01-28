package polling

import "testing"

func TestCheckForUpdate(t *testing.T) {
	const wantAvailable = false
	const wantLastTag = "1.1.3"

	gotAvailable, gotLastTag, err := CheckForUpdate()
	if err != nil {
		t.Error(err)
	}

	if gotAvailable != wantAvailable {
		t.Errorf("got '%t' want '%t'", gotAvailable, wantAvailable)
	}

	if gotLastTag != wantLastTag {
		t.Errorf("got '%s' want '%s'", gotLastTag, wantLastTag)
	}
}

func TestBackOff(t *testing.T) {
	const wantMax = 2
	const wantMin = 4
	got := backOff(wantMin)

	if wantMin <= got && got <= wantMax {
		t.Errorf("got '%d' want '%d' <= x <= '%d'", got, wantMin, wantMax)
	}
}

func BenchmarkCheckForUpdate(b *testing.B) {
	_, _, err := CheckForUpdate()
	if err != nil {
		b.Error(err)
	}
}
