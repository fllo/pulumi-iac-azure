package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		resourceGroup, err := resources.NewResourceGroup(ctx, "fllo-rg",
			&resources.ResourceGroupArgs{Tags: pulumi.StringMap{"owner": pulumi.String("fllo")}, Location: pulumi.String(`westeurope`)})
		if err != nil {
			return err
		}

		_, err = storage.NewStorageAccount(ctx, "sa", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
			Kind: pulumi.String("StorageV2"),
		})
		if err != nil {
			return err
		}

		// Export the primary key of the Storage Account
		//ctx.Export("primaryStorageKey", pulumi.All(resourceGroup.Name, account.Name).ApplyT(
		//	func(args []interface{}) (string, error) {
		//		resourceGroupName := args[0].(string)
		//		accountName := args[1].(string)
		//		accountKeys, err := storage.ListStorageAccountKeys(ctx, &storage.ListStorageAccountKeysArgs{
		//			ResourceGroupName: resourceGroupName,
		//			AccountName:       accountName,
		//		})
		//		if err != nil {
		//			return "", err
		//		}
		//
		//		return accountKeys.Keys[0].Value, nil
		//	},
		//))

		return nil
	})
}
