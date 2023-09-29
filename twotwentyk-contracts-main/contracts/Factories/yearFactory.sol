// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../Parts/year.sol';

contract YearFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event YearCreated(address year);

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

    function createYear(uint256 collectionId, address owner) external 
        onlyAdmin returns(address year){        

        Year newYear = new Year(collectionId);
        year = address(newYear);

        newYear.transferOwnership(owner);            

        emit YearCreated(year);

        return (year);
    }   
}