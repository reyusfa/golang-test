package organizebook

import "testing"

func TestOrganizeBooksError(t *testing.T) {
	t.Run("when argument is empty string", func (t *testing.T) {
		_, err := OrganizeBooks("")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := OrganizeBooks("3A13 AX19")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := OrganizeBooks("9E22 ")

		if err == nil {
			t.Fatal("expected an error")
		}
	})
}

func TestOrganizeBooksResult(t *testing.T) {
	t.Run("when argument is '3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14'", func (t *testing.T) {
		got, _ := OrganizeBooks("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
		want := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '3A13 5X19 9Y20'", func (t *testing.T) {
		got, _ := OrganizeBooks("3A13 5X19 9Y20")
		want := "9Y20 5X19 3A13"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
