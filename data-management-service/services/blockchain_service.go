package blockchain

import (
    "com.digitalidentity.datamanagementservice/config"
    "context"
    "math/big"

    // Ethereum paketleri
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    // Akıllı kontrat bağlaması (abigen ile oluşturulmuş)
    "com.digitalidentity.datamanagementservice/blockchain/contract"
)

func getClient() (*ethclient.Client, error) {
    client, err := ethclient.Dial(config.AppConfig.EthereumEndpoint)
    return client, err
}

func getContractInstance(client *ethclient.Client) (*contract.PermissionContract, error) {
    contractAddress := common.HexToAddress(config.AppConfig.ContractAddress)
    instance, err := contract.NewPermissionContract(contractAddress, client)
    return instance, err
}

func GrantPermission(dataID string, userID string, permission string) error {
    client, err := getClient()
    if err != nil {
        return err
    }

    instance, err := getContractInstance(client)
    if err != nil {
        return err
    }

    // İşlem için yetkilendirme (private key vb. gereklidir)
    auth, err := bind.NewTransactorWithChainID(/* ... */)
    if err != nil {
        return err
    }

    tx, err := instance.GrantPermission(auth, dataID, userID, permission)
    if err != nil {
        return err
    }

    // İşlemin onaylanmasını beklemek gerekebilir
    _ = tx

    return nil
}

func CheckPermission(dataID string, userID string, requiredPermission string) (bool, error) {
    client, err := getClient()
    if err != nil {
        return false, err
    }

    instance, err := getContractInstance(client)
    if err != nil {
        return false, err
    }

    opts := &bind.CallOpts{Context: context.TODO()}
    permission, err := instance.GetPermission(opts, dataID, userID)
    if err != nil {
        return false, err
    }

    // İzin kontrolü (örneğin, "owner" tüm izinlere sahiptir)
    if permission == requiredPermission || permission == "owner" {
        return true, nil
    }

    return false, nil
}

// Diğer fonksiyonlar (RevokePermission vb.)
