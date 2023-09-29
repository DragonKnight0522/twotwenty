// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../Parts/craftingCard.sol';

contract CraftingCardFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event CraftingCreated(address crafting);

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

    function createCrafting(uint256 collectionId, address owner) external 
        onlyAdmin returns(address crafting){        

        CraftingCard newCrafting = new CraftingCard(collectionId);
        crafting = address(newCrafting);

        newCrafting.transferOwnership(owner);            

        emit CraftingCreated(crafting);

        return (crafting);
    }   
}