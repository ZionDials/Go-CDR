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
	"fmt"
	"strconv"
	"strings"

	"github.com/ziondials/go-cdr/logger"
)

const PhoneNumberSplit = "#:"

func ConvertHexadecimalToInt(s *string) *int64 {
	if s == nil {
		return nil
	}

	integer, err := strconv.ParseInt(strings.TrimSpace(*s), 16, 64)
	if err != nil {
		logger.Error("Error converting hexadecimal to int: %s", err)
		return nil
	}
	return &integer
}

func ConvertStringToInt64(s *string) (*int64, error) {
	if s == nil {
		return nil, nil
	}
	newString := stringToIntReg.ReplaceAllString(*s, "")
	if newString != "" {
		integer, err := strconv.ParseInt(strings.TrimSpace(newString), 10, 64)
		if err != nil {
			logger.Error("Error converting string to int: %s", err)
			return nil, fmt.Errorf("error converting string to int: %s", err)
		}
		return &integer, nil
	} else {
		return nil, nil
	}
}

func ConvertStringToInt(s *string) (*int, error) {
	if s == nil {
		return nil, nil
	}
	newString := stringToIntReg.ReplaceAllString(*s, "")
	if newString != "" {
		integer, err := strconv.Atoi(strings.TrimSpace(newString))
		if err != nil {
			logger.Error("Error converting string to int: %s", err)
			return nil, fmt.Errorf("error converting string to int: %s", err)
		}
		return &integer, nil
	} else {
		return nil, nil
	}
}

func ExtractPhoneNumberFromString(s *string) *string {
	if s == nil {
		return nil
	}
	if strings.Contains(*s, PhoneNumberSplit) {
		SplitString := strings.Split(*s, PhoneNumberSplit)[1]

		return &SplitString
	} else {
		return nil
	}
}

func ConvertStringToFloat64(s *string) (*float64, error) {
	if s == nil {
		return nil, nil
	}
	newString := stringToIntReg.ReplaceAllString(*s, "")
	if newString != "" {
		integer, err := strconv.ParseFloat(strings.TrimSpace(newString), 64)
		if err != nil {
			logger.Error("Error converting string to int: %s", err)
			return nil, fmt.Errorf("error converting string to int: %s", err)
		}
		return &integer, nil
	} else {
		return nil, nil
	}
}
