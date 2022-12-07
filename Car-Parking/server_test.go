package main

import (
	"net/http"
	"net/http/httptest"
	_ "strconv"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	body := strings.NewReader(`{
		"Fname": "TestData1",
		"LName": "TestData1",
		"UserID": "0",
	}`)

	req := httptest.NewRequest("POST", "/user", body)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetUser(t *testing.T) {

	req := httptest.NewRequest("GET", "/user?page=1", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}
func TestUpdateUser(t *testing.T) {
	body := strings.NewReader(`{
		"fname": "TestData11",
		"lname": "TestData11"
	}`)

	req := httptest.NewRequest("PUT", "/user/0", body)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/user/0", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestAddCar(t *testing.T) {

	body := strings.NewReader(`{
		"carnumber": "UK12FG5643",
		"CarModel": "Kia",
		"UserID": "0",
		""UniqueSlotNo" :"-3A0"
	}`)

	req := httptest.NewRequest("POST", "/car", body)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddCar(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetCar(t *testing.T) {
	req := httptest.NewRequest("GET", "/car?page=1", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCar(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetCarByNumber(t *testing.T) {
	req := httptest.NewRequest("GET", "/car/UK12FG5643", nil)
	recorder := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				w: recorder,
				r: req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCarByNumber(tt.args.w, tt.args.r)
		})
	}
}

// func TestGetCarByNumber(t *testing.T) {

// 	req := httptest.NewRequest("GET", "/car/UK12FG5643", nil)
// 	recorder := httptest.NewRecorder()

// 	type args struct {
// 		writer http.ResponseWriter
// 		req    *http.Request
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "Test1",
// 			args: args{
// 				writer: recorder,
// 				req:    req,
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			GetCarByNumber(tt.args.writer, tt.args.req)
// 			res := recorder.Result()
// 			if res.StatusCode != http.StatusOK {
// 				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
// 			}
// 		})
// 	}

// }

func TestUpdateCar(t *testing.T) {
	body := strings.NewReader(`{
		"carnumber": "UK12FG5643",
		"CarModel": "Kia Seltos",
	}`)

	req := httptest.NewRequest("PUT", "/car/UK12FG5643", body)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCar(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestUpdateCarTimeIn(t *testing.T) {

	req := httptest.NewRequest("PUT", "/car/UK12FG5643", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCarTimeIn(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestUpdateCarTimeOut(t *testing.T) {

	req := httptest.NewRequest("PUT", "/car/UK12FG5643", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCarTimeOut(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteCar(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/car/UK12FG5643", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteCar(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestAddSlot(t *testing.T) {
	body := strings.NewReader(`{
		"floorno": 4,
		"uniqueslotno": "4A4",
		"occupancy": true,
	}`)

	req := httptest.NewRequest("POST", "/slot", body)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetSlot(t *testing.T) {
	req := httptest.NewRequest("GET", "/slot?page=1", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetFreeSlot(t *testing.T) {
	req := httptest.NewRequest("GET", "/slot/free?page=1", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetFreeSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestUpdateSlot(t *testing.T) {
	body := strings.NewReader(`{
		"occupancy": false,
	}`)

	req := httptest.NewRequest("PUT", "/slot/4A4", body)
	recorder := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{
			name: "Test1",
			args: args{
				w: recorder,
				r: req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateSlot(tt.args.w, tt.args.r)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteSlot(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/slot/4A4", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestWrongRoute(t *testing.T) {
	req := httptest.NewRequest("GET", "/wrongpath", nil)
	recorder := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.\
		{
			name: "Test Case 1",
			args: args{
				w: recorder,
				r: req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WrongRoute(tt.args.w, tt.args.r)
		})
	}
}
