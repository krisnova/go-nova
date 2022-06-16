/*===========================================================================*\
 *           MIT License Copyright (c) 2022 Kris Nóva <kris@nivenly.com>     *
 *                                                                           *
 *                ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓                *
 *                ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗   ┃                *
 *                ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗  ┃                *
 *                ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║  ┃                *
 *                ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║  ┃                *
 *                ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║  ┃                *
 *                ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝  ┃                *
 *                ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛                *
 *                                                                           *
 *                       This machine kills fascists.                        *
 *                                                                           *
\*===========================================================================*/

package main

import (
	"fmt"
	"os"
)

var env = &EnvironmentOptions{}

type EnvironmentOptions struct {

	// example fields
	token string
}

var (
	envOpt   = &EnvironmentOptions{}
	registry = []*EnvironmentVariable{
		{
			Name:        "NOVA_TOKEN",
			Value:       "",
			Destination: &envOpt.token,
			Required:    true,
		},
	}
)

type EnvironmentVariable struct {
	Name        string
	Value       string
	Destination *string
	Required    bool
}

func Environment() error {
	for _, v := range registry {
		v.Value = os.Getenv(v.Name)
		if v.Required && v.Value == "" {
			// If required and the variable is empty
			return fmt.Errorf("empty or undefined environmental variable [%s]", v.Name)
		}
		*v.Destination = v.Value
	}
	return nil
}
