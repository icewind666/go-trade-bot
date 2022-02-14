package models

import "trading-bot/src/main/investapi"

// UserInfo holds user permissions on broker system
type UserInfo struct {
	CanWorkWithBond          bool
	CanWorkWithForeignShares bool
	CanWorkWithRussianShares bool
	CanWorkWithForeignEtf    bool
	CanWorkWithLeverage      bool
	IsQualified              bool
}

func MapFrom(response *investapi.GetInfoResponse) *UserInfo {
	return &UserInfo{
		CanWorkWithBond:          IsInArray("bond", response.QualifiedForWorkWith),
		CanWorkWithForeignShares: IsInArray("foreign_shares", response.QualifiedForWorkWith),
		CanWorkWithRussianShares: IsInArray("russian_shares", response.QualifiedForWorkWith),
		CanWorkWithForeignEtf:    IsInArray("foreign_etf", response.QualifiedForWorkWith),
		CanWorkWithLeverage:      IsInArray("leverage", response.QualifiedForWorkWith),
	}
}

func IsInArray(searchedValue string, values []string) bool {
	for _, x := range values {
		if x == searchedValue {
			return true
		}
	}
	return false
}
