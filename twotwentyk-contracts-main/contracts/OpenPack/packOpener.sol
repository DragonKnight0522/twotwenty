// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Interfaces/cardInterface.sol";
import "../Interfaces/packInterface.sol";

/// @title A contract to manage the opening of card packs.
/// @dev The contract is Ownable, and the owner has exclusive rights to perform certain actions.
contract PackOpener is Initializable, OwnableUpgradeable {
    ICardPack cardPackContract;
    IDayMonth dayMonthContract;
    IYear yearContract;
    ICategory categoryContract;
    ICraftingCard craftingCardContract;
    ITrigger triggerContract;
    address adminWallet;

    /// @notice Initializes the contract with specified parameters.
    /// @dev The contract takes the number of each type of card, and the addresses of the card contracts.

    function initialize(        
    ) public initializer {        
        adminWallet = msg.sender;        
        
        __Ownable_init();                
    }

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin wallet");
        _;
    }

    /// @notice Ensures that only unopened packs can be operated on.
    /// @param _tokenId The ID of the pack that the owner wants to operate on.
    modifier onlyNonOpenedPacks(uint _tokenId) {
        require(
            false == cardPackContract.isOpened(_tokenId),
            "Pack already opened"
        );
        _;
    }

    /// @notice Allows the owner of a pack to open it.
    /// @dev When a pack is opened, new cards of each type are created.
    /// @param _tokenId The ID of the pack that the owner wants to open.
    /// @param _cardPackAddress The address of the CardPack contract.    
    /// @param _dayMonthAddress The address of the DayMonth contract.
    /// @param _yearAddress The address of the Year contract.
    /// @param _categoryAddress The address of the Category contract.
    /// @param _craftingCardAddress The address of the CraftingCard contract.        
    /// @param _tokenDayMonths The URI for the DayMonth cards.
    /// @param _tokenYears The URI for the Year cards.
    /// @param _tokenCategories The URI for the Category cards.
    /// @param _tokenCraftingCards The URI for the CraftingCard cards.
    function openPack(        
        uint256 _tokenId,        
        address _cardPackAddress,
        address _categoryAddress,        
        address _yearAddress,
        address _dayMonthAddress,
        address _craftingCardAddress,
        address _triggerContract,
        
        string[] memory _tokenCategories,
        string[] memory _tokenYears,
        string[] memory _tokenDayMonths,
        string[] memory _tokenCraftingCards,
        string[] memory _tokenTriggers
    ) public onlyAdmin {        
        
        cardPackContract = ICardPack(_cardPackAddress);
        require(
            false == cardPackContract.isOpened(_tokenId),
            "Pack already opened"
        );

        categoryContract = ICategory(_categoryAddress);
        dayMonthContract = IDayMonth(_dayMonthAddress);
        yearContract = IYear(_yearAddress);        
        craftingCardContract = ICraftingCard(_craftingCardAddress);
        triggerContract = ITrigger(_triggerContract);

        address tokenOwner = cardPackContract.ownerOf(_tokenId);
        cardPackContract.changeToOpened(_tokenId);

        for(uint256 i = 0; i < _tokenCategories.length; i++) {
            categoryContract.mintCategory(_tokenCategories[i], tokenOwner);
        }

        for(uint256 i = 0; i < _tokenYears.length; i++) {
            yearContract.mintYear(_tokenYears[i], tokenOwner);
        }

        for(uint256 i = 0; i < _tokenDayMonths.length; i++) {
            dayMonthContract.mintDayMonth(_tokenDayMonths[i], tokenOwner);
        }

        for(uint256 i = 0; i < _tokenCraftingCards.length; i++) {
            craftingCardContract.mintCraftingCard(_tokenCraftingCards[i], tokenOwner);
        }

        for(uint256 i = 0; i < _tokenTriggers.length; i++) {
            triggerContract.mintTrigger(_tokenTriggers[i], tokenOwner);
        }
        
    }
    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }
}
