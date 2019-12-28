/*
Copyright © 2019 Mobius <sv0220@163.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	goflag "flag"

	"github.com/Ankr-network/dccn-tools/ceph-toolbox/pkg/ceph"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init ceph cluster",
	Long:  ``,
	Run:   ceph.Run,
}

func init() {
	rootCmd.AddCommand(initCmd)
	goflag.Parse()
	//initCmd.PersistentFlags().StringP("foo", "f", "", "A help for foo")

}
