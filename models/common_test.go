package models

import (
	"testing"

	"note.cstdlib.com/models"
)

func TestInt64ToStr62(t *testing.T) {
	int64List := []int64{0, 1, 61, 62, 839299365868340223, 1014}
	str62List := []string{"1", "R", "i", "R1", "iiiiiiiiii", "Tf"}

	for i := 0; i < len(int64List); i++ {
		ans := models.Int64ToStr62(int64List[i])
		if str62List[i] != ans {
			t.Error("int64:", int64List[i], ", expect str62:", str62List[i], ", actual:", ans)
		}
	}
}

func TestStr62ToInt64(t *testing.T) {
	str62List := []string{"1", "R", "i", "R1", "iiiiiiiiii", "Tf"}
	int64List := []int64{0, 1, 61, 62, 839299365868340223, 1014}

	for i := 0; i < len(str62List); i++ {
		ans := models.Str62ToInt64(str62List[i])
		if int64List[i] != ans {
			t.Error("str62:", str62List[i], ", expect int64:", int64List[i], ", actual:", ans)
		}
	}
}

func TestGenerateRandomStr62(t *testing.T) {
	strmp := make(map[string]int)
	times := 1000
	for i := 0; i < times; i++ {
		str := models.GenerateRandomStr62(8)
		if len(str) != 8 {
			t.Error("expect length:", 8, ", actual length:", len(str))
		}
		if _, ok := strmp[str]; ok {
			t.Error("generate same str in", times, "times")
			continue
		}
		strmp[str] = 1
	}

	if str := models.GenerateRandomStr62(0); str != "" {
		t.Error("expect empty string, actual:", str)
	}
	if str := models.GenerateRandomStr62(11); str != "" {
		t.Error("expect empty string, actual:", str)
	}
	if str := models.GenerateRandomStr62(-1); str != "" {
		t.Error("expect empty string, actual:", str)
	}
}
