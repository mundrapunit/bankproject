//a struct to rep user account

package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	u "../utils"
	"github.com/gorilla/mux"
)

/*
type Bank struct {
	gorm.Model
	Id    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token";sql:"-"`
}
*/

// Branch : model : all the feild
type Branch struct {
	IFSC     string `json:"ifsc"`
	BankID   int    `json:"bank_id"`
	Branch   string `json:"branch"`
	Address  string `json:"address"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Name     string `json:"bank_name"`
}

// GetBranchDetails : fetxh the brnach details
func getBranchDetail(ifsc string) *Branch {
	br := &Branch{}
	GetDB().Table("branches").Where("ifsc = ?", ifsc).First(br)
	fmt.Println(br)
	return br
}

// GetBranchDetail : clouser for branch details
var GetBranchDetail = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ifsc := vars["ifsc"]
	data := getBranchDetail(ifsc)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetBranchDetails : fetxh the branch details
func getBranchDetails(branch string, city string, limit int, offset int) []*Branch {
	//br := &Branch{}
	br := make([]*Branch, 0)
	if limit > 0 {
		GetDB().Table("bank_branches_tem").Where("branch = ? AND city = ?", branch, city).Limit(limit).Offset(offset).Find(&br)
	} else {
		GetDB().Table("bank_branches").Where("branch = ? AND city = ?", branch, city).Find(&br)
	}
	fmt.Println("br:", br)

	return br
}

// Throtel : use to limit and offset
type Throtel struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// GetBranchDetails : clouser for branch details
var GetBranchDetails = func(w http.ResponseWriter, r *http.Request) {

	throtel := &Throtel{}
	err := json.NewDecoder(r.Body).Decode(throtel)
	if err != nil {
		fmt.Println("Error:", err)
		// TODO : neet to return
	}
	fmt.Println("Let set what come :", throtel)
	vars := mux.Vars(r)
	branch := vars["branch"]
	city := vars["city"]
	limit := throtel.Limit
	offset := throtel.Offset
	fmt.Println("branch:", branch)
	fmt.Println("city:", city)
	fmt.Println("limit:", limit)
	fmt.Println("offset:", offset)

	data := getBranchDetails(branch, city, limit, offset)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
