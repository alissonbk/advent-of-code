package partone

type gamePart struct {
	color string
	qty   int
	valid bool
}

// only 12 red cubes, 13 green cubes, and 14 blue cubes
func (gp *gamePart) validate() {
	switch gp.color {
	case "red":
		if gp.qty > 12 {
			gp.valid = false
		}
		break
	case "green":
		if gp.qty > 13 {
			gp.valid = false
		}
		break
	case "blue":
		if gp.qty > 14 {
			gp.valid = false
		}
		break
	default:
		panic("Could not find gamePart color!")
	}
}

type gameSet []gamePart
