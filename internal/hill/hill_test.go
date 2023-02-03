package hill

import (
	"fmt"
	"github.com/ravielze/Crypto-1/internal/utils"
	"testing"
)

func convertIndexToRowAndColumn(index int) (int, int) {
	return index / 3, index % 3
}

func Cryptanalysis(plain string, encrypted string) string {
	matrixPlain := utils.InitializeMatrix(3, 3)
	matrixEncrypted := utils.InitializeMatrix(3, 3)

	for i := 0; i < 9; i++ {
		row, col := convertIndexToRowAndColumn(i)
		matrixPlain[col][row] = int(plain[i]) - 65
		matrixEncrypted[col][row] = int(encrypted[i]) - 65
	}

	inverseMat := matrixPlain.Inverse()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			inverseMat[i][j] = utils.Modulo(inverseMat[i][j], 26)
		}
	}

	result := matrixEncrypted.Multiply(inverseMat)

	var key []rune
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = utils.Modulo(result[i][j], 26) + 65
			key = append(key, rune(result[i][j]))
		}
	}
	return string(key)
}

func TestHill_Cryptanalysis(t *testing.T) {
	knownPlain := "HELLOCAPTAINHADDOCK"
	expected := "TFJOXUPOUXYTTRDSXQMONIYPEUFJDQUBGIMOCJQTNBEHCZEKROVBNTWLMVXMOWZLUCHOXYGSKBQGUAOBQZKIXYJIETSWVXHVKCUAOTOFYIZAKJGXKAWGQTRVFDZAJNQDUIWZCMYWNFIUPYMCZXIAKYUCQIAZPIQMGAMGUAKKKHMWKDUXQDUAAKYOWEHLJPWYFKXSARBLLHGAJKTQNTRTPWSCIZASCGSLKVDHTUZSWBNBTJGYYUPQMFSYZAUTOQCDNGQMFSRLRTUWEMKADIVYLTJKFHLKJUWTSSHMHJFGTRIBYIDAHQEPMPIQCROWDYRYZNSPNOJHQVKKTOCBPNFAJNLYJZNVBAYJWRGMCHJPWBDHHTPOXSIJVQWDMSIGMTRVEVXDILKVAYTNUNJXEZLAPGYETRVZNVHSVWLGICDXQFOALDVPASUSYXPFHUWTILUQHTJQVGWFSPAEKBRBNIINYKHNTNUKJVDHVLXQKUZNVQXUOZZOJZYNPIVYSVFVTZMMUUPWTGHRIOWCBKZYAGUMRCKHIQZSIGISPGBXPYXMOAWGAGHQVUWTEIGPBMOMBWIOPQEVKMRQATNBMILHHLVUXGMOUWTZCLBKGWIJHFRNGOSCMUHDWHBB"

	for i := 0; i < len(expected)-9; i++ {
		cryptanalysisKey := Cryptanalysis(knownPlain[0:9], expected[i:i+9])

		decrypt := NewHill(cryptanalysisKey)
		decrypted := decrypt.Decrypt(expected)

		fmt.Println(i, utils.GenerateKeyMatrix(cryptanalysisKey, 3), decrypted)
	}
}
