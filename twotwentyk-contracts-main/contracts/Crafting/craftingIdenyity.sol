// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface IPart {
    function ownerOf(uint256 tokenId) external view returns (address);

    /**
     * @dev Checks if a token has been marked as crafted.
     * @param tokenId The ID of the token to check.
     * @return A boolean indicating if the token is crafted.
     */
    function isCrafted(uint256 tokenId) external view returns (bool);

    /**
     * @dev Changes the status of a card pack to crafted.
     * @param tokenId The ID of the token to be marked as crafted.
     */
    function changeToCrafted(uint256 tokenId) external;
}

contract CollectionCrafting is Initializable, ERC721Upgradeable, OwnableUpgradeable {
    IPart public dayMonthContract;
    IPart public yearContract;
    IPart public categoryContract;
    IPart public craftingCardContract;
    address private craftingContract; // Address of the crafting contract
    address adminWallet;
    uint256 private collectionCounter;
    mapping(uint256 => uint256) private dayMonthTokenIdByCollection;
    mapping(uint256 => uint256) private yearTokenIdByCollection;
    mapping(uint256 => uint256) private categoryTokenIdByCollection;
    mapping(uint256 => uint256) private craftingCardTokenIdByCollection;
    mapping(uint256 => bool) public crafted;
    /**
     * @dev Stores the token URIs for each token ID.
     */
    mapping(uint256 => string) private _tokenURIs;
    event CollectionCrafted(
        uint256 indexed collectionId,
        uint256 dayMonthTokenId,
        uint256 yearTokenId,
        uint256 categoryTokenId,
        uint256 craftingCardTokenId,
        string tokenURI
    );
    /**
     * @dev Emitted when the crafting contract address is changed.
     * @param previousCrafter The address of the previous crafting contract.
     * @param newCrafter The address of the new crafting contract.
     */
    event CraftingContractChanged(
        address indexed previousCrafter,
        address indexed newCrafter
    );
        /**
     * @dev Emitted when a card part is marked as crafted.
     * @param tokenId The ID of the token.
     */
    event PartCrafted(uint256 indexed tokenId);    

    function initialize(
        
    ) public initializer {        
        adminWallet = msg.sender;
        collectionCounter = 0;
        
        __ERC721_init("Collection", "COLL");
        __Ownable_init();                
    }

    /**
    @dev Modifier to check if the sender is the crafting contract.
    */
    modifier onlyCrafting() {
        require(
            msg.sender == craftingContract,
            "You are not the Crafting Contract :)"
        );
        _;
    }

    modifier tokenExists(uint256 tokenId) {
        require(_exists(tokenId), "Category: does not exist");
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin wallet");
        _;
    }

    /**
     * @dev Returns the URI for a given token ID.
     * @param tokenId The ID of the token to query the URI for.
     * @return A string representing the URI for the given token ID.
     */
    function tokenURI(
        uint256 tokenId
    ) public view override returns (string memory) {
        require(_exists(tokenId), "URI query for nonexistent token");
        string memory baseURI = _baseURI();
        if (bytes(baseURI).length == 0) {
            return _tokenURIs[tokenId];
        } else {
            return string(abi.encodePacked(baseURI, _tokenURIs[tokenId]));
        }
    }

    function craftCollection(
        address _dayMonthContract,
        address _yearContract,
        address _categoryContract,
        address _craftingCardContract,

        uint256 dayMonthTokenId,
        uint256 yearTokenId,
        uint256 categoryTokenId,
        uint256 craftingCardTokenId,
        string memory _tokenURI
    ) public onlyAdmin {
        
        dayMonthContract = IPart(_dayMonthContract);
        yearContract = IPart(_yearContract);
        categoryContract = IPart(_categoryContract);
        craftingCardContract = IPart(_craftingCardContract);

        // Check ownership of NFTs
        address dayMonthOwner = dayMonthContract.ownerOf(dayMonthTokenId);
        address yearOwner = yearContract.ownerOf(yearTokenId);
        address categoryOwner = categoryContract.ownerOf(categoryTokenId);
        address craftingCardOwner = craftingCardContract.ownerOf(
            craftingCardTokenId
        );
        // Check ownership of NFTs
        bool dayMonthBool = dayMonthContract.isCrafted(dayMonthTokenId);
        bool yearBool = yearContract.isCrafted(yearTokenId);
        bool categoryBool = categoryContract.isCrafted(categoryTokenId);
        bool craftingCardBool = craftingCardContract.isCrafted(
            craftingCardTokenId
        );

        require(
            dayMonthOwner == yearOwner &&
                yearOwner == categoryOwner &&
                categoryOwner == craftingCardOwner,
            "Only the owner of all tokens can craft a new collection"
        );
        require(dayMonthBool == false, "Day/Month already used");
        require(yearBool == false, "Year already used");
        require(categoryBool == false, "Category already used");
        require(craftingCardBool == false, "Crafting Card already used");

        uint256 collectionId = collectionCounter;
        _safeMint(dayMonthOwner, collectionId);

        dayMonthContract.changeToCrafted(dayMonthTokenId);
        yearContract.changeToCrafted(yearTokenId);
        categoryContract.changeToCrafted(categoryTokenId);
        craftingCardContract.changeToCrafted(craftingCardTokenId);

        dayMonthTokenIdByCollection[collectionId] = dayMonthTokenId;
        yearTokenIdByCollection[collectionId] = yearTokenId;
        categoryTokenIdByCollection[collectionId] = categoryTokenId;
        craftingCardTokenIdByCollection[collectionId] = craftingCardTokenId;

        _tokenURIs[collectionId] = _tokenURI;
        crafted[collectionId] = false;
        collectionCounter++;

        emit CollectionCrafted(
            collectionId,
            dayMonthTokenId,
            yearTokenId,
            categoryTokenId,
            craftingCardTokenId,
            _tokenURI
        );
    }

    function getCollectionTokens(
        uint256 collectionId
    ) public view returns (uint256, uint256, uint256, uint256) {
        require(_exists(collectionId), "Collection does not exist");
        return (
            dayMonthTokenIdByCollection[collectionId],
            yearTokenIdByCollection[collectionId],
            categoryTokenIdByCollection[collectionId],
            craftingCardTokenIdByCollection[collectionId]
        );
    }

        /**
    @dev Changes the crafting contract address.
    @param _newCrafter The new address of the crafting contract.
    */
    function changeCrafter(address _newCrafter) public onlyOwner {
        require(_newCrafter != address(0), "Crafting Address: No zero address");
        craftingContract = _newCrafter;
        emit CraftingContractChanged(craftingContract, _newCrafter);
    }

    /**
    @dev Changes the status of a card pack to crafted.
    @param tokenId The ID of the token to be marked as crafted.
*/
    function changeToCrafted(
        uint256 tokenId
    ) public onlyCrafting tokenExists(tokenId) {
        crafted[tokenId] = true;
        emit PartCrafted(tokenId);
    }

    /**
     * @dev Checks if a token has been marked as crafted.
     * @param tokenId The ID of the token to check.
     * @return A boolean indicating if the token is crafted.
     */
    function isCrafted(
        uint256 tokenId
    ) public view tokenExists(tokenId) returns (bool) {
        return crafted[tokenId];
    }

    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }
}
