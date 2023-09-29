// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";


/// @title A contract for interacting with the CraftingCard contract.
interface ICraftingCard {
    /// @notice Creates a new crafting card.
    /// @param tokenURI The URI of the crafting card.
    /// @return The new crafting card's token ID.
    function createCraftingCard(
        string memory tokenURI,
        address owner
    ) external returns (uint256);
}

/// @title A contract for interacting with the Category contract.
interface ICategory {
    /// @notice Creates a new category.
    /// @param tokenURI The URI of the category.
    /// @return The new category's token ID.
    function createCategory(
        string memory tokenURI,
        address owner
    ) external returns (uint256);
}

/// @title A contract for interacting with the Trigger contract.
interface ITrigger {
    /// @notice Creates a new category.
    /// @param tokenURI The URI of the category.
    /// @return The new category's token ID.
    function createTrigger(
        string memory tokenURI,
        address owner
    ) external returns (uint256);
}

/// @title A contract for interacting with the Year contract.
interface IYear {
    /// @notice Creates a new year.
    /// @param tokenURI The URI of the year.
    /// @return The new year's token ID.
    function createYear(
        string memory tokenURI,
        address owner
    ) external returns (uint256);
}

/// @title A contract for interacting with the DayMonth contract.
interface IDayMonth {
    /// @notice Creates a new day and month.
    /// @param tokenURI The URI of the day and month.
    /// @return The new day and month's token ID.
    function createDayMonth(
        string memory tokenURI,
        address owner
    ) external returns (uint256);
}

/// @title A contract for interacting with the CardPack contract.
interface ICardPack {
    /// @notice Creates a new card pack.
    /// @param tokenURI The URI of the card pack.
    /// @return The new card pack's token ID.
    function createCard(
        string memory tokenURI,
        address owner
    ) external returns (uint256);

    /// @notice Changes the status of a card pack to "opened".
    /// @param tokenId The ID of the token to be opened.
    function changeToOpened(uint256 tokenId) external;

    /// @notice Fetches the owner of a particular card pack.
    /// @param tokenId The ID of the token whose owner is to be fetched.
    /// @return The address of the owner.
    function ownerOf(uint256 tokenId) external view returns (address);

    /// @notice Checks whether a card pack is opened or not.
    /// @param tokenId The ID of the token to be checked.
    /// @return A boolean indicating whether the card pack is opened or not.
    function isOpened(uint256 tokenId) external view returns (bool);
}

/// @title A contract to manage the opening of card packs.
/// @dev The contract is Ownable, and the owner has exclusive rights to perform certain actions.
contract PackOpenerV1 is Initializable, OwnableUpgradeable {
    ICardPack cardPackContract;
    IDayMonth dayMonthContract;
    IYear yearContract;
    ICategory categoryContract;
    ICraftingCard craftingCardContract;
    ITrigger triggerContract;
    address adminWallet;

    /// @notice Initializes the contract with specified parameters.
    /// @dev The contract takes the number of each type of card, and the addresses of the card contracts.
    /// @param _cardPackAddress The address of the CardPack contract.
    /// @param _dayMonthAddress The address of the DayMonth contract.
    /// @param _yearAddress The address of the Year contract.
    /// @param _categoryAddress The address of the Category contract.
    /// @param _craftingCardContract The address of the CraftingCard contract.    

    function initialize(
        address _cardPackAddress,
        address _dayMonthAddress,
        address _yearAddress,
        address _categoryAddress,
        address _craftingCardContract,
        address _triggerContract
    ) public initializer {
        cardPackContract = ICardPack(_cardPackAddress);
        dayMonthContract = IDayMonth(_dayMonthAddress);
        yearContract = IYear(_yearAddress);
        categoryContract = ICategory(_categoryAddress);
        craftingCardContract = ICraftingCard(_craftingCardContract);
        triggerContract = ITrigger(_triggerContract);
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
    /// @param _dayMonthAmount The number of DayMonth cards in a pack.
    /// @param _yearAmount The number of Year cards in a pack.
    /// @param _categoryAmount The number of Category cards in a pack.
    /// @param _craftingCardAmount The number of CraftingCard cards in a pack.
    /// @param _tokenDayMonth The URI for the DayMonth cards.
    /// @param _tokenYear The URI for the Year cards.
    /// @param _tokenCategory The URI for the Category cards.
    /// @param _tokenCraftingCard The URI for the CraftingCard cards.
    function openPack(
        uint256 _tokenId,
        uint256 _dayMonthAmount,
        uint256 _yearAmount,
        uint256 _categoryAmount,
        uint256 _craftingCardAmount,
        uint256 _triggerAmount,
        string[] memory _tokenDayMonth,
        string[] memory _tokenYear,
        string[] memory _tokenCategory,
        string[] memory _tokenCraftingCard,
        string[] memory _tokenTrigger
    ) public onlyAdmin onlyNonOpenedPacks(_tokenId) {
        require(
            _dayMonthAmount == _tokenDayMonth.length,
            "Day Month Amount != URIs.length"
        );
        require(_yearAmount == _tokenYear.length, "Year Amount != URIs.length");
        require(
            _categoryAmount == _tokenCategory.length,
            "Category Amount != URIs.length"
        );
        require(
            _craftingCardAmount == _tokenCraftingCard.length,
            "CraftingCard Amount != URIs.length"
        );
        require(
            _triggerAmount == _tokenTrigger.length,
            "Trigger Amount != URIs.length"
        );
        address tokenOwner = cardPackContract.ownerOf(_tokenId);
        cardPackContract.changeToOpened(_tokenId);
        for (uint256 i = 0; i < _dayMonthAmount; i++) {
            dayMonthContract.createDayMonth(_tokenDayMonth[i], tokenOwner);
        }
        for (uint256 j = 0; j < _yearAmount; j++) {
            yearContract.createYear(_tokenYear[j], tokenOwner);
        }
        for (uint256 k = 0; k < _categoryAmount; k++) {
            categoryContract.createCategory(_tokenCategory[k], tokenOwner);
        }
        for (uint256 l = 0; l < _craftingCardAmount; l++) {
            craftingCardContract.createCraftingCard(
                _tokenCraftingCard[l],
                tokenOwner
            );
        }
        for (uint256 m = 0; m < _triggerAmount; m++) {
            triggerContract.createTrigger(_tokenTrigger[m], tokenOwner);
        }
    }
    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }
}
