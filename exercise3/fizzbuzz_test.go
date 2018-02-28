package test_fizzbuzz

import ("testing")

func Test_equal_1_display_1(t *testing.T) {
    display := fizzbuzz(1)
    if display != "1" {
        t.Error("fizzbuzz(1) != \"1\" Actual is " + display)
    }
}

func Test_morethan_1_divide_3_equal_0_display_Fizz(t *testing.T) {
    display := fizzbuzz(3)
    if display != "Fizz" {
        t.Error("fizzbuzz(3) != \"Fizz\" Actual is " + display)
    }
}

func Test_morethan_1_divide_3_notequal_0_display_4(t *testing.T) {
    display := fizzbuzz(4)
    if display != "4" {
        t.Error("fizzbuzz(4) != \"4\" Actual is " + display)
    }
}

func Test_morethan_1_divide_5_equal_0_display_Buzz(t *testing.T) {
    display := fizzbuzz(5)
    if display != "Buzz" {
        t.Error("fizzbuzz(5) != \"Buzz\" Actual is " + display)
    }
}

func Test_morethan_1_divide_5_notequal_0_display_7(t *testing.T) {
    display := fizzbuzz(7)
    if display != "7" {
        t.Error("fizzbuzz(7) != \"7\" Actual is " + display)
    }
}

func Test_morethan_1_divide_3_and_5_equal_0_display_FizzBuzz(t *testing.T) {
    display := fizzbuzz(15)
    if display != "FizzBuzz" {
        t.Error("fizzbuzz(15) != \"FizzBuzz\" Actual is " + display)
    }
}

func Test_morethan_1_divide_3_and_5_notequal_0_display_17(t *testing.T) {
    display := fizzbuzz(17)
    if display != "17" {
        t.Error("fizzbuzz(17) != \"17\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_3_equal_0_display_Fizz(t *testing.T) {
    display := fizzbuzz(99)
    if display != "Fizz" {
        t.Error("fizzbuzz(99) != \"Fizz\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_3_notequal_0_display_4(t *testing.T) {
    display := fizzbuzz(98)
    if display != "98" {
        t.Error("fizzbuzz(98) != \"98\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_5_equal_0_display_Buzz(t *testing.T) {
    display := fizzbuzz(95)
    if display != "Buzz" {
        t.Error("fizzbuzz(95) != \"Buzz\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_5_notequal_0_display_7(t *testing.T) {
    display := fizzbuzz(94)
    if display != "94" {
        t.Error("fizzbuzz(94) != \"94\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_3_and_5_equal_0_display_FizzBuzz(t *testing.T) {
    display := fizzbuzz(90)
    if display != "FizzBuzz" {
        t.Error("fizzbuzz(90) != \"FizzBuzz\" Actual is " + display)
    }
}

func Test_lessthan_100_divide_3_and_5_notequal_0_display_17(t *testing.T) {
    display := fizzbuzz(98)
    if display != "98" {
        t.Error("fizzbuzz(98) != \"98\" Actual is " + display)
    }
}

func Test_equal_100_display_Buzz(t *testing.T) {
    display := fizzbuzz(100)
    if display != "Buzz" {
        t.Error("fizzbuzz(100) != \"Buzz\" Actual is " + display)
    }
}

func Test_lessthan_1_notdisplay(t *testing.T) {
    display := fizzbuzz(0)
    if display != "" {
        t.Error("fizzbuzz(0) != \"\" Actual is " + display)
    }
}

func Test_morethan_100_notdisplay(t *testing.T) {
    display := fizzbuzz(101)
    if display != "" {
        t.Error("fizzbuzz(101) != \"\" Actual is " + display)
    }
}
