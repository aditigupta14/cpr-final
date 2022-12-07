package DAL

import (
	"net/url"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "adding user",
			args: args{
				user: User{FName: "Riya", LName: "Mehta", UserID: "0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUser(tt.args.user)
		})
	}
}

func TestGetUser(t *testing.T) {

	type args struct {
		query url.Values
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		// TODO: Add test cases.
		{
			name: "getting user",
			args: args{
				query: url.Values{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.query.Set("page", "1000")
			if got := GetUser(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGetUser(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []User
// 	}{
// 		// TODO: Add test cases.

// 		{
// 			name: "getting user data",
// 			want: []User{{FName: "Aditi", LName: "Gupta", UserID: "2"}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		var user User
// 		user.FName = "Aditi"
// 		user.LName = "Gupta"
// 		user.UserID = "2"
// 		users := []User{user}
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := users; !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestUpdateUserByID(t *testing.T) {
	type args struct {
		userid string
		user   User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updating user data",
			args: args{
				user:   User{FName: "Priya", LName: "Gupta"},
				userid: "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUserByID(tt.args.userid, tt.args.user)
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	type args struct {
		userid string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{

			name: "deleting user data",
			args: args{
				userid: "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUserByID(tt.args.userid)
		})
	}
}

func TestAddCar(t *testing.T) {
	type args struct {
		car Car
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "adding car",
			args: args{
				car: Car{CarNumber: "UK23FG9234", CarModel: "Nano", UserID: "0", UniqueSlotNo: "-3A0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddCar(tt.args.car)
		})
	}
}

func TestGetCar(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name string
		args args
		want []Car
	}{
		// TODO: Add test cases.
		{
			name: "getting car",
			args: args{
				query: url.Values{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.query.Set("page", "1000")
			if got := GetCar(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGetCar(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []Car
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "getting car data",
// 			want: []Car{{CarNumber: "UK23FG9234", CarModel: "Nano", UserID: "0", UniqueSlotNo: "-3A0"}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		var car Car
// 		car.CarNumber = "UK23FG9234"
// 		car.CarModel = "Nano"
// 		car.UserID = "0"
// 		car.UniqueSlotNo = "-3A0"
// 		cars := []Car{car}
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := cars; !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetCar() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetCarByNumber(t *testing.T) {
	type args struct {
		carnumber string
	}
	tests := []struct {
		name string
		args args
		want Car
	}{
		// TODO: Add test cases.
		{
			name: "getting car by number",
			args: args{
				carnumber: "UK23FG9234",
			},
			want: Car{CarNumber: "UK23FG9234", CarModel: "Nano", UserID: "0", UniqueSlotNo: "-3A0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCarByNumber(tt.args.carnumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCarByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGetCarByNumber(t *testing.T) {
// 	type args struct {
// 		carnumber string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want Car
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "getting car by its number",
// 			args: args{
// 				carnumber: "UK23FG9234",
// 			},

// 			want: Car{CarNumber: "UK23FG9234", CarModel: "Nano", UserID: "0", UniqueSlotNo: "-3A0"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var car Car
// 			car.CarNumber = tt.args.carnumber
// 			car.CarModel = "Nano"
// 			car.UserID = "0"
// 			car.UniqueSlotNo = "-3A0"
// 			if got := car; !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetCarByNumber() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestUpdateCar(t *testing.T) {
	type args struct {
		carnumber string
		car       Car
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updating car data",
			args: args{

				car:       Car{CarNumber: "UK23FG9234", CarModel: "Kia"},
				carnumber: "UK23FG9234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCar(tt.args.carnumber, tt.args.car)
		})
	}
}

func TestUpdateCarTimeIn(t *testing.T) {
	type args struct {
		carnumber string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updating car timein",
			args: args{
				carnumber: "UK23FG9234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCarTimeIn(tt.args.carnumber)
		})
	}
}

func TestUpdateCarTimeOut(t *testing.T) {
	type args struct {
		carnumber string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updating car timeout",
			args: args{
				carnumber: "UK23FG9234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCarTimeOut(tt.args.carnumber)
		})
	}
}

func TestDeleteCar(t *testing.T) {
	type args struct {
		carnumber string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{

			name: "deleting car data",
			args: args{
				carnumber: "UK23FG9234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteCar(tt.args.carnumber)
		})
	}
}

func TestAddSlot(t *testing.T) {
	type args struct {
		slot Slot
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "adding user",
			args: args{
				slot: Slot{FloorNo: -3, UniqueSlotNo: "-3A0", Occupancy: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddSlot(tt.args.slot)
		})
	}
}

func TestGetSlot(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name string
		args args
		want []Slot
	}{
		// TODO: Add test cases.
		{
			name: "getting slot",
			args: args{
				query: url.Values{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.query.Set("page", "1000")
			if got := GetSlot(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFreeSlot(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name string
		args args
		want []Slot
	}{
		// TODO: Add test cases.
		{
			name: "getting free slot",
			args: args{
				query: url.Values{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.query.Set("page", "1000")
			if got := GetFreeSlot(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}

}

// func TestGetSlot(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []Slot
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "getting slot data",
// 			want: []Slot{{FloorNo: -3, UniqueSlotNo: "-3A0", Occupancy: false}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		var slot Slot
// 		slot.FloorNo = -3
// 		slot.UniqueSlotNo = "-3A0"
// 		slot.Occupancy = false

// 		slots := []Slot{slot}
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := slots; !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetSlot() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetFreeSlot(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []Slot
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "getting free slot data",
// 			want: []Slot{{FloorNo: -3, UniqueSlotNo: "-3A0", Occupancy: true}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		var slot Slot
// 		slot.FloorNo = -3
// 		slot.UniqueSlotNo = "-3A0"
// 		slot.Occupancy = true

// 		slots := []Slot{slot}
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := slots; !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetSlot() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestUpdateSlot(t *testing.T) {
	type args struct {
		uniqueslotno string
		slot         Slot
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updating slot data",
			args: args{
				slot:         Slot{Occupancy: false},
				uniqueslotno: "4A4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateSlot(tt.args.uniqueslotno, tt.args.slot)
		})
	}
}

func TestDeleteSlot(t *testing.T) {
	type args struct {
		uniqueslotno string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{

			name: "deleting slot data",
			args: args{
				uniqueslotno: "-3A0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteSlot(tt.args.uniqueslotno)
		})
	}
}
