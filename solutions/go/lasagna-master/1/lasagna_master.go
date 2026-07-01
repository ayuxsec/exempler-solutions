package lasagnamaster

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	if avgPrepTimePerLayer == 0 {
        avgPrepTimePerLayer = 2
    }
    return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64)  {
    for i := range layers {
        if layers[i] == "noodles" {
            noodles += 50
        }
        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList []string, myList []string) ([]string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
    return myList
}

func ScaleRecipe(quantitiesForTwo []float64, amountToCook int) ([]float64) {
    scaledRecipe := make([]float64, len(quantitiesForTwo))

    for i, v := range quantitiesForTwo {
        scaledRecipe[i] = v / 2.0 * float64(amountToCook)
    }

    return scaledRecipe
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
