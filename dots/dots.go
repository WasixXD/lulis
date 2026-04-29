package dots

const BatchSize = 10

func Naive(a [BatchSize][9]int, b [9]int) int {
	sum := 0
	for i, x := 0, 0; i < 9 && x < BatchSize; i, x = i+1, x+1 {
		sum += a[x][i] * b[i]
	}
	return sum
}

func Unroll(a [BatchSize][9]int, b [9]int) int {
	sum := 0
	for x := 0; x < BatchSize; x++ {
		sum +=
			a[x][0]*b[0] +
				a[x][1]*b[1] +
				a[x][2]*b[2] +
				a[x][3]*b[3] +
				a[x][4]*b[4] +
				a[x][5]*b[5] +
				a[x][6]*b[6] +
				a[x][7]*b[7] +
				a[x][8]*b[8]
	}
	return sum
}

func BCE(a [BatchSize][9]int, b [9]int) int {
	sum := 0
	for x := 0; x < BatchSize; x++ {
		for i := 0; i < 9; i += 3 {
			aTmp := a[x][i : i+3 : i+3]
			bTmp := b[i : i+3 : i+3]
			sum +=
				aTmp[0]*bTmp[0] +
					aTmp[1]*bTmp[1] +
					aTmp[2]*bTmp[2]
		}
	}
	return sum
}

func FullUnroll(a [4][9]int, b [9]int) int {
	return (a[0][0]*b[0] +
		a[0][1]*b[1] +
		a[0][2]*b[2] +
		a[0][3]*b[3] +
		a[0][4]*b[4] +
		a[0][5]*b[5] +
		a[0][6]*b[6] +
		a[0][7]*b[7] +
		a[0][8]*b[8] +

		a[1][0]*b[0] +
		a[1][1]*b[1] +
		a[1][2]*b[2] +
		a[1][3]*b[3] +
		a[1][4]*b[4] +
		a[1][5]*b[5] +
		a[1][6]*b[6] +
		a[1][7]*b[7] +
		a[1][8]*b[8] +

		a[2][0]*b[0] +
		a[2][1]*b[1] +
		a[2][2]*b[2] +
		a[2][3]*b[3] +
		a[2][4]*b[4] +
		a[2][5]*b[5] +
		a[2][6]*b[6] +
		a[2][7]*b[7] +
		a[2][8]*b[8] +

		a[3][0]*b[0] +
		a[3][1]*b[1] +
		a[3][2]*b[2] +
		a[3][3]*b[3] +
		a[3][4]*b[4] +
		a[3][5]*b[5] +
		a[3][6]*b[6] +
		a[3][7]*b[7] +
		a[3][8]*b[8])
}

func Dot4(a [BatchSize][9]int, b [9]int) (int, int, int, int) {
	var s0, s1, s2, s3 int
	for i := 0; i < 9; i++ {
		w := b[i]
		s0 += a[0][i] * w
		s1 += a[1][i] * w
		s2 += a[2][i] * w
		s3 += a[3][i] * w
	}
	return s0, s1, s2, s3
}

func Sum4(a [BatchSize][9]int) (int, int, int, int) {
	var s0, s1, s2, s3 int
	for i := 0; i < 9; i++ {
		s0 += a[0][i]
		s1 += a[1][i]
		s2 += a[2][i]
		s3 += a[3][i]
	}
	return s0, s1, s2, s3
}
