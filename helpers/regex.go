package helpers

import "regexp"

var (
	// ConvertStringToUnixTime
	unixTimePrimaryReg    = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)
	unixTimeSecondaryReg  = regexp.MustCompile(`^\d{2}\/\d{2}\/\d{4}\s\d{2}:\d{2}:\d{2}.\d{3}`)
	unixTimeTertiaryReg   = regexp.MustCompile(`^\d{2}\s\d{2}\s\d{4}\s\d{2}\s\d{2}\s\d{2}.\d{3}`)
	unixTimeQuaternaryReg = regexp.MustCompile(`^\*\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)
	unixTimeQuinaryReg    = regexp.MustCompile(`^\.\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)

	// ExtractTimeLocationFromString
	extractTimeLocationPrimaryReg    = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)
	extractTimeLocationSecondaryReg  = regexp.MustCompile(`^\d{2}\/\d{2}\/\d{4}\s\d{2}:\d{2}:\d{2}.\d{3}`)
	extractTimeLocationTertiaryReg   = regexp.MustCompile(`^\d{2}\s\d{2}\s\d{4}\s\d{2}\s\d{2}\s\d{2}.\d{3}`)
	extractTimeLocationQuaternaryReg = regexp.MustCompile(`^\*\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)
	extractTimeLocationQuinaryReg    = regexp.MustCompile(`^\.\d{2}:\d{2}:\d{2}.\d{3}\s[A-Z]{3}\s[a-zA-Z]{3}\s[a-zA-Z]{3}\s\d{1,2}\s\d{4}`)

	// ConvertStringToInt64
	stringToIntReg = regexp.MustCompile("[^0-9]+")
)
