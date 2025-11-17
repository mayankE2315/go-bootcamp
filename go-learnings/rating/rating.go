package rating

import (
	"fmt"
	"time"
)

type Rating struct {
	id        int
	avgRating float64
	ratings   []UserRating
}

type UserRating struct {
	uid      int
	rating   float64
	comments Comment
}

type Comment struct {
	comment string
	date    time.Time
}

func (r *Rating) Add(uid int, rating float64, comment string) error {
	if rating < 0 || rating > 5 {
		return fmt.Errorf("invalid input %f", rating)
	}
	r.ratings = append(r.ratings, UserRating{uid: uid, rating: rating, comments: Comment{comment: comment, date: time.Now()}})
	r.avgRating = (r.avgRating + rating) / float64(len(r.ratings))
	return nil
}

func (r Rating) String() string {
	return fmt.Sprintf("Rating{id: %d, avgRating: %f, ratings: %v}", r.id, r.avgRating, r.ratings)
}

func (u UserRating) String() string {
	return fmt.Sprintf("UserRating{uid: %d, rating: %f, comments: %v}", u.uid, u.rating, u.comments)
}

func (c Comment) String() string {
	return fmt.Sprintf("Comment{comment: %s, date: %s}", c.comment, c.date.String())
}
