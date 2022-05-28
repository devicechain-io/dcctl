/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package graphql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	dmgql "github.com/devicechain-io/dc-device-management/gqlclient"
	dmmodel "github.com/devicechain-io/dc-device-management/model"
	"github.com/spf13/cobra"
)

// Gets a GraphQL client based on command flags and other settings.
func GetDeviceManagementGraphQLClient(cmd *cobra.Command) graphql.Client {
	return GetGraphQLClientForCommand(cmd, "device-management")
}

// Assure a device type (check for existing or create new).
func AssureDeviceType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("device type", token)
	req := dmmodel.DeviceTypeCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceTypeByToken.Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceType.Token)
	}
}

// Assure a device (check for existing or create new).
func AssureDevice(ctx context.Context, cli graphql.Client, token string, deviceTypeToken string, name *string,
	description *string, metadata *string) {
	assure("device", token)
	req := dmmodel.DeviceCreateRequest{
		Token:           token,
		DeviceTypeToken: deviceTypeToken,
		Name:            name,
		Description:     description,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureDevice(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceByToken.Token)
	}
	if cresp != nil {
		created(cresp.CreateDevice.Token)
	}
}

// Assure a device relationship type (check for existing or create new).
func AssureDeviceRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("device relationship type", token)
	req := dmmodel.DeviceRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceRelationshipTypeByToken.Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceRelationshipType.Token)
	}
}

// Assure a device relationship (check for existing or create new).
func AssureDeviceRelationship(ctx context.Context, cli graphql.Client, token string, source string,
	target string, relation string, metadata *string) {
	assure("device relationship type", token)
	req := dmmodel.DeviceRelationshipCreateRequest{
		Token:            token,
		SourceDevice:     source,
		TargetDevice:     target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceRelationshipByToken.Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceRelationship.Token)
	}
}
