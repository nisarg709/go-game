package handlers

import (
	"eman/passport/achievements/src/services"
	"github.com/globalsign/mgo/bson"
	"encoding/json"
	"errors"
)

type SurveyCompletedPayload struct {
	UserId string `json:"user"`
}

func (self SurveyCompletedPayload) IsValid() bool {
	return self.UserId != ""
}

func SurveyCompletedHandler(context *services.Context, event *services.Event) error {

	payload := SurveyCompletedPayload{}
	err := json.Unmarshal([]byte(event.Data), &payload)

	if err != nil || payload.IsValid() != true {
		return errors.New("invalid event payload")
	}

	collection := context.Mongo.DB("passport").C("users")

	query := bson.M{"_id": bson.ObjectIdHex(payload.UserId), "badges": "survey_completed"}
	count, err := collection.Find(query).Count()

	if count > 0 {
		return nil // Abort: Achievement already exists
	}

	if err := collection.UpdateId(bson.ObjectIdHex(payload.UserId), bson.M{"$push": bson.M{"badges": "survey_completed"}}); err != nil {
		return err
	}

	return nil
}
