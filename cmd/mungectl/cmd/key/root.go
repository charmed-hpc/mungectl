// Copyright 2024 Canonical Ltd.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package key

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const keyHelp = "Manage munge key file"

var KeyCmd = &cobra.Command{
	Use:   "key",
	Short: keyHelp,
}

func init() {
	KeyCmd.CompletionOptions.DisableDefaultCmd = true
	KeyCmd.AddCommand(generateCmd, getCmd, setCmd)
	KeyCmd.PersistentFlags().StringP("keyfile", "k", "", "Specify keyfile path")
	// errcheck: an error would be returned here if the keyfile flag did not exist,
	// but since the flag does exist, we can safely ignore the error returned here.
	_ = viper.BindPFlag("keyfile", KeyCmd.PersistentFlags().Lookup("keyfile"))
}
