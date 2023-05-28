// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package helpers

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidNTPReferenceAsterisk = errors.New("device has invalid ntp reference as indicated by asterisk (*) prepended to time")
	ErrInvalidNTPReferencePeriod   = errors.New("device has invalid ntp reference as indicated by period (.) prepended to time")
	ErrInvalidTimeFormat           = errors.New("time format is invalid")
)

func ConvertStringToUnixTime(s *string, t *time.Location) (*int64, error) {
	if s == nil || *s == "" {
		return nil, nil
	}
	if unixTimePrimaryReg.MatchString(*s) {
		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime stdReg MatchString: %s", err)
		}
		parsedTime := t.UTC().Unix()
		return &parsedTime, nil
	}
	if unixTimeSecondaryReg.MatchString(*s) {
		t, err := time.Parse("01/02/2006 15:04:05.000", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime secondaryReg MatchString: %s", err)
		}
		parsedTime := t.UTC().Unix()
		return &parsedTime, nil
	}
	if unixTimeTertiaryReg.MatchString(*s) {
		t, err := time.Parse("01 02 2006 15 04 05.000", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime tertiaryReg MatchString: %s", err)
		}
		parsedTime := t.UTC().Unix()
		return &parsedTime, nil
	}
	if unixTimeQuaternaryReg.MatchString(*s) {
		runes := []rune(*s)
		stringLength := len(*s)
		substr := string(runes[1:stringLength])
		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", substr)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime quaternaryReg MatchString: %s", err)
		}
		parsedTime := t.UTC().Unix()
		return &parsedTime, ErrInvalidNTPReferenceAsterisk
	}
	if unixTimeQuinaryReg.MatchString(*s) {
		runes := []rune(*s)
		stringLength := len(*s)
		substr := string(runes[1:stringLength])
		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", substr)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime quinaryReg MatchString: %s", err)
		}
		parsedTime := t.UTC().Unix()
		return &parsedTime, ErrInvalidNTPReferencePeriod
	}
	return nil, ErrInvalidTimeFormat

}

func ExtractTimeLocationFromString(s *string) (*time.Location, error) {
	if s == nil || *s == "" {
		return nil, nil
	}
	if extractTimeLocationPrimaryReg.MatchString(*s) {

		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime stdReg MatchString: %s", err)
		}
		return t.Location(), nil
	}
	if extractTimeLocationSecondaryReg.MatchString(*s) {
		t, err := time.Parse("01/02/2006 15:04:05.000", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime secondaryReg MatchString: %s", err)
		}
		return t.Location(), nil
	}
	if extractTimeLocationTertiaryReg.MatchString(*s) {
		t, err := time.Parse("01 02 2006 15 04 05.000", *s)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime tertiaryReg MatchString: %s", err)
		}
		return t.Location(), nil
	}
	if extractTimeLocationQuaternaryReg.MatchString(*s) {
		runes := []rune(*s)
		stringLength := len(*s)
		substr := string(runes[1:stringLength])
		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", substr)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime quaternaryReg MatchString: %s", err)
		}
		return t.Location(), ErrInvalidNTPReferenceAsterisk
	}
	if extractTimeLocationQuinaryReg.MatchString(*s) {
		runes := []rune(*s)
		stringLength := len(*s)
		substr := string(runes[1:stringLength])
		t, err := time.Parse("15:04:05.000 MST Mon Jan 2 2006", substr)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringToUnixTime quinaryReg MatchString: %s", err)
		}
		return t.Location(), ErrInvalidNTPReferencePeriod
	}
	return nil, ErrInvalidTimeFormat
}
