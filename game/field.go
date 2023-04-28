package game

// 移動のロジックを実装

func (field *Field) down() {
	for cnt := 0; cnt < 4; cnt++ {
		for i := 2; i >= 0; i-- {
			//  i行目
			for j := 0; j <= 3; j++ {
				if field[i+1][j] == 0 {
					field[i+1][j] = field[i][j]
					field[i][j] = 0
				}
			}
		}
	}
}
func (field *Field) up() {
	for cnt := 0; cnt < 4; cnt++ {
		for i := 1; i <= 3; i++ {
			for j := 0; j <= 3; j++ {
				//Jは触らない
				if field[i-1][j] == 0 {
					field[i-1][j] = field[i][j]
					field[i][j] = 0
				}
			}
		}
	}
}
func (field *Field) right() {
	for cnt := 0; cnt < 4; cnt++ {
		for i := 0; i <= 3; i++ {
			// iは触らない
			for j := 0; j < 3; j++ {
				if field[i][j+1] == 0 {
					field[i][j+1] = field[i][j]
					field[i][j] = 0
				}
			}
		}
	}
}
func (field *Field) left() {
	for cnt := 0; cnt < 4; cnt++ {
		for i := 0; i <= 3; i++ {
			// iは触らない
			for j := 3; j >= 1; j-- {
				if field[i][j-1] == 0 {
					field[i][j-1] = field[i][j]
					field[i][j] = 0
				}
			}
		}
	}
}

func (field *Field) downPlus() {
	for j := 0; j <= 3; j++ {
		//jは触らない
		// 2回足されるとき。
		if field[0][j] == field[1][j] && field[2][j] == field[3][j] {
			field[0][j] = 0
			field[1][j] *= 2
			field[2][j] = 0
			field[3][j] *= 2
		} else if field[2][j] == field[3][j] {
			field[2][j] = 0
			field[3][j] *= 2
		} else if field[1][j] == field[2][j] {
			field[2][j] *= 2
			field[1][j] = 0
		} else if field[0][j] == field[1][j] {
			field[0][j] = 0
			field[1][j] *= 2
		}
		field.down()
	}
}
func (field *Field) upPlus() {
	for j := 0; j <= 3; j++ {
		//jは触らない
		// 2回足されるとき。
		if field[0][j] == field[1][j] && field[2][j] == field[3][j] {
			field[0][j] = 0
			field[1][j] *= 2
			field[2][j] = 0
			field[3][j] *= 2
		} else if field[0][j] == field[1][j] {
			field[0][j] = 0
			field[1][j] *= 2
		} else if field[1][j] == field[2][j] {
			field[2][j] *= 2
			field[1][j] = 0
		} else if field[2][j] == field[3][j] {
			field[2][j] = 0
			field[3][j] *= 2
		}
		field.up()
	}
}
func (field *Field) rightPlus() {
	for i := 0; i <= 3; i++ {
		//iは触らない
		// 2回足されるとき。
		if field[i][0] == field[i][1] && field[i][2] == field[i][3] {
			field[i][0] *= 2
			field[i][1] = 0
			field[i][2] *= 2
			field[i][3] = 0
		} else if field[i][2] == field[i][3] {
			field[i][2] *= 2
			field[i][3] = 0
		} else if field[i][1] == field[i][2] {
			field[i][1] *= 2
			field[i][2] = 0
		} else if field[i][0] == field[i][1] {
			field[i][0] *= 2
			field[i][1] = 0
		}
		field.right()
	}
}
func (field *Field) leftPlus() {
	for i := 0; i <= 3; i++ {
		//iは触らない
		// 2回足されるとき。
		if field[i][0] == field[i][1] && field[i][2] == field[i][3] {
			field[i][0] *= 2
			field[i][1] = 0
			field[i][2] *= 2
			field[i][3] = 0
		} else if field[i][1] == field[i][2] {
			field[i][1] *= 2
			field[i][2] = 0
		} else if field[i][0] == field[i][1] {
			field[i][0] *= 2
			field[i][1] = 0
		} else if field[i][2] == field[i][3] {
			field[i][2] *= 2
			field[i][3] = 0
		}
		field.left()
	}
}
