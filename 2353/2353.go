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


func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    for i, item := range this.Foods {
        if food == item {
            this.Ratings[i] = newRating
            return
        }
    }
}


func (this *FoodRatings) HighestRated(cuisine string) string {
    var highest, highestIndex int

    for i := 0; i < len(this.Foods); i++ {
        if this.Cuisines[i] == cuisine && this.Ratings[i] > highest {
            highestIndex = i
            highest = this.Ratings[i]
        }
    }
    return this.Foods[highestIndex]
}