package ch4

import "testing"

func TestSwitch(t *testing.T){
	for i := 0; i <= 5; i++{
		switch i {
		case 0,2:
			t.Log(i,"Even")
		case 1,3:
			t.Log(i,"Odd")
		default:
			t.Log(i,"is not Even or Odd in 1-3")
		}
	}
}

func TestSwitchCondition(t *testing.T){
	for i := 0; i <= 6; i++{
		switch  {
		case 0 <= i && i <= 2:
			t.Log(i, "between 0-2")
		case 3 <= i && i <= 5:
			t.Log(i, "between 3-5")
		default:
			t.Log(i)
		}

		switch  {
		case i % 2 == 0:
			t.Log(i,"is Even")
		case i % 2 == 1:
			t.Log(i,"is Odd")
		default:
			t.Log("emm")
		}
	}
}
