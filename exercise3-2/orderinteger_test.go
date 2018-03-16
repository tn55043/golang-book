package test_orderinteger

import ("testing")

func Test_given_50219_display_95021(t *testing.T) {
	slice := []int{50, 2, 1, 9}
    display := orderintegers(slice)
    if display != "95021" {
        t.Error("orderintegers{50, 2, 1, 9} != \"95021\" Actual is " + display)
    }
}

func Test_given_55056_display_56550(t *testing.T) {
	slice := []int{5, 50, 56}
    display := orderintegers(slice)
    if display != "56550" {
        t.Error("orderintegers{5, 50, 56} != \"56550\" Actual is " + display)
    }
}

func Test_given_42042423_display_42423420(t *testing.T) {
	slice := []int{420, 42, 423}
    display := orderintegers(slice)
    if display != "42423420" {
        t.Error("orderintegers{420, 42, 423} != \"42423420\" Actual is " + display)
    }
}
