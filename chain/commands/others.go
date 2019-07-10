// Copyright © 2017 ZhongAn Technology
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"os"

	glb "github.com/dappledger/AnnChain/chain/commands/global"
	"github.com/dappledger/AnnChain/chain/types"
	"github.com/dappledger/AnnChain/gemmill"
	gtypes "github.com/dappledger/AnnChain/gemmill/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewShowCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "show",
		Short: "show infomation about this node",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	c.AddCommand(NewPubkeyCommand())

	return c
}

func NewPubkeyCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "pubkey",
		Short: "print this node's public key",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ang, err := gemmill.NewAngine(nil, &gemmill.Tunes{Runtime: viper.GetString("runtime")})
			if err != nil {
				cmd.Println("Create angine error: " + err.Error())
				os.Exit(1)
			}
			cmd.Println(ang.PrivValidator().PubKey)
		},
	}

	return c
}

func NewVersionCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print version of the binary",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(types.GetCommitVersion())
		},
	}

	return c
}

func NewResetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "reset",
		Short: "Reset PrivValidator, clean the data and shards",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			runtime, _ := cmd.Flags().GetString("runtime")
			if err = glb.CheckAndReadRuntimeConfig(runtime); err == nil {
				setFlags(cmd, glb.GConf())
			}
			return err
		},
		Run: resetCommandFunc,
	}

	return c
}

func resetCommandFunc(cmd *cobra.Command, args []string) {
	angineconf := glb.GConf()
	os.RemoveAll(angineconf.GetString("db_dir"))
	resetPrivValidator(angineconf.GetString("priv_validator_file"))
}

func resetPrivValidator(privValidatorFile string) {
	var (
		privValidator *gtypes.PrivValidator
		err           error
	)

	if _, err = os.Stat(privValidatorFile); err == nil {
		privValidator, err = gtypes.LoadPrivValidator(privValidatorFile)
		if err != nil {
			fmt.Println("Load PrivValidator error: ", err)
			os.Exit(1)
		}
		privValidator.Reset()
		fmt.Println("Reset PrivValidator", "file", privValidatorFile)
	} else {
		privValidator, err = gtypes.GenPrivValidator(glb.DefaultCrypto, nil)
		if err != nil {
			fmt.Println("Generate PrivValidator error: ", err)
			os.Exit(1)
		}
		privValidator.SetFile(privValidatorFile)
		if err := privValidator.Save(); err != nil {
			fmt.Println("Save PrivValidator error: ", err)
			os.Exit(1)
		}
		fmt.Printf("Generated PrivValidator file: [%v] successfully!\n", privValidatorFile)
	}
}
