package main

import (
	"2021_hancom_lecture/skyCastle/sky_membership"
	"fmt"
)

func main() {
// Register a member
new_member := sky_membership.Register("Kim", "Male", 30)
fmt.Println("main => ", new_member)

// Sponsor
new_member.Sponsor(1000)

currentAsset := new_member.CheckAsset()
fmt.Println(currentAsset)

// Withdraw asset
err := new_member.WithdrawAsset(2000)
if err != nil {
	fmt.Println(err)
}

fmt.Println(new_member)

	// private
// 	new_member := sky_membership.Premium_membership{
// 	name: "Lee",
// 	gender: "Female",
// 	age: 40,
// 	Asset: 10000,
// }

// public
// new_member2 := sky_membership.Premium_membership {
// 	Name: "Lee",
// 	Gender: "Female",
// 	Age: 40,
// 	Asset: 10000,
// }
// fmt.Println(new_member)
// fmt.Println(new_member2)
}