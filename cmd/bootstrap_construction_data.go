/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Khan/genqlient/graphql"
	gql "github.com/devicechain-io/dcctl/graphql"
	"github.com/spf13/cobra"
)

const (
	DATASET_CONSTRUCTION = "Construction"
)

// Runes available for generating vehicle identifiers.
var vinRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Create common command for creating DeviceChain resources
var constructionDataCmd = &cobra.Command{
	Use:   "construction",
	Short: "Bootstrap construction sample data",
	Long:  `Bootstraps system microservices with sample data for construction use case`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return bootstrapConstructionData(context.Background(), cmd)
	},
	SilenceUsage: true,
}

// Bootstraps system microservices with construction sample dataset.
func bootstrapConstructionData(ctx context.Context, cmd *cobra.Command) error {
	title(DATASET_CONSTRUCTION)
	gqlcli := gql.GetDeviceManagementGraphQLClient(cmd)

	// Bootstrap asset data.
	bootstrapAssetData(ctx, gqlcli)

	// Bootstrap device data.
	bootstrapDeviceData(ctx, gqlcli)

	footer(DATASET_CONSTRUCTION)
	return nil
}

// Bootstraps asset data for construction dataset.
func bootstrapAssetData(ctx context.Context, gqlcli graphql.Client) {
	header("Asset Types", DATASET_CONSTRUCTION)
	bootstrapAssetTypes(ctx, gqlcli)

	header("Asset Groups", DATASET_CONSTRUCTION)
	bootstrapAssetGroups(ctx, gqlcli)

	header("Asset Group Relationship Types", DATASET_CONSTRUCTION)
	bootstrapAssetGroupRelationshipTypes(ctx, gqlcli)

	header("Assets", DATASET_CONSTRUCTION)
	bootstrapAssets(ctx, gqlcli)
}

// Bootstrap asset types.
func bootstrapAssetTypes(ctx context.Context, client graphql.Client) {
	// Cat D6 (bulldozer)
	gql.AssureAssetType(ctx, client, "catd6", s("Cat D6"),
		unspace(`Move material at a lower cost with a fully automatic transmission, outstanding fuel efficiency 
		and reduced service/maintenance costs. The broadest range of technology features in the industry work 
		together seamlessly to help you make the most of your equipment investment.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat C9.3B",
			"netPower": "215 HP",
			"operatingWeight": "50733 lb"
		}`))
	// Cat 725 (truck)
	gql.AssureAssetType(ctx, client, "cat725", s("Cat 725 Articulated Truck"),
		unspace(`The Cat® 725 features a world-class cab design, re-engineered using global operator feedback to 
		advance comfort and ease of operation. Enhancements include new controls, transmission-protection features, 
		hoist-assist system, advanced automatic traction control system, automatic retarder control, stability-assist 
		machine rollover warning system, and a fuel saving ECO mode.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat C9.3",
			"ratedPayload": "26.5 ton",
			"heaped": "9.6 yd³"
		}`))
	// Cat 730 (truck)
	gql.AssureAssetType(ctx, client, "cat730", s("Cat 730 Articulated Truck"),
		unspace(`The Cat® 730 features a world-class cab design, re-engineered using global operator feedback to advance 
		comfort and ease of operation. Enhancements include new controls, transmission-protection features, hoist-assist 
		system, advanced automatic traction control system, stability-assist machine rollover warning system, and a fuel 
		saving ECO mode.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat® C13",
			"ratedPayload": "31 ton",
			"heaped": "23 yd³"
		}`))
	// Cat 313 (excavator)
	gql.AssureAssetType(ctx, client, "cat313", s("Cat 313 Small Excavator"),
		unspace(`The 313 excavator offers superior performance and operator efficiency. Standard, easy-to-use Cat technologies, 
		a new cab focused on operator comfort and productivity, and low fuel and maintenance costs allow you to move material 
		all day with speed and precision while keeping more money in your pocket.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "108 HP",
			"operatingWeight": "30400 lb",
			"maxDigDepth": "19.8 ft"
		}`))
	// Cat 317 (excavator)
	gql.AssureAssetType(ctx, client, "cat317", s("Cat 317 Small Excavator"),
		unspace(`The 317 Hydraulic Excavator boosts productivity on your jobsite. Standard, easy-to-use Cat® technologies, 
		performance upgrades, and low fuel and maintenance costs allow you to move material all day with speed and precision 
		while keeping more money in your pocket.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "130 HP",
			"operatingWeight": "40200 lb",
			"maxDigDepth": "21 ft"
		}`))
	// Cat D1 (mini dozer)
	gql.AssureAssetType(ctx, client, "catd1", s("Cat D1"),
		unspace(`The new Cat® D1 delivers superior performance and the broadest choice of technology features to 
		help you get the most from your dozer. Nimble and responsive, it has power for dozing and ﬁnesse for 
		grading. Fully hydrostatic transmission gives you seamless acceleration, so you can get the job done quickly. 
		The load sensing system automatically optimizes ground speed based on load, for the greatest productivity and 
		fuel efficiency.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat C3.6",
			"netPower": "80 HP",
			"operatingWeight": "17855 lb"
		}`))
	// Cat 302 CR (mini excavator)
	gql.AssureAssetType(ctx, client, "cat302cr", s("Cat 302 CR Mini Excavator"),
		unspace(`The Cat® 302 CR Mini Excavator delivers power and performance in a compact size to help you work in a 
		wide range of applications.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "21 HP",
			"operatingWeight": "3913 lb",
			"maxDigDepth": "100 in"
		}`))
	// Cat 305 CR (mini excavator)
	gql.AssureAssetType(ctx, client, "cat305cr", s("Cat 305 CR Mini Excavator"),
		unspace(`The Cat® 305 CR Mini Excavator delivers power and performance in a compact size to help you work in a 
		wide range of applications.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "45 HP",
			"operatingWeight": "12688 lb",
			"maxDigDepth": "144.5 in"
		}`))
	// Cat 415 IL (wheel loader)
	gql.AssureAssetType(ctx, client, "cat415il", s("Cat 415 IL Backhoe Loader"),
		unspace(`The Cat® 415 IL Industrial Loader delivers great performance, improved fuel efficiency, and a superior 
		hydraulic system.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "69 HP",
			"operatingWeight": "17637 lb",
			"engineModel": "Cat C3.6"
		}`))
	// Cat 416 (wheel loader)
	gql.AssureAssetType(ctx, client, "cat416", s("Cat 416 Backhoe Loader"),
		unspace(`The Cat® 416 Backhoe Loader delivers exceptional performance, increased fuel efficiency, superior hydraulic 
		system and an updated operator station.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"netPower": "86 HP",
			"operatingWeight": "14 ft",
			"engineModel": "24251 lb"
		}`))
}

// Bootstrap asset groups.
func bootstrapAssetGroups(ctx context.Context, client graphql.Client) {
	// Bulldozers
	gql.AssureAssetGroup(ctx, client, "bulldoz", s("Bulldozers"),
		unspace(`Group which includes bulldozers of various types`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil, nil)
	// Trucks
	gql.AssureAssetGroup(ctx, client, "truck", s("Trucks"),
		unspace(`Group which includes trucks of various types`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil, nil)
	// Excavators
	gql.AssureAssetGroup(ctx, client, "excavator", s("Excavators"),
		unspace(`Group which includes excavators of various types`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil, nil)
	// Wheel Loaders
	gql.AssureAssetGroup(ctx, client, "wloaders", s("Wheel Loaders"),
		unspace(`Group which includes wheel loaders of various types`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil, nil)
}

// Bootstrap asset group relationship types.
func bootstrapAssetGroupRelationshipTypes(ctx context.Context, client graphql.Client) {
	// Tracks location of
	gql.AssureAssetGroupRelationshipType(ctx, client, "contains", s("Contains"),
		unspace(`The group contains the target asset`), nil)
}

// Bootstrap multiple assets of a given type by generating VINs for each.
func bootstrapAssetsOfType(ctx context.Context, client graphql.Client, typeToken string, groupToken string, count int) {
	atypes, err := gql.GetAssetTypesByToken(ctx, client, []string{typeToken})
	if err != nil {
		panic(err)
	}
	atype := atypes[typeToken]
	agroups, err := gql.GetAssetGroupsByToken(ctx, client, []string{groupToken})
	if err != nil {
		panic(err)
	}
	agroup := agroups[groupToken]
	for idx := 0; idx < count; idx++ {
		vin := randomVin(15)
		gql.AssureAsset(ctx, client, vin, atype.GetToken(), s(fmt.Sprintf("%s VIN:%s", atype.GetName(), vin)),
			unspace(fmt.Sprintf("%s VIN:%s", atype.GetDescription(), vin)),
			s(fmt.Sprintf(*unspace(`
			{
				"vin": "%s",
				"purchaseDate": "%s"
			}`), vin, time.Now().Format("2006-01-02"))))
		gql.AssureAssetGroupRelationship(ctx, client, fmt.Sprintf("%s-contains-%s", agroup.GetToken(), vin),
			agroup.GetToken(), vin, "contains", nil)
	}
}

// Bootstrap assets.
func bootstrapAssets(ctx context.Context, client graphql.Client) {
	bootstrapAssetsOfType(ctx, client, "catd6", "bulldoz", 3)
	bootstrapAssetsOfType(ctx, client, "cat725", "truck", 3)
	bootstrapAssetsOfType(ctx, client, "cat730", "truck", 3)
	bootstrapAssetsOfType(ctx, client, "cat313", "excavator", 3)
	bootstrapAssetsOfType(ctx, client, "cat317", "excavator", 3)
	bootstrapAssetsOfType(ctx, client, "catd1", "bulldoz", 3)
	bootstrapAssetsOfType(ctx, client, "cat302cr", "excavator", 3)
	bootstrapAssetsOfType(ctx, client, "cat305cr", "excavator", 3)
	bootstrapAssetsOfType(ctx, client, "cat415il", "wloaders", 3)
	bootstrapAssetsOfType(ctx, client, "cat416", "wloaders", 3)
}

func init() {
	bootstrapCmd.AddCommand(constructionDataCmd)
}

// Bootstraps device data for construction dataset.
func bootstrapDeviceData(ctx context.Context, gqlcli graphql.Client) {
	header("Device Types", DATASET_CONSTRUCTION)
	bootstrapDeviceTypes(ctx, gqlcli)

	header("Devices", DATASET_CONSTRUCTION)
	bootstrapDevices(ctx, gqlcli)

	header("Device Relationship Types", DATASET_CONSTRUCTION)
	bootstrapDeviceRelationshipTypes(ctx, gqlcli)

	header("Device Relationships", DATASET_CONSTRUCTION)
	bootstrapDeviceRelationships(ctx, gqlcli)

	header("Device Groups", DATASET_CONSTRUCTION)
	bootstrapDeviceGroups(ctx, gqlcli)

	header("Device Group Relationship Types", DATASET_CONSTRUCTION)
	bootstrapDeviceGroupRelationshipTypes(ctx, gqlcli)

	header("Device Group Relationships", DATASET_CONSTRUCTION)
	bootstrapDeviceGroupRelationships(ctx, gqlcli)
}

// Bootstrap device types.
func bootstrapDeviceTypes(ctx context.Context, client graphql.Client) {
	// Cat D1
	gql.AssureDeviceType(ctx, client, "catd1", s("Cat D1"),
		unspace(`The new Cat® D1 delivers superior performance and the broadest choice of technology features to 
		help you get the most from your dozer. Nimble and responsive, it has power for dozing and ﬁnesse for 
		grading. Fully hydrostatic transmission gives you seamless acceleration, so you can get the job done quickly. 
		The load sensing system automatically optimizes ground speed based on load, for the greatest productivity and 
		fuel efficiency.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat C3.6",
			"powerNet": "80 HP",
			"operatingWeight": "17855 lb"
		}`))
}

// Bootstrap devices.
func bootstrapDevices(ctx context.Context, client graphql.Client) {
	// Cat D1 SDK7GV3WXZ3FBXZ
	gql.AssureDevice(ctx, client, "SDK7GV3WXZ3FBXZ", "catd1", s("Cat D1 VIN:SDK7GV3WXZ3FBXZ"),
		unspace(`This is a Cat D1 with VIN SDK7GV3WXZ3FBXZ`),
		unspace(`
		{
			"vin": "SDK7GV3WXZ3FBXZ",
			"owner": "CatCorp",
			"purchaseDate": "2022/01/01"
		}`))
	// Cat D1 WDVM4L7YPRM7HU2
	gql.AssureDevice(ctx, client, "WDVM4L7YPRM7HU2", "catd1", s("Cat D1 VIN:WDVM4L7YPRM7HU2"),
		unspace(`This is a Cat D1 with VIN WDVM4L7YPRM7HU2`),
		unspace(`
		{
			"vin": "WDVM4L7YPRM7HU2",
			"owner": "CatCorp",
			"purchaseDate": "2022/02/01"
		}`))
}

// Bootstrap device relationship types.
func bootstrapDeviceRelationshipTypes(ctx context.Context, client graphql.Client) {
	// Tracks location of
	gql.AssureDeviceRelationshipType(ctx, client, "tracksLocationOf", s("Tracks location of"),
		unspace(`The source device tracks the location of the target device`),
		unspace(`
		{
			"accuracy": "1 meter"
		}`))
	// Tracks temperature of
	gql.AssureDeviceRelationshipType(ctx, client, "tracksTempOf", s("Tracks temperature of"),
		unspace(`The source device tracks the temperature of the target device`),
		unspace(`
		{
			"accuracy": "1 degree C"
		}`))
}

// Bootstrap device relationships.
func bootstrapDeviceRelationships(ctx context.Context, client graphql.Client) {
	// SDK7GV3WXZ3FBXZ tracksLocationOf WDVM4L7YPRM7HU2
	gql.AssureDeviceRelationship(ctx, client, "SDK7GV3WXZ3FBXZ-tracksLocationOf-WDVM4L7YPRM7HU2",
		"SDK7GV3WXZ3FBXZ", "WDVM4L7YPRM7HU2", "tracksLocationOf",
		unspace(`
		{
			"accuracy": "1 meter"
		}`))
}

// Bootstrap device groups.
func bootstrapDeviceGroups(ctx context.Context, client graphql.Client) {
	// Small Dozers
	gql.AssureDeviceGroup(ctx, client, "smalldoz", s("Small Dozers"),
		unspace(`Under 105 hp, the Cat® small dozers are designed to optimize speed, transportability, maneuverability, 
		versatility and finish grading accuracy. These crawler dozers are ideal for residential construction performing 
		such tasks as clearing and grading lots, sloping the sides of roads, back-filling, and final grade work for 
		landscaping and driveway construction.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"maxWeight": "20000 lb"
		}`))
}

// Bootstrap device group relationship types.
func bootstrapDeviceGroupRelationshipTypes(ctx context.Context, client graphql.Client) {
	// Tracks location of
	gql.AssureDeviceGroupRelationshipType(ctx, client, "contains", s("Contains"),
		unspace(`The group contains the target device`), nil)
}

// Bootstrap device group relationships.
func bootstrapDeviceGroupRelationships(ctx context.Context, client graphql.Client) {
	// smalldoz contains SDK7GV3WXZ3FBXZ
	gql.AssureDeviceGroupRelationship(ctx, client, "smalldoz-contains-SDK7GV3WXZ3FBXZ",
		"smalldoz", "SDK7GV3WXZ3FBXZ", "contains", nil)
	// smalldoz contains WDVM4L7YPRM7HU2
	gql.AssureDeviceGroupRelationship(ctx, client, "smalldoz-contains-WDVM4L7YPRM7HU2",
		"smalldoz", "WDVM4L7YPRM7HU2", "contains", nil)
}

// Create a random VIN of the given length.
func randomVin(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = vinRunes[rand.Intn(len(vinRunes))]
	}
	return string(b)
}

// Seed random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}
