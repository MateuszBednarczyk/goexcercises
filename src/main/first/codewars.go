package first

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3

	if avg >= 0 && avg < 60 {
		return 'F'
	} else if avg >= 60 && avg < 70 {
		return 'D'
	} else if avg >= 70 && avg < 80 {
		return 'C'
	} else if avg >= 80 && avg < 90 {
		return 'B'
	} else {
		return 'A'
	}
}
