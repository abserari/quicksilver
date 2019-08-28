package main

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

func main() {
	client, err := stream.NewClient(
		"4vhnphzf39ae",
		"ctbf5w4j3nfg4kwq4c3tudc2dakhyca9earf8pxbtrvcadxdjsezp986f46vm88r",
	)

	chris, err := client.FlatFeed("user", "chris")
	if err != nil {
		panic(err)
	}

	// Add an Activity; message is a custom field - tip: you can add unlimited custom fields!
	_, err := chris.AddActivity(stream.Activity{
		Actor:     "chris",
		Verb:      "add",
		Object:    "picture:10",
		ForeignID: "picture:10",
		Extra: map[string]interface{}{
			"message": "Beautiful bird!",
		},
	})
	if err != nil {
		panic(err)
	}

	// Create a following relationship between Jack's "timeline" feed and Chris' "user" feed:
	jack, err := client.FlatFeed("timeline", "jack")
	if err != nil {
		panic(err)
	}

	err = jack.Follow(chris)
	if err != nil {
		panic(err)
	}

	// Read Jack's timeline and Chris' post appears in the feed:
	resp, err := jack.GetActivities(stream.WithActivitiesLimit(10))
	if err != nil {
		panic(err)
	}
	for _, activity := range resp.Results {
		// ...
	}

	// Remove an Activity by referencing it's foreign_id
	err = chris.RemoveActivityByForeignID("picture:10")
	if err != nil {
		panic(err)
	}
}
