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

import "strings"

// Take String, Delimiter, and Equality Charachter and return a map of key value pairs
func ConvertStringToKeyValuePairs(s *string, d string, e string) map[string]string {
	m := make(map[string]string)
	// Split String by Delimiter
	split := strings.Split(*s, d)
	for _, v := range split {
		// Split String by Equality Charachter
		split2 := strings.Split(v, e)
		// If the split2 array has a length of 2, add the key value pair to the map
		if len(split2) == 2 {
			m[split2[0]] = split2[1]
		}
	}
	return m
}
