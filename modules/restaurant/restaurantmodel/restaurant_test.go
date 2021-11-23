package restaurantmodel

import "testing"

type DataTable struct {
	Input RestaurantCreate
	Expect error
}

func TestValidate(t *testing.T) {
	table := []DataTable{
		{Input: RestaurantCreate{
			Name: "",
		}, Expect: ErrNameCannotBeEmpty},
		{Input: RestaurantCreate{
			Name: "I have title",
		}, Expect: nil},
	}

	for i := range table {
		err := table[i].Input.Validate()
		if err != table[i].Expect {
			t.Errorf("Test Validate() failed, expected %v, but got %v", table[i].Expect, err.Error())
		}
	}

	t.Log("Test restaurant Validate() passed")
}
