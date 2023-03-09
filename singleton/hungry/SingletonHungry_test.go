package hungry

import "testing"

func TestSingletonHungry(t *testing.T) {
	s1 := GetSingelton()
	s2 := GetSingelton()
	t.Log("s1 == s2 ? ", s1 == s2)
}
