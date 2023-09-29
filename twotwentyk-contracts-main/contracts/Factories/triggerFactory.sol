// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../Parts/triggers.sol';

contract TriggerFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event DayMonthCreated(address daymonth);

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

    function createTrigger(uint256 collectionId, address owner) external 
        onlyAdmin returns(address trigger){        

        Trigger newTrigger = new Trigger(collectionId);
        trigger = address(newTrigger);

        newTrigger.transferOwnership(owner);            

        emit DayMonthCreated(trigger);

        return (trigger);
    }   
}