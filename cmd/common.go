// Package cmd implements fyde-cli commands
package cmd

/*
Copyright © 2019 Fyde, Inc. <hello@fyde.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	apipolicies "github.com/fyde/fyde-cli/client/access_policies"
	apiproxies "github.com/fyde/fyde-cli/client/access_proxies"
	apiresources "github.com/fyde/fyde-cli/client/access_resources"
	apievents "github.com/fyde/fyde-cli/client/device_events"
	apidevices "github.com/fyde/fyde-cli/client/devices"
	apigroups "github.com/fyde/fyde-cli/client/groups"
	apiusers "github.com/fyde/fyde-cli/client/users"
)

func preRunCheckEndpoint(cmd *cobra.Command, args []string) error {
	if authViper.GetString(ckeyAuthEndpoint) == "" || global.Client == nil {
		return fmt.Errorf("endpoint not set! Run `%s endpoint [hostname]` first", ApplicationName)
	}

	return nil
}

func preRunCheckAuth(cmd *cobra.Command, args []string) error {
	err := preRunCheckEndpoint(cmd, args)
	if err != nil {
		return err
	}

	switch authViper.GetString(ckeyAuthMethod) {
	case authMethodBearerToken:
		if authViper.GetString(ckeyAuthAccessToken) == "" ||
			authViper.GetString(ckeyAuthClient) == "" ||
			authViper.GetString(ckeyAuthUID) == "" {
			return fmt.Errorf("not logged in! Run `%s login` first", ApplicationName)
		}
	case "":
	default:
		return fmt.Errorf("not logged in! Run `%s login` first", ApplicationName)
	}

	return nil
}

func processErrorResponse(err error) error {
	// TODO prepare for other error response types
	// (maybe use reflection if we can always get the Payload from within the error type)
	switch r := err.(type) {
	// user errors
	case *apiusers.GetUserNotFound:
		return fmt.Errorf("user not found")
	case *apiusers.ListUsersUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))

	// group errors
	case *apigroups.GetGroupNotFound:
		return fmt.Errorf("group not found")
	case *apigroups.ListGroupsUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))

	// device errors
	case *apidevices.ListDevicesUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))
	case *apidevices.GetDeviceNotFound:
		return fmt.Errorf("device not found")

	// resource errors
	case *apiresources.ListResourcesUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))
	case *apiresources.GetResourceNotFound:
		return fmt.Errorf("resource not found")

	// proxy errors
	case *apiproxies.ListProxiesUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))
	case *apiproxies.GetProxyNotFound:
		return fmt.Errorf("proxy not found")

	// policy errors
	case *apipolicies.ListPoliciesUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))
	case *apipolicies.GetPolicyNotFound:
		return fmt.Errorf("policy not found")

	// device event errors
	case *apievents.ListDeviceEventsUnauthorized:
		return fmt.Errorf(strings.Join(r.Payload.Errors, "\n"))
	case *apievents.GetDeviceEventNotFound:
		return fmt.Errorf("record not found")

	default:
		return err
	}
}

func preRunFlagChecks(cmd *cobra.Command, args []string) error {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}

	if _, ok := cmd.Annotations[flagInitPagination]; ok {
		err := preRunFlagCheckPagination(cmd, args)
		if err != nil {
			return err
		}
	}

	if _, ok := cmd.Annotations[flagInitSort]; ok {
		err := preRunFlagCheckSort(cmd, args)
		if err != nil {
			return err
		}
	}

	if _, ok := cmd.Annotations[flagInitFilter]; ok {
		err := preRunFlagCheckFilter(cmd, args)
		if err != nil {
			return err
		}
	}

	if _, ok := cmd.Annotations[flagInitOutput]; ok {
		err := preRunFlagCheckOutput(cmd, args)
		if err != nil {
			return err
		}
	}

	return nil
}
