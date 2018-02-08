package ingredient_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"

	"github.com/chvck/meal-planner/model/ingredient"
)

func TestString(t *testing.T) {
	i := ingredient.Ingredient{
		ID:       1,
		Measure:  null.String{NullString: sql.NullString{String: "tbsp"}},
		Name:     "Paprika",
		Quantity: 2,
		RecipeID: 2,
	}

	assert.Equal(t, "2 tbsp Paprika", fmt.Sprint(i))
}
