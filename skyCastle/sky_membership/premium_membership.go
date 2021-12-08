package sky_membership

import "errors"

// SKY_Membership struct
type Premium_membership struct {
	name string
	gender string
	age int
	asset int
}
var ErrNotEnoughAsset = errors.New(
	"You can not withdraw this amount because asset is not enough",
)

func Register(_name, _gender string, _age int) *Premium_membership {
	new_membership := Premium_membership{
		name: _name,
		gender: _gender,
		age: _age,
		asset: 0,
	}
	return &new_membership // 실제 new_membership 주소를 가리키는 것
}

func (pm *Premium_membership) Sponsor(_amount int) {
	pm.asset += _amount
}

func (pm Premium_membership) CheckAsset() int {
	return pm.asset;
}

func (pm *Premium_membership) WithdrawAsset(_amount int) error {
switch {
	case pm.asset < _amount: 
		return ErrNotEnoughAsset
	case pm.asset >= _amount:
		pm.asset -= _amount
}
	
	return nil
}