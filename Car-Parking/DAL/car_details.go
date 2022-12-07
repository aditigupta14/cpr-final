package DAL

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// User represents the user for this application
// swagger:model User
type User struct {

	// ID of the user
	// in: string
	// required: true
	UserID string

	// First Name of the user
	// in: string
	// required: true
	FName string

	// Last Name of the user
	// in: string
	// required: true
	LName string
}

// Car represents the car of the user
// swagger:model Car
type Car struct {
	// Carnumber of the car
	// in: string
	// required: true
	CarNumber string

	// Carmodel of the car
	// in: string
	CarModel string

	// Car Timein
	// in: time.Time
	TimeIn time.Time

	// Car Timeout
	// in: time.Time
	TimeOut time.Time

	// UserID of the user to whom the car belongs
	// in: string
	// required: true
	UserID string

	// Unique Slot Number in which the car is parked
	// in: string
	// required: true
	UniqueSlotNo string
}

// Slot represents the Parking Slot
// swagger:model Slot
type Slot struct {
	// Floor Number
	// in: int
	// required: true
	FloorNo int

	// Unique Slot Number
	// in: string
	// required: true
	UniqueSlotNo string

	// Occupancy of the slot
	// in: bool
	// required: true
	Occupancy bool
}

func AddUser(user User) error {

	session := Connect()

	collection := session.Database("webinardb").Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"userid": -1},
		Options: options.Index().SetUnique(true),
	}

	name, _ := collection.Indexes().CreateOne(context.TODO(), indexModel)
	// if err != nil {

	// 	return echo.NewHTTPError(400, err)

	// }

	fmt.Println("Name of Index Created: " + name)

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return echo.NewHTTPError(400, "Add a unique UserID")
	}

	log.Println(result)
	log.Println(err)

	return nil

}

func GetUser(query url.Values) []User {

	var users []User
	isPageQuery := query.Has("page")
	session := Connect()

	collection := session.Database("webinardb").Collection("users")
	if isPageQuery {

		page, _ := strconv.ParseInt(query["page"][0], 10, 64)

		var limit int64 = 5

		_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&users).Find()

		if err != nil {

			panic(err)

		}

	} else {

		cursor, _ := collection.Find(context.TODO(), bson.M{})

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {

			var user User

			err := cursor.Decode(&user)

			if err != nil {

				log.Fatal(err)

			}

			users = append(users, user)

		}

	}

	return users

}

//start
// page, _ := strconv.ParseInt(query["page"][0], 10, 64)
// //page1, _ := strconv.ParseInt(query["page"][1], 10, 64)
// fmt.Println(page)
// //fmt.Println(page1)
// var limit int64 = 5

// if page > 0 {
// 	_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&users).Find()

// 	if err != nil {

// 		panic(err)

// 	}
// }
//end
// cur, _ := collection.Find(context.TODO(), bson.M{})

// var limit int64 = 4
// var page int64 = 1

// pfilter := bson.M{}
// paginatedData, err := New(collection).Limit(limit).Page(page).Sort("floorno", -1).Filter(pfilter).Decode(&users).Find()
// if err != nil {
// 	panic(err)
// }

// // print pagination data
// fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)

// defer cur.Close(context.TODO())

// for cur.Next(context.TODO()) {

// 	var user User

// 	err := cur.Decode(&user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	users = append(users, user)
// }
// 	return users
// }

func UpdateUserByID(userid string, user User) {

	session := Connect()

	collection := session.Database("webinardb").Collection("users")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"userid": userid}
	update := bson.M{"$set": bson.M{"fname": user.FName, "lname": user.LName}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	fmt.Println(result, err)

}

func DeleteUserByID(userid string) {

	session := Connect()

	collection := session.Database("webinardb").Collection("users")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"userid": userid}
	fmt.Println(filter)

	result, err := collection.DeleteOne(context.TODO(), filter)

	fmt.Println(result, err)

	// if err != nil {
	// 	log.Println(err)

	// }

}

func AddCar(car Car) error {

	session := Connect()

	collection := session.Database("webinardb").Collection("cars")
	//collection2 := session.Database("webinardb").Collection("slots")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"carnumber": -1},
		Options: options.Index().SetUnique(true),
	}

	name, _ := collection.Indexes().CreateOne(context.TODO(), indexModel)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Name of Index Created: " + name)

	//car.TimeIn = time.Now()
	//car.TimeOut = time.Time{}

	result, err := collection.InsertOne(context.TODO(), car)
	//car.CarID = result.InsertedID.(primitive.ObjectID)

	//idPrimitive, _ := primitive.ObjectIDFromHex(car.SlotID)
	//fmt.Println(idPrimitive)

	// filter := bson.M{"uniqueslotno": car.UniqueSlotNo}
	// update := bson.M{"$set": bson.M{"occupancy": false}}

	// result2, err2 := collection2.UpdateOne(context.TODO(), filter, update)

	fmt.Println(result, err)
	if err != nil {
		return echo.NewHTTPError(400, "Add a unique Car Number")
	}
	//fmt.Println(result2, err2)
	return nil
}

func GetCar(query url.Values) []Car {

	var cars []Car
	isPageQuery := query.Has("page")
	session := Connect()

	collection := session.Database("webinardb").Collection("cars")

	if isPageQuery {

		page, _ := strconv.ParseInt(query["page"][0], 10, 64)

		var limit int64 = 5

		_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&cars).Find()

		if err != nil {

			panic(err)

		}

	} else {

		cursor, _ := collection.Find(context.TODO(), bson.M{})

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {

			var car Car

			err := cursor.Decode(&car)

			if err != nil {

				log.Fatal(err)

			}

			cars = append(cars, car)

		}

	}

	// page, _ := strconv.ParseInt(query["page"][0], 10, 64)

	// var limit int64 = 5

	// if page > 0 {

	// 	_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&cars).Find()

	// 	if err != nil {

	// 		panic(err)

	// 	}
	// }
	// cur, _ := collection.Find(context.TODO(), bson.M{})

	// var limit int64 = 4
	// var page int64 = 1

	// pfilter := bson.M{}
	// paginatedData, err := New(collection).Limit(limit).Page(page).Sort("floorno", -1).Filter(pfilter).Decode(&cars).Find()
	// if err != nil {
	// 	panic(err)
	// }

	// // print pagination data
	// fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)

	// defer cur.Close(context.TODO())

	// for cur.Next(context.TODO()) {

	// 	var car Car

	// 	err := cur.Decode(&car)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cars = append(cars, car)
	// }
	return cars

}

func GetCarByNumber(carnumber string) Car {

	var car Car
	session := Connect()

	collection := session.Database("webinardb").Collection("cars")

	filter := bson.M{"carnumber": carnumber}

	err := collection.FindOne(context.TODO(), filter).Decode(&car)

	if err != nil {
		log.Println(err)

	}
	return car
}

func UpdateCar(carnumber string, car Car) {

	session := Connect()

	collection := session.Database("webinardb").Collection("cars")
	//collection2 := session.Database("webinardb").Collection("slots")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)
	//idPrimitive2, _ := primitive.ObjectIDFromHex(car.Slot.ID)
	//fmt.Println(idPrimitive) //, idPrimitive2)

	filter := bson.M{"carnumber": carnumber}
	update := bson.M{"$set": bson.M{"carmodel": car.CarModel, "carnumber": car.CarNumber}} //, "slot": bson.M{"floorno": car.Slot.FloorNo, "uniqueslotno": car.Slot.UniqueSlotNo}}}

	//filter2 := bson.M{"_id": idPrimitive2}
	//update2 := bson.M{"$set": bson.M{"occupancy": car.Slot.Occupancy, "timeout": car.Slot.TimeOut}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	//result2, err2 := collection2.UpdateOne(context.TODO(), filter2, update2)

	fmt.Println(result, err)
	//fmt.Println(result2, err2)

}

func UpdateCarTimeIn(carnumber string) {

	session := Connect()

	collection := session.Database("webinardb").Collection("cars")
	collection2 := session.Database("webinardb").Collection("slots")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)
	//idPrimitive2, _ := primitive.ObjectIDFromHex(car.SlotID)
	filter := bson.M{"carnumber": carnumber}
	update := bson.M{"$set": bson.M{"timein": time.Now()}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	var car Car
	err3 := collection.FindOne(context.TODO(), filter).Decode(&car)

	if err3 != nil {
		// if err3 == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return
		// }
		panic(err3)
	}

	//fmt.Printf("%s", car)
	//fmt.Printf(car.SlotID)
	//sid, _ := primitive.ObjectIDFromHex(car.SlotID)
	filter2 := bson.M{"uniqueslotno": car.UniqueSlotNo}
	update2 := bson.M{"$set": bson.M{"occupancy": false}}

	result2, err2 := collection2.UpdateOne(context.TODO(), filter2, update2)

	fmt.Println(result, err)
	fmt.Println(result2, err2)

}

func UpdateCarTimeOut(carnumber string) {

	session := Connect()

	collection := session.Database("webinardb").Collection("cars")
	collection2 := session.Database("webinardb").Collection("slots")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)
	//idPrimitive2, _ := primitive.ObjectIDFromHex(car.SlotID)
	filter := bson.M{"carnumber": carnumber}
	update := bson.M{"$set": bson.M{"timeout": time.Now()}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	var car Car
	err3 := collection.FindOne(context.TODO(), filter).Decode(&car)

	if err3 != nil {
		// if err3 == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return
		// }
		panic(err3)
	}

	//fmt.Printf("%s", car)
	//fmt.Printf(car.SlotID)
	//sid, _ := primitive.ObjectIDFromHex(car.SlotID)
	filter2 := bson.M{"uniqueslotno": car.UniqueSlotNo}
	update2 := bson.M{"$set": bson.M{"occupancy": true}}

	result2, err2 := collection2.UpdateOne(context.TODO(), filter2, update2)

	fmt.Println(result, err)
	fmt.Println(result2, err2)

}

func UpdateSlot(uniqueslotno string, slot Slot) {

	session := Connect()

	collection := session.Database("webinardb").Collection("slots")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"uniqueslotno": uniqueslotno}
	update := bson.M{"$set": bson.M{"occupancy": slot.Occupancy}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	fmt.Println(result, err)

}

func DeleteCar(carnumber string) {

	session := Connect()

	collection := session.Database("webinardb").Collection("cars")
	collection2 := session.Database("webinardb").Collection("slots")
	//collection3 := session.Database("webinardb").Collection("users")

	//car.TimeOut = time.Now()

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)
	//idPrimitive2, _ := primitive.ObjectIDFromHex(car.SlotID)
	//idPrimitive3, _ := primitive.ObjectIDFromHex(car.UserID)
	filter := bson.M{"carnumber": carnumber}
	var car Car
	err := collection.FindOne(context.TODO(), filter).Decode(&car)

	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return
		// }
		panic(err)
	}
	//sid, _ := primitive.ObjectIDFromHex(car.SlotID)
	//uid, _ := primitive.ObjectIDFromHex(car.UserID)
	//fmt.Println(idPrimitive, idPrimitive2)

	filter2 := bson.M{"uniqueslotno": car.UniqueSlotNo}
	update2 := bson.M{"$set": bson.M{"occupancy": true}}

	result2, err2 := collection2.UpdateOne(context.TODO(), filter2, update2)
	//result3, err3 := collection2.UpdateOne(context.TODO(), filter3, update3)
	result, err := collection.DeleteOne(context.TODO(), filter)

	//result3, err3 := collection3.DeleteOne(context.TODO(), filter3)

	fmt.Println(result, err)
	fmt.Println(result2, err2)
	//fmt.Println(result3, err3)

}

func AddSlot(slot Slot) error {

	session := Connect()

	collection := session.Database("webinardb").Collection("slots")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"uniqueslotno": -1},
		Options: options.Index().SetUnique(true),
	}

	name, _ := collection.Indexes().CreateOne(context.TODO(), indexModel)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Name of Index Created: " + name)

	result, err := collection.InsertOne(context.TODO(), slot)

	//slot.ID = result.InsertedID

	//fmt.Println(result, err)
	if err != nil {
		return echo.NewHTTPError(400, "Add a unique slot number")
	}

	log.Println(result)
	log.Println(err)
	//fmt.Println(result2, err2)
	return nil
}

func GetSlot(query url.Values) []Slot {

	var slots []Slot
	isPageQuery := query.Has("page")

	session := Connect()

	// var limit int64 = 4
	// var page int64 = 1
	collection := session.Database("webinardb").Collection("slots")

	if isPageQuery {

		page, _ := strconv.ParseInt(query["page"][0], 10, 64)

		var limit int64 = 5

		_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&slots).Find()

		if err != nil {

			panic(err)

		}

	} else {

		cursor, _ := collection.Find(context.TODO(), bson.M{})

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {

			var slot Slot

			err := cursor.Decode(&slot)

			if err != nil {

				log.Fatal(err)

			}

			slots = append(slots, slot)

		}

	}

	// page, _ := strconv.ParseInt(query["page"][0], 10, 64)

	// var limit int64 = 5

	// if page > 0 {

	// 	_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&slots).Find()

	// 	if err != nil {

	// 		panic(err)

	// 	}
	// }

	// cur, _ := collection.Find(context.TODO(), bson.M{})

	// filter := bson.M{}
	// //projection := bson.M{"floorno": 1, "uniqueslotno": 1}
	// paginatedData, err := New(collection).Limit(limit).Page(page).Sort("floorno", -1).Filter(filter).Decode(&slots).Find()
	// if err != nil {
	// 	panic(err)
	// }

	// // print pagination data
	// fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)

	// defer cur.Close(context.TODO())

	// for cur.Next(context.TODO()) {

	// 	var slot Slot

	// 	err := cur.Decode(&slot)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	slots = append(slots, slot)
	// }
	return slots

}

func GetFreeSlot(query url.Values) []Slot {

	var slots []Slot

	fmt.Println("query1: ", query)

	isPageQuery := query.Has("page")

	session := Connect()

	collection := session.Database("webinardb").Collection("slots")

	if isPageQuery {

		page, _ := strconv.ParseInt(query["page"][0], 10, 64)

		var limit int64 = 5

		_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{"occupancy": true}).Decode(&slots).Find()

		if err != nil {

			panic(err)

		}

	} else {

		cursor, _ := collection.Find(context.TODO(), bson.M{})

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {

			var slot Slot

			err := cursor.Decode(&slot)

			if err != nil {

				log.Fatal(err)

			}

			slots = append(slots, slot)

		}

	}

	// page, _ := strconv.ParseInt(query["page"][0], 10, 64)

	// var limit int64 = 5

	// if page > 0 {
	// 	_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{"occupancy": true}).Decode(&slots).Find()

	// 	if err != nil {

	// 		panic(err)

	// 	}
	// }
	// filter := bson.M{"occupancy": true}
	// cur, _ := collection.Find(context.TODO(), filter)

	// var limit int64 = 4
	// var page int64 = 1

	// paginatedData, err := New(collection).Limit(limit).Page(page).Sort("floorno", -1).Filter(filter).Decode(&slots).Find()
	// if err != nil {
	// 	panic(err)
	// }

	// // print pagination data
	// fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)

	// defer cur.Close(context.TODO())

	// for cur.Next(context.TODO()) {

	// 	var slot Slot

	// 	err := cur.Decode(&slot)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	slots = append(slots, slot)
	// }
	return slots

}

func DeleteSlot(uniqueslotno string) {

	session := Connect()

	collection := session.Database("webinardb").Collection("slots")

	//idPrimitive, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"uniqueslotno": uniqueslotno}
	fmt.Println(filter)
	result, err := collection.DeleteOne(context.TODO(), filter)

	fmt.Println(result, err)

	// if err != nil {
	// 	log.Println(err)

	// }

}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongo_db:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected...")
	return client
}
