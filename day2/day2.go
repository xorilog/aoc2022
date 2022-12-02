package day2

const (
	Rock     int = 1
	Paper    int = 2
	Scissors int = 3
	Lost     int = 0
	Draw     int = 3
	Win      int = 6
)

// MatchResult takes two weapons as int, (the player is a, b is the adversary) do the RockPaperScissors operation and
// return the winner as int
func MatchResult(a, b int) int {
	var result int

	if a == b {
		return Draw
	}
	switch a {
	case Rock:
		if b == Paper {
			result = Lost
		} else {
			result = Win
		}
	case Paper:
		if b == Scissors {
			result = Lost
		} else {
			result = Win
		}
	case Scissors:
		if b == Rock {
			result = Lost
		} else {
			result = Win
		}
	}
	return result
}

// GetWeaponMapping takes a string and return its weapon int equivalent
func GetWeaponMapping(s string) int {
	var weapon int
	switch s {
	case "A", "X":
		weapon = Rock
	case "B", "Y":
		weapon = Paper
	case "C", "Z":
		weapon = Scissors
	}
	return weapon
}

// GetRoundStrategy takes a string and return it's strategy int equivalent
func GetRoundStrategy(s string) int {
	var strategy int
	switch s {
	case "X":
		strategy = Lost
	case "Y":
		strategy = Draw
	case "Z":
		strategy = Win
	}
	return strategy
}

// SumOfGames takes a slices of string as input and return the amount of point accumulated as int
func SumOfGames(l []string) int {
	var sum int

	for _, v := range l {

		a := GetWeaponMapping(string(v[0]))
		b := GetWeaponMapping(string(v[2]))

		result := MatchResult(b, a)
		sum = sum + b + result
	}

	return sum
}

// SumOfGamesRoundStrategy takes a slices of string as input and return the amount of point accumulated as int when
// using the Win, Lose, Draw Strategy
func SumOfGamesRoundStrategy(l []string) int {
	var sum int

	for _, v := range l {

		a := GetWeaponMapping(string(v[0]))
		b := GetRoundStrategy(string(v[2]))

		for _, s := range []string{"A", "B", "C"} {
			w := GetWeaponMapping(s)
			if result := MatchResult(w, a); result == b {
				sum = sum + w + result
			}
		}
	}

	return sum
}
