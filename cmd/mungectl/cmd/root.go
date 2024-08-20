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

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	key "github.com/charmed-hpc/mungectl/cmd/mungectl/cmd/key"
)

const help = "Control the munge authentication daemon"

var rootCmd = &cobra.Command{
	Use:   "mungectl",
	Short: help,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(key.KeyCmd)
}

// Initialize mungectl configuration.
func initConfig() {
	viper.SetConfigName("mungectl")
	viper.SetConfigType("yaml")

	if os.Getenv("SNAP") != "" {
		common := os.Getenv("SNAP_COMMON")
		viper.AddConfigPath(path.Join(common, "/etc/munge/mungectl.yaml"))
		viper.SetDefault("keyfile", path.Join(common, "/etc/munge/munge.key"))
	} else {
		viper.AddConfigPath("/etc/munge/mungectl.yaml")
		viper.SetDefault("keyfile", "/etc/munge/munge.key")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
