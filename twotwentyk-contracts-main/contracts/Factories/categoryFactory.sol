// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../Parts/category.sol';

contract CategoryFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event CategoryCreated(address category);

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin contract");
        _;
    }

    function initialize() external initializer {                        
        __Ownable_init();                
    }

    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }

    function createCategory(uint256 collectionId, address owner) external 
        onlyAdmin returns(address category){        

        Category newCategory = new Category(collectionId);
        category = address(newCategory);

        newCategory.transferOwnership(owner);            

        emit CategoryCreated(category);

        return (category);
    }   
}