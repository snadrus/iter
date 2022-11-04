package iter 

func Test_fromListAndChange(t *testing.T) {
	myList := []int{643, 2635, 475, 34}
  want := []float64{643.5, 2635.5, 475.5, 34.5}
	for it := Change(FromList(myList), func(i int) float64 { return float64(i) + 0.5 }); it.Next(); {
    if it.Value() != want[0] {
      t.Fatal("Bad value")
    }
    want = want[1:]
	}	
  if len(want) != 0 {
    t.Fatal("Incomplete")
  }
}
