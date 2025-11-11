// Copyright (C) Kumo inc. and its affiliates.
// Author: Jeff.li lijippy@163.com
// All rights reserved.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package kprinter

import (
	"fmt"

	"github.com/pterm/pterm"
)

// Green returns green colored text
func Yellow(text string) string {
	return pterm.Yellow(text)
}

// Green returns green colored text
func Yellowf(text string, args ...any) string {
	return pterm.Yellow(fmt.Sprintf(text, args...))
}

// Green returns green colored text
func Green(text string) string {
	return pterm.Green(text)
}

// Green returns green colored text
func Greenf(text string, args ...any) string {
	return pterm.Green(fmt.Sprintf(text, args...))
}

// Green returns green colored text
func Magenta(text string) string {
	return pterm.Magenta(text)
}

// Green returns green colored text
func Magentaf(text string, args ...any) string {
	return pterm.Magenta(fmt.Sprintf(text, args...))
}

// Red returns red colored text
func Red(text string) string {
	return pterm.Red(text)
}

// Red returns red colored text
func Redf(text string, args ...any) string {
	return pterm.Red(fmt.Sprintf(text, args...))
}

// Green returns green colored text
func Blue(text string) string {
	return pterm.Blue(text)
}

// Green returns green colored text
func Bluef(text string, args ...any) string {
	return pterm.Blue(fmt.Sprintf(text, args...))
}

// RedOnWhite returns red text on white background
func RedOnWhitef(text string, args ...any) string {
	return pterm.NewStyle(pterm.FgRed, pterm.BgWhite).Sprint(fmt.Sprintf(text, args...))
}

// RedOnWhite returns red text on white background
func RedOnWhite(text string) string {
	return pterm.NewStyle(pterm.FgRed, pterm.BgWhite).Sprint(text)
}

// Bold returns bold text
func Bold(text string, args ...any) string {
	return pterm.Bold.Sprint(fmt.Sprintf(text, args...))
}
