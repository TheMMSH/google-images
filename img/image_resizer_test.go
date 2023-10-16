package img

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestResizeImageResizesTestFiles(t *testing.T) {
	asserts := assert.New(t)
	sut := New(100, 200)

	f, _ := os.ReadFile("./data/test/test_1.jpg")
	validOut, _ := os.ReadFile("./data/validout/test_1.jpg")
	res, _ := sut.ResizeImage(f)

	asserts.Equal(validOut, res)

	f2, _ := os.ReadFile("./data/test/test_2.jpg")
	validOut2, _ := os.ReadFile("./data/validout/test_2.jpg")
	res2, _ := sut.ResizeImage(f2)

	asserts.Equal(validOut2, res2)
}
