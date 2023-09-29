// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
@title CardPack
@dev This contract represents a Card Pack ERC721 token contract, allowing users to create and manage card packs.
*/
contract CardPack is ERC721, Ownable {
    uint256 totalCounter;
    uint256 public tokenCounterStandard; // Counter for the token IDs
    uint256 public tokenCounterPremium; // Counter for the token IDs
    uint256 public tokenCounterElite; // Counter for the token IDs
    mapping(uint256 => bool) public opened; // Mapping to track if a card pack has been opened
    mapping(uint256 => string) private _tokenURIs; // Mapping to store the token URIs for each card pack
    address public minterContract; // Address of the minter contract
    uint256 public totalLimit;
    uint256 public standardLimit;
    uint256 public premiumLimit;
    uint256 public eliteLimit;
    address adminWallet;

    event CardPackCreated(
        uint256 indexed tokenId,
        uint256 packNumber,
        string tokenURI,
        string indexed typeOf
    );
    event CardPackOpened(uint256 indexed tokenId);
    event MinterContractChanged(
        address indexed previousMinter,
        address indexed newMinter
    );
    event LimitChanged(uint256 newLimit, string typeOf);

    /**
    @dev Initializes the CardPack contract by setting the token name and symbol.
    */   
    constructor(
        uint256 _standardAmount,
        uint256 _premiumAmount,
        uint256 _eliteAmount,
        uint256 _totalAmount
    ) ERC721("Card Pack", "CP") {            
        standardLimit = _standardAmount;
        premiumLimit = _premiumAmount;
        eliteLimit = _eliteAmount;        
        totalLimit = _totalAmount;        
        
        totalCounter = 0;
        tokenCounterStandard = 0;
        tokenCounterPremium = 0;
        tokenCounterElite = 0;
    }

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin wallet");
        _;
    }

    /**
    @dev Modifier to check if the sender is the minter contract.
    */
    modifier onlyMinterContract() {
        require(
            msg.sender == minterContract,
            "Card Pack: You are not the Minting Contract :)"
        );
        _;
    }

    /**
     * @dev Modifier to check if a card pack with the given token ID exists.
     * @param tokenId The ID of the token to check.
     */
    modifier packExists(uint256 tokenId) {
        require(_exists(tokenId), "CardPack: Pack does not exist");
        _;
    }

    modifier tokenURICheck(string memory _tokenURI) {
        require(bytes(_tokenURI).length > 0, "TokenURI should not be empty");
        _;
    }

    modifier checkLimit() {
        require(totalLimit > totalCounter, "All the tokens have been minted");
        _;
    }

    /**
    @dev Internal function to handle burning of a token.
    @param tokenId The ID of the token to be burned.
    */
    function _burn(uint256 tokenId) internal override(ERC721) {
        super._burn(tokenId);
    }

    /**
    @dev Returns the URI for a given token ID.
    @param tokenId The ID of the token to query the URI for.
    @return A string representing the URI for the given token ID.
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

    /**
    @dev Creates a new card pack token.
    @param _tokenURI The URI for the token.
    @return The ID of the newly created token.
    */
    function createStandardCard(
        string memory _tokenURI
    ) public tokenURICheck(_tokenURI) checkLimit onlyAdmin returns (uint256) {
        require(
            standardLimit > tokenCounterStandard,
            "All Standard tokens have been minted"
        );
        _safeMint(msg.sender, totalCounter);
        _tokenURIs[totalCounter] = _tokenURI;
        opened[totalCounter] = false;
        tokenCounterStandard++;
        totalCounter++;
        emit CardPackCreated(
            tokenCounterStandard - 1,
            totalCounter - 1,
            _tokenURI,
            "Standard"
        );
        return totalCounter;
    }

    /**
    @dev Creates a new card pack token.
    @param _tokenURI The URI for the token.
    @return The ID of the newly created token.
    */
    function createPremiumCard(
        string memory _tokenURI
    ) public tokenURICheck(_tokenURI) checkLimit onlyAdmin returns (uint256) {
        require(
            premiumLimit > tokenCounterPremium,
            "All Premium tokens have been minted"
        );
        _safeMint(msg.sender, totalCounter);
        _tokenURIs[totalCounter] = _tokenURI;
        opened[totalCounter] = false;
        tokenCounterPremium++;
        totalCounter++;
        emit CardPackCreated(
            tokenCounterPremium - 1,
            totalCounter - 1,
            _tokenURI,
            "Premium"
        );
        return totalCounter;
    }

    /**
    @dev Creates a new card pack token.
    @param _tokenURI The URI for the token.
    @return The ID of the newly created token.
    */
    function createEliteCard(
        string memory _tokenURI
    ) public tokenURICheck(_tokenURI) checkLimit onlyAdmin returns (uint256) {
        require(
            eliteLimit > tokenCounterElite,
            "All Elite tokens have been minted"
        );
        _safeMint(msg.sender, tokenCounterElite);
        _tokenURIs[tokenCounterElite] = _tokenURI;
        opened[tokenCounterElite] = false;
        tokenCounterElite++;
        totalCounter++;
        emit CardPackCreated(
            tokenCounterElite - 1,
            totalCounter - 1,
            _tokenURI,
            "Elite"
        );
        return totalCounter;
    }

    /**
    @dev Changes the status of a card pack to opened.
    @param tokenId The ID of the token to be marked as opened.
    */
    function changeToOpened(
        uint256 tokenId
    ) public onlyMinterContract packExists(tokenId) {
        opened[tokenId] = true;
        _burn(tokenId);
        emit CardPackOpened(tokenId);
    }

    /**
    @dev Checks if a card pack has been opened.
    @param tokenId The ID of the token to check.
    @return A boolean indicating whether the card pack has been opened.
    */
    function isOpened(
        uint256 tokenId
    ) public view packExists(tokenId) returns (bool) {
        return opened[tokenId];
    }

    /**
     * @dev Changes the minter contract address.
     * @param _newMinter The new address of the minter contract.
     */
    function changeMinter(address _newMinter) public onlyOwner {
        require(_newMinter != address(0), "CardPack: No zero address");
        emit MinterContractChanged(minterContract, _newMinter);
        minterContract = _newMinter;
    }

    function changeToTotalLimit(uint256 newLimit) public onlyOwner {
        require(newLimit > totalLimit, "New limit should be higher");
        totalLimit = newLimit;
        emit LimitChanged(newLimit, "Standard");
    }

    function changeToStandardLimit(uint256 newLimit) public onlyOwner {
        require(newLimit > standardLimit, "New limit should be higher");
        standardLimit = newLimit;
        emit LimitChanged(newLimit, "Standard");
    }

    function changeToPremiumLimit(uint256 newLimit) public onlyOwner {
        require(newLimit > premiumLimit, "New limit should be higher");
        premiumLimit = newLimit;
        emit LimitChanged(newLimit, "Premium");
    }

    function changeToEliteLimit(uint256 newLimit) public onlyOwner {
        require(newLimit > eliteLimit, "New limit should be higher");
        eliteLimit = newLimit;
        emit LimitChanged(newLimit, "Elite");
    }

    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }
}
