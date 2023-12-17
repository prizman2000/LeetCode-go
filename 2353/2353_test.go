package foodRatings

import (
	"testing"
	"reflect"
)

func Test_foodRatings(t *testing.T) {
	type args struct {
		foods 		[]string
		cuisines	[]string
		ratings		[]int
	
	}
	test := struct {
		name string
		args args
		want string
	} {
		name: "Example 1", 
		args: args{
			foods: []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
			cuisines: []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
			ratings: []int{9, 12, 8, 15, 14, 7},
		}, 
		want: "sushi",
	}
	t.Run(test.name, func(t *testing.T) {
		obj := Constructor(test.args.foods, test.args.cuisines, test.args.ratings)
		obj.ChangeRating("sushi", 17)
		if got := obj.HighestRated("japanese"); !reflect.DeepEqual(got, test.want) {
			t.Errorf("HighestRated() = %v, want %v", got, test.want)
		}
	}) 
}