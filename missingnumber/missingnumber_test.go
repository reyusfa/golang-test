package missingnumber

import "reflect"
import "testing"

func TestFindMissingNumberError(t *testing.T) {
	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := FindMissingNumber("1234567890")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := FindMissingNumber("111210")

		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("when argument is not valid", func (t *testing.T) {
		_, err := FindMissingNumber("1234567891011")

		if err == nil {
			t.Fatal("expected an error")
		}
	})
}

func TestFindMissingNumberResult(t *testing.T) {
	t.Run("when argument is '23242526272830'", func (t *testing.T) {
		got, _ := FindMissingNumber("23242526272830")
		want := []int{29}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '101102103104105106107108109111112113'", func (t *testing.T) {
		got, _ := FindMissingNumber("101102103104105106107108109111112113")
		want := []int{110}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '12346789'", func (t *testing.T) {
		got, _ := FindMissingNumber("12346789")
		want := []int{5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '79101112'", func (t *testing.T) {
		got, _ := FindMissingNumber("79101112")
		want := []int{8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '7891012'", func (t *testing.T) {
		got, _ := FindMissingNumber("7891012")
		want := []int{11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '9799100101102'", func (t *testing.T) {
		got, _ := FindMissingNumber("9799100101102")
		want := []int{98}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when argument is '100001100002100003100004100006'", func (t *testing.T) {
		got, _ := FindMissingNumber("100001100002100003100004100006")
		want := []int{100005}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
