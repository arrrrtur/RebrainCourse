package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseInt(t *testing.T) {
	if reversed, _ := ReverseInt(123); reversed != 321 {
		t.Errorf("123 is not 321")
	}

	if reversed, _ := ReverseInt(-649); reversed != -946 {
		t.Errorf("-649 is not -946")
	}

	if reversed, _ := ReverseInt(0); reversed != 0 {
		t.Errorf("0 is not 0")
	}
}

func TestReverseIntWithTestify(t *testing.T) {
	req := require.New(t)
	reversed, err := ReverseInt(123)
	req.NoError(err)
	req.Equal(reversed, 321)

}
