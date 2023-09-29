// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPackFactory{
    function createCardPack(
            uint256 _standardAmount,
            uint256 _premiumAmount,
            uint256 _eliteAmount,
            uint256 _totalAmount,
            address owner     
    ) external returns(address);
}

interface ICategoryFactory {   
    function createCategory(uint256 collectionId, address owner) external returns(address);
}

interface IYearFactory {   
    function createYear(uint256 collectionId, address owner) external returns(address);
}

interface IDayMonthFactory {   
    function createDayMonth(uint256 collectionId, address owner) external returns(address);
}

interface ICraftingFactory {   
    function createCrafting(uint256 collectionId, address owner) external returns(address);
}

interface ITriggerFactory {   
    function createTrigger(uint256 collectionId, address owner) external returns(address);
}