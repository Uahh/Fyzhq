package fyzhq

import "testing"

func TestTransform(t *testing.T) {
	text := Transform("你好,我很可爱,请给我钱", "jp")
	t.Log(text)
}
