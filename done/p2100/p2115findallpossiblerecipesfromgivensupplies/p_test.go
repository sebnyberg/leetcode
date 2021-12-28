package p2115findallpossiblerecipesfromgivensupplies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAllRecipes(t *testing.T) {
	for _, tc := range []struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
		want        []string
	}{
		{[]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}, []string{"bread"}},
		{[]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich"}},
		{[]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich", "burger"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.recipes), func(t *testing.T) {
			require.Equal(t, tc.want, findAllRecipes(tc.recipes, tc.ingredients, tc.supplies))
		})
	}
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Recipes/ingredients are the same in the sense that they are nodes in a
	// graph. The only difference is that ingredients have edges toward recipes
	adj := make([][]int, 0, len(recipes)+len(ingredients))
	recipesAndIngredients := make([]string, 0)
	nameToIdx := make(map[string]int)
	indeg := make([]int, 0, len(recipes)+len(ingredients))
	maybeAdd := func(name string) {
		if _, exists := nameToIdx[name]; !exists {
			nameToIdx[name] = len(recipesAndIngredients)
			recipesAndIngredients = append(recipesAndIngredients, name)
			adj = append(adj, make([]int, 0))
			indeg = append(indeg, 0)
		}
	}

	for i, recipe := range recipes {
		maybeAdd(recipe)
		ii := nameToIdx[recipe]
		for _, ingredient := range ingredients[i] {
			maybeAdd(ingredient)
			jj := nameToIdx[ingredient]
			adj[jj] = append(adj[jj], ii) // this ingredient leads to its recipe
			indeg[ii]++                   // one more requirement for the recipe
		}
	}

	for _, s := range supplies {
		if _, exists := nameToIdx[s]; !exists {
			continue
		}
		// Ingredient fulfilled
		for _, recipeIdx := range adj[nameToIdx[s]] {
			indeg[recipeIdx]--
		}
	}

	// Find recipes with an indegree of zero (no requirements)
	result := make([]string, 0)
	for _, recipe := range recipes {
		if indeg[nameToIdx[recipe]] == 0 {
			result = append(result, recipe)
		}
	}

	for i := 0; i < len(result); i++ {
		// Recipe has been cooked, change any recipes that depend on this one to
		// be fulfilled
		ii := nameToIdx[result[i]]
		for _, parentRecipeIdx := range adj[ii] {
			indeg[parentRecipeIdx]--
			if indeg[parentRecipeIdx] == 0 {
				result = append(result, recipesAndIngredients[parentRecipeIdx])
			}
		}
	}

	return result
}
