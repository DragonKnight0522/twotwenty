// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title DayMonth
 * @dev This contract represents a DayMonth ERC721 token contract, allowing users to create and manage day-month tokens.
 */
contract DayMonth is ERC721, Ownable {    
    
    uint256 public tokenCounter; // Counter for the token IDs
    address private minterContract; // Address of the minter contract
    address private craftingContract; // Address of the crafting contract
    uint256 public collectionId;    
    /**
     * @dev Stores the token URIs for each token ID.
     */
    mapping(uint256 => string) private tokenURIs;

    /**
     * @dev Tracks whether a token has been crafted.
     */
    mapping(uint256 => bool) private crafted;

    /**
     * @dev Emitted when a card part is created.
     * @param tokenId The ID of the token.
     * @param tokenURI The URI associated with the token.
     */
    event DayMonthMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI);

    /**
     * @dev Emitted when a card part is marked as crafted.
     * @param tokenId The ID of the token.
     */
    event PartCrafted(uint256 indexed tokenId);

    /**
     * @dev Emitted when the minter contract address is changed.
     * @param previousMinter The address of the previous minter contract.
     * @param newMinter The address of the new minter contract.
     */
    event MinterContractChanged(
        address indexed previousMinter,
        address indexed newMinter
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
    @dev Initializes the Category contract by setting the token name and symbol.
    */
    constructor(uint256 _collectionId) 
        ERC721("Day Month", "DMT") {            
            collectionId = _collectionId;            
    }

    modifier tokenExists(uint256 tokenId) {
        require(_exists(tokenId), "Category: does not exist");
        _;
    }

    /**
     * @dev Modifier to check if the sender is the crafting contract.
     */
    modifier onlyCrafting() {
        require(
            msg.sender == craftingContract,
            "You are not the Crafting Contract :)"
        );
        _;
    }

    /**
     * @dev Modifier to check if the sender is the minter contract.
     */
    modifier onlyMinter() {
        require(
            msg.sender == minterContract,
            "DayMonth: You are not the Minting Contract :)"
        );
        _;
    }

    /**
     * @dev Internal function to handle burning of a token.
     * @param tokenId The ID of the token to be burned.
     */
    function _burn(uint256 tokenId) internal override(ERC721) {
        super._burn(tokenId);
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
            return tokenURIs[tokenId];
        } else {
            return string(abi.encodePacked(baseURI, tokenURIs[tokenId]));
        }
    }

    /**
     * @dev Creates a new day-month token.
     * @param _tokenURI The URI for the token.     
     * @param _owner The owner account.     
     */
    function mintDayMonth(
        string memory _tokenURI,
        address _owner
    ) public onlyMinter returns(uint256){
        require(bytes(_tokenURI).length > 0, "TokenURI should not be empty");
        tokenURIs[tokenCounter] = _tokenURI;
        _safeMint(_owner, tokenCounter);       

        tokenCounter++;

        return tokenCounter - 1;
    }

    /**
     * @dev Changes the minter contract address.
     * @param _newMinter The new address of the minter contract.
     */
    function changeMinter(address _newMinter) public onlyOwner {
        require(_newMinter != address(0), "Minter: No zero address");
        address oldMinter = minterContract;
        minterContract = _newMinter;
        emit MinterContractChanged(oldMinter, _newMinter);
    }

    /**
    @dev Changes the crafting contract address.
    @param _newCrafter The new address of the crafting contract.
    */
    function changeCrafter(address _newCrafter) public onlyOwner {
        require(_newCrafter != address(0), "Crafting: No zero address");
        address oldCrafter = craftingContract;
        craftingContract = _newCrafter;
        emit CraftingContractChanged(oldCrafter, _newCrafter);
    }

    /**
    @dev Changes the status of a card pack to crafted.
    @param tokenId The ID of the token to be marked as crafted.
*/
    function changeToCrafted(
        uint256 tokenId
    ) public onlyCrafting tokenExists(tokenId) {
        crafted[tokenId] = true;
        _burn(tokenId);
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
}
