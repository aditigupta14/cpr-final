package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	_ "strconv"
	"strings"

	"Parking.com/DAL"
	"github.com/gorilla/mux"
)

func EnableCors(w *http.ResponseWriter) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")

}

// swagger:operation POST /user addUser
//
// # Include documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/User"
//
// responses:
//   '200':
//     description: user response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/User"

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user DAL.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)

	fmt.Println(err)

	fmt.Println(user)

	val := DAL.AddUser(user)

	//w.WriteHeader(http.StatusCreated) // 201
	EnableCors(&w)
	if val == nil {
		respondWithJSON(w, http.StatusOK, "User Added")
	} else {
		respondWithJSON(w, http.StatusBadRequest, val)
	}
}

// swagger:operation GET /user getUser
// // # Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: page
//     in: query
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: user response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/User"

func GetUser(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	query := r.URL.Query()
	isPageQuery := query.Has("page")

	//fmt.Println(query)
	//if reflect.ValueOf(query).Kind() == reflect.Map {
	users := DAL.GetUser(query)
	if users != nil {
		respondWithJSON(w, http.StatusOK, users)
		if isPageQuery {

			temp, _ := strconv.Atoi(query["page"][0])

			temp += 1

			currPage := strconv.Itoa(temp)

			nextPage := strings.Replace("/user?page=0", "0", currPage, 12)
			fmt.Fprintln(w, "\nNext Page: ", nextPage)

		}

	} else {
		http.Error(w, "Invalid Page number", 400)
	}

	// } else {
	// 	http.Error(w, "Enter page number", 400)
	// }
	//convert golang object to json and sending to Response object
	//json.NewEncoder(w).Encode(users)

}

// swagger:operation PUT /user/{userid} updateUser
//
// # Insert documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: userid
//     in: path
//     description: User id that updates a user
//     required: true
//     type: string
//
//   - name: body
//     in: body
//     schema:
//			 "$ref": "#/definitions/User"
//
// responses:
//   '200':
//     description: user response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/User"

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	userid := params["userid"]

	var user DAL.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)

	fmt.Println(err)

	fmt.Println(user)

	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	// ... handle error
	// 	panic(err)
	// }

	DAL.UpdateUserByID(userid, user)

}

// swagger:operation DELETE /user/{userid} deleteUser
//
// # Insert documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: userid
//     in: path
//     description: User id that deletes a user
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: user response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/User"

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	userid := params["userid"]
	fmt.Println("userid is", userid)

	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	// ... handle error
	// 	panic(err)
	// }

	DAL.DeleteUserByID(userid)

}

// swagger:operation POST /car addCar
//
// # Include documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/Car"
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func AddCar(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	var car DAL.Car

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&car)

	fmt.Println(err)

	fmt.Println(car)

	val := DAL.AddCar(car)

	if val == nil {
		respondWithJSON(w, http.StatusOK, "Car Added")
	} else {
		respondWithJSON(w, http.StatusBadRequest, val)
		//w.WriteHeader(http.StatusCreated) // 201
	}
}

// swagger:operation GET /car getCar
//
// # Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: page
//     in: query
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func GetCar(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	query := r.URL.Query()
	isPageQuery := query.Has("page")
	cars := DAL.GetCar(query)

	//convert golang object to json and sending to Response object
	//json.NewEncoder(w).Encode(cars)
	if cars != nil {
		respondWithJSON(w, http.StatusOK, cars)
		if isPageQuery {

			temp, _ := strconv.Atoi(query["page"][0])

			temp += 1

			currPage := strconv.Itoa(temp)

			nextPage := strings.Replace("/car?page=0", "0", currPage, 12)
			fmt.Fprintln(w, "\nNext Page: ", nextPage)

		}

	} else {
		http.Error(w, "Invalid Page number", 400)
	}

}

// swagger:operation GET /car/{carnumber} getCarbyNumber
//
// Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
// - name: carnumber
//   in: path
//   description: Car number that need to be returned
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func GetCarByNumber(w http.ResponseWriter, r *http.Request) {
	//Reading/Writing Header Values
	// w.Header().Set("Content-Type", "application/json")

	// header1 := r.Header["Content-Type"]
	// header2 := r.Header["Token"]

	// fmt.Println(header1, header2)

	//Reading Query string values
	//hostname?carmodel=abc123&carnumber=abc

	//param1 := r.URL.Query().Get("param1")

	//io.WriteString(w, "GetCars is called....")

	//Get request data from URL route parameter

	params := mux.Vars(r)

	carnumber := params["carnumber"]

	car := DAL.GetCarByNumber(carnumber)

	if car == (DAL.Car{}) {

		//http.Error(w, "Wrong Car number+++++++++++++", 400)
		respondWithJSON(w, http.StatusBadRequest, "Wrong Car Number")

		return
	} else {
		//json.NewEncoder(w).Encode(car)
		respondWithJSON(w, http.StatusOK, car)
	}
	EnableCors(&w)
}

// swagger:operation PUT /car/{carnumber} updateCar
//
// # Include documentation
//
// ---
// produces:
// - application/json
// parameters:
//
//   - name: carnumber
//     in: path
//     description: Car Number to update a car
//     required: true
//     type: string
//
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/Car"
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	carnumber := params["carnumber"]

	var car DAL.Car

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&car)

	fmt.Println(err)

	fmt.Println(car)

	DAL.UpdateCar(carnumber, car)

}

// swagger:operation PUT /cartimein/{carnumber} updateTimeInCar
//
// # Include documentation
//
// ---
// produces:
// - application/json
// parameters:
//
//   - name: carnumber
//     in: path
//     description: Car Number to update car timein
//     required: true
//     type: string
//
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/Car"
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func UpdateCarTimeIn(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	carnumber := params["carnumber"]

	DAL.UpdateCarTimeIn(carnumber)

}

// swagger:operation PUT /cartimeout/{carnumber} updateTimeoutCar
//
// # Include documentation
//
// ---
// produces:
// - application/json
// parameters:
//
//   - name: carnumber
//     in: path
//     description: Car Number to update car timeout
//     required: true
//     type: string
//
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/Car"
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func UpdateCarTimeOut(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	carnumber := params["carnumber"]

	DAL.UpdateCarTimeOut(carnumber)

}

// swagger:operation DELETE /car/{carnumber} deleteCar
//
// # Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: carnumber
//     in: path
//     description: Carnumber to delete a car
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: car response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Car"

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	carnumber := params["carnumber"]

	DAL.DeleteCar(carnumber)

}

// swagger:operation POST /slot addSlot
//
// # Include documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: Body
//     in: body
//     schema:
//     		"$ref": "#/definitions/Slot"
//
// responses:
//   '200':
//     description: slot response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Slot"

func AddSlot(w http.ResponseWriter, r *http.Request) {
	var slot DAL.Slot

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&slot)

	fmt.Println(err)

	fmt.Println(slot)

	val := DAL.AddSlot(slot)
	EnableCors(&w)
	if val == nil {
		respondWithJSON(w, http.StatusOK, "Slot Added")
	} else {
		respondWithJSON(w, http.StatusBadRequest, val)
	}
	//w.WriteHeader(http.StatusCreated) // 201
}

// swagger:operation GET /slot getSlot
//
// Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: page
//     in: query
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: slot response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Slot"

func GetSlot(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	query := r.URL.Query()
	isPageQuery := query.Has("page")
	slots := DAL.GetSlot(query)

	//convert golang object to json and sending to Response object
	//json.NewEncoder(w).Encode(slots)
	if slots != nil {
		respondWithJSON(w, http.StatusOK, slots)
		if isPageQuery {

			temp, _ := strconv.Atoi(query["page"][0])

			temp += 1

			currPage := strconv.Itoa(temp)

			nextPage := strings.Replace("/slot?page=0", "0", currPage, 12)
			fmt.Fprintln(w, "\nNext Page: ", nextPage)

		}

	} else {
		http.Error(w, "Invalid Page number", 400)
	}

}

// swagger:operation GET /slot/free getFreeSlot
//
// # Insert documentation
//
// ---
// produces:
// - application/json
//
// parameters:
//   - name: page
//     in: query
//     required: true
//     type: string
//
//
// responses:
//   '200':
//     description: slot response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Slot"

func GetFreeSlot(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	query := r.URL.Query()
	fmt.Println("query: ", query)
	isPageQuery := query.Has("page")
	slots := DAL.GetFreeSlot(query)

	//convert golang object to json and sending to Response object
	//json.NewEncoder(w).Encode(slots)
	if slots != nil {
		respondWithJSON(w, http.StatusOK, slots)
		if isPageQuery {

			temp, _ := strconv.Atoi(query["page"][0])

			temp += 1

			currPage := strconv.Itoa(temp)

			nextPage := strings.Replace("/slot/free?page=0", "0", currPage, 12)
			fmt.Fprintln(w, "\nNext Page: ", nextPage)

		}

	} else {
		http.Error(w, "Invalid Page number", 400)
	}

}

// swagger:operation PUT /slot/{uniqueslotno} updateSlot
//
// # Insert documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: uniqueslotno
//     in: path
//     description: Unique slot number to update a slot
//     required: true
//     type: string
//
//   - name: body
//     in: body
//     schema:
//			 "$ref": "#/definitions/Slot"
//
// responses:
//   '200':
//     description: slot response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Slot"

func UpdateSlot(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	uniqueslotno := params["uniqueslotno"]

	var slot DAL.Slot

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&slot)

	fmt.Println(err)

	fmt.Println(slot)

	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	// ... handle error
	// 	panic(err)
	// }

	DAL.UpdateSlot(uniqueslotno, slot)

}

// swagger:operation DELETE /slot/{uniqueslotno} deleteSlot
//
// # Insert documentation
//
// ---
// produces:
// - application/json
// parameters:
//   - name: uniqueslotno
//     in: path
//     description: Unique slot number that deletes a slot
//     required: true
//     type: string
//
// responses:
//   '200':
//     description: slot response
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Slot"

func DeleteSlot(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	uniqueslotno := params["uniqueslotno"]
	fmt.Println("usn is:", uniqueslotno)
	DAL.DeleteSlot(uniqueslotno)

}

func WrongRoute(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	respondWithJSON(w, http.StatusBadRequest, "Invalid path")

}

func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func main() {

	Router := mux.NewRouter()

	Router.HandleFunc("/user", AddUser).Methods("POST")
	Router.HandleFunc("/user", GetUser).Methods("GET")
	Router.HandleFunc("/user/{userid}", UpdateUser).Methods("PUT")
	Router.HandleFunc("/user/{userid}", DeleteUser).Methods("DELETE")

	Router.HandleFunc("/car", AddCar).Methods("POST")
	Router.HandleFunc("/car", GetCar).Methods("GET")
	Router.HandleFunc("/car/{carnumber}", GetCarByNumber).Methods("GET")
	Router.HandleFunc("/car/{carnumber}", UpdateCar).Methods("PUT")
	Router.HandleFunc("/cartimein/{carnumber}", UpdateCarTimeIn).Methods("PUT")
	Router.HandleFunc("/cartimeout/{carnumber}", UpdateCarTimeOut).Methods("PUT")
	Router.HandleFunc("/car/{carnumber}", DeleteCar).Methods("DELETE")

	Router.HandleFunc("/slot", AddSlot).Methods("POST")
	Router.HandleFunc("/slot", GetSlot).Methods("GET")
	Router.HandleFunc("/slot/free", GetFreeSlot).Methods("GET")
	Router.HandleFunc("/slot/{uniqueslotno}", UpdateSlot).Methods("PUT")
	Router.HandleFunc("/slot/{uniqueslotno}", DeleteSlot).Methods("DELETE")

	Router.HandleFunc("*", WrongRoute).Methods("POST")
	Router.HandleFunc("*", WrongRoute).Methods("GET")
	Router.HandleFunc("*", WrongRoute).Methods("PUT")
	Router.HandleFunc("*", WrongRoute).Methods("DELETE")

	http.ListenAndServe(":8080", Router)

}
