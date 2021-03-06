package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func removeIngredientFromAllAllergens(allergens map[string]mapset.Set, ingredient string) {
	for _, allergenSet := range allergens {
		allergenSet.Remove(ingredient)
	}
}

func main() {
	allergens := make(map[string]mapset.Set, 0)
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ingredientSet := make(map[string]struct{})
	for _, food := range strings.Split(string(b), "\n") {
		if food == "" {
			continue
		}
		re := regexp.MustCompile(`^(.*)\(contains (.*)\)$`)
		match := re.FindStringSubmatch(food)
		ingredientsArr := strings.Split(match[1], " ")
		allergensArr := strings.Split(match[2], ", ")
		for _, allergen := range allergensArr {
			allergenSet := mapset.NewSet()
			for _, ingredient := range ingredientsArr {
				if ingredient == "" {
					continue
				}
				allergenSet.Add(ingredient)
				ingredientSet[ingredient] = struct{}{}
			}

			if _, exists := allergens[allergen]; exists {
				allergens[allergen] = allergens[allergen].Intersect(allergenSet)
			} else {
				allergens[allergen] = allergenSet
			}
		}
	}

	done := false
	for !done {
		done = true
		for _, allergenSet := range allergens {
			if allergenSet.Cardinality() == 1 {
				ingredient := allergenSet.Pop().(string)
				removeIngredientFromAllAllergens(allergens, ingredient)
				allergenSet.Add(ingredient)
				delete(ingredientSet, ingredient)
			} else {
				done = false
			}
		}
	}

	count := 0

	for _, food := range strings.Split(string(b), "\n") {
		if food == "" {
			continue
		}
		for ingredient := range ingredientSet {
			if strings.Contains(food, " "+ingredient+" ") || strings.HasSuffix(food, " "+ingredient) || strings.HasPrefix(food, ingredient+" ") {
				count++
			}
		}
	}

	fmt.Println(count)
}
