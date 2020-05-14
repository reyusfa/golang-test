package palindromenumber

import "testing"

func TestCountPalindromesError(t *testing.T) {
	t.Run("when argument is an empty string", func (t *testing.T) {
		_, err := CountPalindromes("")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := CountPalindromes("1 ")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := CountPalindromes("1a 2")

		if err == nil {
			t.Fatal("expected an error")
		}
	})
}

func TestCountPalindromesResult(t *testing.T) {
	t.Run("when argument is '1 10'", func (t *testing.T) {
		got, _ := CountPalindromes("1 10")
		want := 9

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '99 100'", func (t *testing.T) {
		got, _ := CountPalindromes("99 100")
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
