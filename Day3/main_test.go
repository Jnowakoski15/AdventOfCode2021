package main

func Test1(t *testing.T) {
	const input = "test.txt"
	rez := part1(input)
	exp := 150
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func Test2(t *testing.T) {
	const input = "test.txt"
	rez := part2(input)
	exp := 900
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
