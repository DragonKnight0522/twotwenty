// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ICategory {
    function mintCategory(
        string memory _tokenURI,
        address _owner
    ) external;
}

interface IYear {
   function mintYear(
        string memory _tokenURI,
        address _owner
    ) external;
}

interface IDayMonth {
    function mintDayMonth(
        string memory _tokenURI,
        address _owner
    ) external;
}

interface ICraftingCard {   
    function mintCraftingCard(
        string memory _tokenURI,
        address _owner
    ) external;
}

interface ITrigger {   
    function mintTrigger(
        string memory _tokenURI,
        address _owner
    ) external;
}