package structs

import "testing"

//func TestArea(t *testing.T) {
//
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		// Helper will call the error within the line where checkArea is called
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12.0, 6.0}
//		checkArea(t, rectangle, 72.0)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//}

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
