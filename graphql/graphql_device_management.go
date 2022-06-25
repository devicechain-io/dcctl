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

type DeviceManagementClient struct {
	graphql.Client
}

// Creates a device management GraphQL client based on command flags and other settings.
func NewDeviceManagementGraphQLClient(cmd *cobra.Command) DeviceManagementClient {
	cli := GetGraphQLClientForCommand(cmd, "device-management")
	dmclient := DeviceManagementClient{
		Client: cli,
	}
	return dmclient
}

// Assure a device type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceType(ctx context.Context, token string, name *string,
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
	resp, wascreated, err := dmgql.AssureDeviceType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device types by token.
func (dmc *DeviceManagementClient) GetDeviceTypesByToken(ctx context.Context, tokens []string) (map[string]dmgql.IDeviceType, error) {
	return dmgql.GetDeviceTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a device (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDevice(ctx context.Context, token string, deviceTypeToken string, name *string,
	description *string, metadata *string) {
	assure("device", token)
	req := dmmodel.DeviceCreateRequest{
		Token:           token,
		DeviceTypeToken: deviceTypeToken,
		Name:            name,
		Description:     description,
		Metadata:        metadata,
	}
	resp, wascreated, err := dmgql.AssureDevice(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get devices by token.
func (dmc *DeviceManagementClient) GetDevicesByToken(ctx context.Context, tokens []string) (map[string]dmgql.IDevice, error) {
	return dmgql.GetDevicesByToken(ctx, dmc.Client, tokens)
}

// Assure a device relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceRelationshipType(ctx context.Context, token string, name *string,
	description *string, metadata *string) {
	assure("device relationship type", token)
	req := dmmodel.DeviceRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureDeviceRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device relationship types by token.
func (dmc *DeviceManagementClient) GetDeviceRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IDeviceRelationshipType, error) {
	return dmgql.GetDeviceRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a device relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceRelationship(ctx context.Context, token string, source string,
	target string, relation string, metadata *string) {
	assure("device relationship", token)
	req := dmmodel.DeviceRelationshipCreateRequest{
		Token:            token,
		SourceDevice:     source,
		TargetDevice:     target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureDeviceRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device relationships by token.
func (dmc *DeviceManagementClient) GetDeviceRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IDeviceRelationship, error) {
	return dmgql.GetDeviceRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure a device group (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceGroup(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("device group", token)
	req := dmmodel.DeviceGroupCreateRequest{
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
	resp, wascreated, err := dmgql.AssureDeviceGroup(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device groups by token.
func (dmc *DeviceManagementClient) GetDeviceGroupsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IDeviceGroup, error) {
	return dmgql.GetDeviceGroupsByToken(ctx, dmc.Client, tokens)
}

// Assure a device group relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceGroupRelationshipType(ctx context.Context,
	token string, name *string, description *string, metadata *string) {
	assure("device group relationship type", token)
	req := dmmodel.DeviceGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureDeviceGroupRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device group relationship types by token.
func (dmc *DeviceManagementClient) GetDeviceGroupRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IDeviceGroupRelationshipType, error) {
	return dmgql.GetDeviceGroupRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a device group relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureDeviceGroupRelationship(ctx context.Context,
	token string, deviceGroup string, device string, relation string, metadata *string) {
	assure("device group relationship", token)
	req := dmmodel.DeviceGroupRelationshipCreateRequest{
		Token:            token,
		DeviceGroup:      deviceGroup,
		Device:           device,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureDeviceGroupRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device group relationships by token.
func (dmc *DeviceManagementClient) GetDeviceGroupRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IDeviceGroupRelationship, error) {
	return dmgql.GetDeviceGroupRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure an asset type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetType(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("asset type", token)
	req := dmmodel.AssetTypeCreateRequest{
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
	resp, wascreated, err := dmgql.AssureAssetType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset types by token.
func (dmc *DeviceManagementClient) GetAssetTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAssetType, error) {
	return dmgql.GetAssetTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an asset (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAsset(ctx context.Context, token string, assetTypeToken string, name *string,
	description *string, metadata *string) {
	assure("asset", token)
	req := dmmodel.AssetCreateRequest{
		Token:          token,
		AssetTypeToken: assetTypeToken,
		Name:           name,
		Description:    description,
		Metadata:       metadata,
	}
	resp, wascreated, err := dmgql.AssureAsset(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get assets by token.
func (dmc *DeviceManagementClient) GetAssetsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAsset, error) {
	return dmgql.GetAssetsByToken(ctx, dmc.Client, tokens)
}

// Assure an asset relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetRelationshipType(ctx context.Context, token string, name *string,
	description *string, metadata *string) {
	assure("asset relationship type", token)
	req := dmmodel.AssetRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureAssetRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset relationship types by token.
func (dmc *DeviceManagementClient) GetAssetRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAssetRelationshipType, error) {
	return dmgql.GetAssetRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an asset relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetRelationship(ctx context.Context, token string, source string,
	target string, relation string, metadata *string) {
	assure("asset relationship", token)
	req := dmmodel.AssetRelationshipCreateRequest{
		Token:            token,
		SourceAsset:      source,
		TargetAsset:      target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureAssetRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset relationships by token.
func (dmc *DeviceManagementClient) GetAssetRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAssetRelationship, error) {
	return dmgql.GetAssetRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure an asset group (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetGroup(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("asset group", token)
	req := dmmodel.AssetGroupCreateRequest{
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
	resp, wascreated, err := dmgql.AssureAssetGroup(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset groups by token.
func (dmc *DeviceManagementClient) GetAssetGroupsByToken(ctx context.Context, tokens []string) (map[string]dmgql.IAssetGroup, error) {
	return dmgql.GetAssetGroupsByToken(ctx, dmc.Client, tokens)
}

// Assure an asset group relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetGroupRelationshipType(ctx context.Context, token string, name *string,
	description *string, metadata *string) {
	assure("asset group relationship type", token)
	req := dmmodel.AssetGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureAssetGroupRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset group relationship types by token.
func (dmc *DeviceManagementClient) GetAssetGroupRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAssetGroupRelationshipType, error) {
	return dmgql.GetAssetGroupRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an asset group relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAssetGroupRelationship(ctx context.Context, token string, assetGroup string,
	asset string, relation string, metadata *string) {
	assure("asset group relationship", token)
	req := dmmodel.AssetGroupRelationshipCreateRequest{
		Token:            token,
		AssetGroup:       assetGroup,
		Asset:            asset,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureAssetGroupRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset group relationships by token.
func (dmc *DeviceManagementClient) GetAssetGroupRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAssetGroupRelationship, error) {
	return dmgql.GetAssetGroupRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure an area type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaType(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("area type", token)
	req := dmmodel.AreaTypeCreateRequest{
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
	resp, wascreated, err := dmgql.AssureAreaType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area types by token.
func (dmc *DeviceManagementClient) GetAreaTypesByToken(ctx context.Context, tokens []string) (map[string]dmgql.IAreaType, error) {
	return dmgql.GetAreaTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an area (check for existing or create new).
func (dmc *DeviceManagementClient) AssureArea(ctx context.Context, token string, areaTypeToken string, name *string,
	description *string, metadata *string) {
	assure("area", token)
	req := dmmodel.AreaCreateRequest{
		Token:         token,
		AreaTypeToken: areaTypeToken,
		Name:          name,
		Description:   description,
		Metadata:      metadata,
	}
	resp, wascreated, err := dmgql.AssureArea(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get areas by token.
func (dmc *DeviceManagementClient) GetAreasByToken(ctx context.Context, tokens []string) (map[string]dmgql.IArea, error) {
	return dmgql.GetAreasByToken(ctx, dmc.Client, tokens)
}

// Assure an area relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaRelationshipType(ctx context.Context, token string, name *string,
	description *string, metadata *string) {
	assure("area relationship type", token)
	req := dmmodel.AreaRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureAreaRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area relationship types by token.
func (dmc *DeviceManagementClient) GetAreaRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAreaRelationshipType, error) {
	return dmgql.GetAreaRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an area relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaRelationship(ctx context.Context, token string, source string,
	target string, relation string, metadata *string) {
	assure("area relationship", token)
	req := dmmodel.AreaRelationshipCreateRequest{
		Token:            token,
		SourceArea:       source,
		TargetArea:       target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureAreaRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area relationships by token.
func (dmc *DeviceManagementClient) GetAreaRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAreaRelationship, error) {
	return dmgql.GetAreaRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure an area group (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaGroup(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("area group", token)
	req := dmmodel.AreaGroupCreateRequest{
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
	resp, wascreated, err := dmgql.AssureAreaGroup(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area groups by token.
func (dmc *DeviceManagementClient) GetAreaGroupsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAreaGroup, error) {
	return dmgql.GetAreaGroupsByToken(ctx, dmc.Client, tokens)
}

// Assure an area group relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaGroupRelationshipType(ctx context.Context, token string, name *string,
	description *string, metadata *string) {
	assure("area group relationship type", token)
	req := dmmodel.AreaGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureAreaGroupRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area group relationship types by token.
func (dmc *DeviceManagementClient) GetAreaGroupRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAreaGroupRelationshipType, error) {
	return dmgql.GetAreaGroupRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure an area group relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureAreaGroupRelationship(ctx context.Context, token string, areaGroup string,
	area string, relation string, metadata *string) {
	assure("area group relationship", token)
	req := dmmodel.AreaGroupRelationshipCreateRequest{
		Token:            token,
		AreaGroup:        areaGroup,
		Area:             area,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureAreaGroupRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area group relationships by token.
func (dmc *DeviceManagementClient) GetAreaGroupRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.IAreaGroupRelationship, error) {
	return dmgql.GetAreaGroupRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure a customer type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerType(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("customer type", token)
	req := dmmodel.CustomerTypeCreateRequest{
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
	resp, wascreated, err := dmgql.AssureCustomerType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer types by token.
func (dmc *DeviceManagementClient) GetCustomerTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerType, error) {
	return dmgql.GetCustomerTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a customer (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomer(ctx context.Context, token string, customerTypeToken string,
	name *string, description *string, metadata *string) {
	assure("customer", token)
	req := dmmodel.CustomerCreateRequest{
		Token:             token,
		CustomerTypeToken: customerTypeToken,
		Name:              name,
		Description:       description,
		Metadata:          metadata,
	}
	resp, wascreated, err := dmgql.AssureCustomer(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customers by token.
func (dmc *DeviceManagementClient) GetCustomersByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomer, error) {
	return dmgql.GetCustomersByToken(ctx, dmc.Client, tokens)
}

// Assure a customer relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerRelationshipType(ctx context.Context, token string,
	name *string, description *string, metadata *string) {
	assure("customer relationship type", token)
	req := dmmodel.CustomerRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureCustomerRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer relationship types by token.
func (dmc *DeviceManagementClient) GetCustomerRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerRelationshipType, error) {
	return dmgql.GetCustomerRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a customer relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerRelationship(ctx context.Context, token string, source string,
	target string, relation string, metadata *string) {
	assure("customer relationship", token)
	req := dmmodel.CustomerRelationshipCreateRequest{
		Token:            token,
		SourceCustomer:   source,
		TargetCustomer:   target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureCustomerRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer relationships by token.
func (dmc *DeviceManagementClient) GetCustomerRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerRelationship, error) {
	return dmgql.GetCustomerRelationshipsByToken(ctx, dmc.Client, tokens)
}

// Assure a customer group (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerGroup(ctx context.Context, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("customer group", token)
	req := dmmodel.CustomerGroupCreateRequest{
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
	resp, wascreated, err := dmgql.AssureCustomerGroup(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer groups by token.
func (dmc *DeviceManagementClient) GetCustomerGroupsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerGroup, error) {
	return dmgql.GetCustomerGroupsByToken(ctx, dmc.Client, tokens)
}

// Assure a customer group relationship type (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerGroupRelationshipType(ctx context.Context, token string,
	name *string, description *string, metadata *string) {
	assure("customer group relationship type", token)
	req := dmmodel.CustomerGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	resp, wascreated, err := dmgql.AssureCustomerGroupRelationshipType(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer group relationship types by token.
func (dmc *DeviceManagementClient) GetCustomerGroupRelationshipTypesByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerGroupRelationshipType, error) {
	return dmgql.GetCustomerGroupRelationshipTypesByToken(ctx, dmc.Client, tokens)
}

// Assure a customer group relationship (check for existing or create new).
func (dmc *DeviceManagementClient) AssureCustomerGroupRelationship(ctx context.Context, token string,
	customerGroup string, customer string, relation string, metadata *string) {
	assure("customer group relationship", token)
	req := dmmodel.CustomerGroupRelationshipCreateRequest{
		Token:            token,
		CustomerGroup:    customerGroup,
		Customer:         customer,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	resp, wascreated, err := dmgql.AssureCustomerGroupRelationship(ctx, dmc.Client, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer group relationships by token.
func (dmc *DeviceManagementClient) GetCustomerGroupRelationshipsByToken(ctx context.Context,
	tokens []string) (map[string]dmgql.ICustomerGroupRelationship, error) {
	return dmgql.GetCustomerGroupRelationshipsByToken(ctx, dmc.Client, tokens)
}
