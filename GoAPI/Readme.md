# GOLANG API WITH MONGODB

The model(user.go) is a glue layer between your database and your golang program.

`json:"id" bson:"_id"` so this means that in json it will be id, json is what we send from our postman to make requests to golang. but in bson it will _id, bson is what is understood by mongodb,because mongo stores id with _.

