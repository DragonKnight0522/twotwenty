// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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