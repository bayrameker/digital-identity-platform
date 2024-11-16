// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PermissionContract {
    struct Permission {
        address user;
        string dataId;
        string permission; // "read", "write", "delete", "owner"
    }

    mapping(string => Permission[]) private dataPermissions; // dataId => permissions

    function grantPermission(string memory dataId, address user, string memory permission) public {
        // Erişim kontrolü yapılabilir (örn. sadece veri sahibi izin verebilir)
        dataPermissions[dataId].push(Permission(user, dataId, permission));
    }

    function getPermission(string memory dataId, address user) public view returns (string memory) {
        Permission[] memory permissions = dataPermissions[dataId];
        for (uint i = 0; i < permissions.length; i++) {
            if (permissions[i].user == user) {
                return permissions[i].permission;
            }
        }
        return "";
    }

    function revokePermission(string memory dataId, address user) public {
        // Erişim kontrolü yapılabilir
        Permission[] storage permissions = dataPermissions[dataId];
        for (uint i = 0; i < permissions.length; i++) {
            if (permissions[i].user == user) {
                permissions[i] = permissions[permissions.length - 1];
                permissions.pop();
                break;
            }
        }
    }
}
