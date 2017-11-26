package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/CityRescue/models/disasterrecovery"
)

func FetchPostsEndpoint(w http.ResponseWriter, r *http.Request) {

	// TODO: Implement fetching posts for a given user

	// We are going to create some mock data and send it out in JSON
	// format.

	// We will actually implement this endpoint, when we cover database
	// persistence later in the course.

	v := mux.Vars(r)

	//w.Write([]byte(v["username"]))
	if v["username"] == "kruti" {

		mockPosts := make([]disasterrecovery.Post, 3)

		post1 := disasterrecovery.NewPost("Kruti", disasterrecovery.Moods["hopeful"], "I need Help!", "Anyone who can help, I will really appreciate!", "https://fasniche.com", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post2 := disasterrecovery.NewPost("Kruti", disasterrecovery.Moods["thrilled"], "Yeyy!!", "Now I've got help!", "https://fasniche.com", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post3 := disasterrecovery.NewPost("Kruti", disasterrecovery.Moods["happy"], "My Helper is awesome", "My helper made sure I was helped!", "https://fasniche.com", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)
		json.NewEncoder(w).Encode(mockPosts)

	} else {
		json.NewEncoder(w).Encode(nil)

	}

}
