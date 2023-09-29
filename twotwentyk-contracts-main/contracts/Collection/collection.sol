// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "hardhat/console.sol";

import "../Interfaces/factoryInterface.sol";

/**
@title Category
@dev This contract represents a Category ERC721 token contract, allowing users to create and manage categories.
*/
contract Collection is Initializable, ERC721Upgradeable, OwnableUpgradeable{
    uint256 public tokenCounter; // Counter for the token IDs
    address private minterContract; // Address of the minter contract    
    IPackFactory packFactory;
    ICategoryFactory categoryFactory;
    IYearFactory yearFactory;
    IDayMonthFactory dayMonthFactory;
    ICraftingFactory craftingFactory;
    ITriggerFactory triggerFactory;
    /**
     * @dev Stores the token URIs for each token ID.
     */
    mapping(uint256 => string) private _tokenURIs;    

    /**
     * @dev Emitted when a card part is created.
     * @param tokenId The ID of the token.
     * @param tokenURI The URI associated with the token.
     */
    event CollectionCreated(uint256 tokenId, string tokenURI, address pack, address category, address year, address daymonth, address crafting, address trigger);

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
    @dev Initializes the Category contract by setting the token name and symbol.
    */    
    function initialize(
        address _packFactory,
        address _categoryFactory,
        address _yearFactory,
        address _dayMonthFactory,
        address _craftingFactory,
        address _triggerFactory
    ) external initializer  {  
        packFactory = IPackFactory(_packFactory);
        categoryFactory = ICategoryFactory(_categoryFactory);
        yearFactory = IYearFactory(_yearFactory);
        dayMonthFactory = IDayMonthFactory(_dayMonthFactory);
        craftingFactory = ICraftingFactory(_craftingFactory);
        triggerFactory = ITriggerFactory(_triggerFactory);
        __ERC721_init("Collection", "CLT");
        __Ownable_init();
    }
    

    modifier tokenExists(uint256 tokenId) {
        require(_exists(tokenId), "Category: does not exist");
        _;
    }

    /**
    @dev Modifier to check if the sender is the minter contract.
    */
    modifier onlyMinter() {
        require(
            msg.sender == minterContract,
            "Category: You are not the Minting Contract :)"
        );
        _;
    }    

    /**
    @dev Internal function to handle burning of a token.
    @param tokenId The ID of the token to be burned.
    */
    function _burn(
        uint256 tokenId
    ) internal virtual override(ERC721Upgradeable) tokenExists(tokenId) {
        super._burn(tokenId);
    }

    /**
    @dev Returns the URI for a given token ID.
    @param tokenId The ID of the token to query the URI for.
    @return A string representing the URI for the given token ID.
    */
    function tokenURI(
        uint256 tokenId
    )
        public
        view
        virtual
        override
        tokenExists(tokenId)
        returns (string memory)
    {
        require(_exists(tokenId), "URI query for nonexistent token");

        string memory baseURI = _baseURI();
        if (bytes(baseURI).length == 0) {
            return _tokenURIs[tokenId];
        } else {
            return string(abi.encodePacked(baseURI, _tokenURIs[tokenId]));
        }
    }

    /**
    @dev Creates a new category token.
    @param _tokenURI The URI for the token.
    @return The ID of the newly created token.Good GG
    */
    function createCollection(
        string memory _tokenURI,
        uint256 _standardAmount,
        uint256 _premiumAmount,
        uint256 _elliteAmount,
        uint256 _totalAmount 
    ) public onlyMinter payable returns (uint256, string memory) {
        require(bytes(_tokenURI).length > 0, "TokenURI should not be empty");
        _tokenURIs[tokenCounter] = _tokenURI;
        _safeMint(msg.sender, tokenCounter);       

            // create card pack collection
            (address pack) = packFactory.createCardPack(_standardAmount, _premiumAmount, _elliteAmount, _totalAmount, msg.sender);

            // create category collection
            (address cateogry) = categoryFactory.createCategory(tokenCounter, msg.sender);
            

            // create year collection
            (address year) = yearFactory.createYear(tokenCounter, msg.sender);

            // create daymonth collection
            (address daymonth) = dayMonthFactory.createDayMonth(tokenCounter, msg.sender);

            // create crafting collection
            (address crafting) = craftingFactory.createCrafting(tokenCounter, msg.sender);
        
            // create trigger collection
            (address trigger) = triggerFactory.createTrigger(tokenCounter, msg.sender);

        console.log("crafting", crafting);
        console.log("trigger", trigger);
        emit CollectionCreated(tokenCounter, _tokenURI, pack, cateogry, year, daymonth, crafting, trigger);

        tokenCounter++;
        return (tokenCounter - 1, _tokenURI);
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
}
