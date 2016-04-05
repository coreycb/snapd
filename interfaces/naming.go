// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package interfaces

import (
	"fmt"

	"github.com/ubuntu-core/snappy/snap"
)

// SecurityTag returns application-specific security tag.
//
// Security tags are used by various security subsystems as "profile names" and
// sometimes also as a part of the file name.
func SecurityTag(appInfo *snap.AppInfo) string {
	return fmt.Sprintf("snap.%s.%s", appInfo.Snap.Name, appInfo.Name)
}

// SecurityTagGlob returns a pattern that matches all security tags belonging to
// the same snap as the given app.
func SecurityTagGlob(snapInfo *snap.Info) string {
	return fmt.Sprintf("snap.%s.%s", snapInfo.Name, "*")
}