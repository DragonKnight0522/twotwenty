// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../Parts/dayMonth.sol';

contract DayMonthFactory is Initializable, OwnableUpgradeable {    
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

    function createDayMonth(uint256 collectionId, address owner) external 
        onlyAdmin returns(address daymonth){        

        DayMonth newDayMonth = new DayMonth(collectionId);
        daymonth = address(newDayMonth);

        newDayMonth.transferOwnership(owner);            

        emit DayMonthCreated(daymonth);

        return (daymonth);
    }   
}