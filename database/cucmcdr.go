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

package database

import "github.com/ziondials/go-cdr/models"

func (ds DataService) CreateCucmCDRs(cdrs []*models.CucmCdr) error {

	if rsp := ds.Session.CreateInBatches(&cdrs, int(ds.Config.Limit)); rsp.Error != nil {
		return rsp.Error
	}

	return nil
}
