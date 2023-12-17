package foodRatings

type FoodRatings struct {
    Foods    []string
    Cuisines []string
    Ratings  []int
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    return FoodRatings {
        Foods: foods,
        Cuisines: cuisines,
        Ratings: ratings,
    }
}


func (fr *FoodRatings) ChangeRating(food string, newRating int)  {
    for i, item := range fr.Foods {
        if food == item {
            fr.Ratings[i] = newRating
            return
        }
    }
}


func (fr *FoodRatings) HighestRated(cuisine string) string {
    var highest, highestIndex int

    for i := 0; i < len(fr.Foods); i++ {
        if fr.Cuisines[i] == cuisine && fr.Ratings[i] > highest {
            highestIndex = i
            highest = fr.Ratings[i]
        }
    }
    return fr.Foods[highestIndex]
}