// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("CraftingCard", () => {
//   let craftingCard: Contract;
//   let owner: Signer;
//   let minterContract: Signer;
//   let craftingContract: Signer;
//   let owner2: Signer;

//   beforeEach(async () => {
//     [owner, minterContract, craftingContract, owner2] =
//       await ethers.getSigners();
//     const CraftingCard = await ethers.getContractFactory("CraftingCard");
//     craftingCard = await CraftingCard.deploy();
//     await craftingCard.deployed();
//     const addressOwner = await owner.getAddress();
//     await craftingCard.connect(owner).changeMinter(addressOwner);
//     await craftingCard.connect(owner).changeCrafter(addressOwner);
//   });

//   describe("createCraftingCard", () => {
//     it("should create a new crafting card token", async () => {
//       const tokenURI = "https://example.com/token/1";

//       await expect(
//         craftingCard.createCraftingCard(tokenURI, await owner.getAddress())
//       )
//         .to.emit(craftingCard, "CardPartCreated")
//         .withArgs(0, tokenURI);

//       const tokenURIResult = await craftingCard.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });

//     it("should revert when creating a crafting card with an empty token URI", async () => {
//       const emptyTokenURI = "";

//       await expect(craftingCard.createCraftingCard(emptyTokenURI)).to.be
//         .reverted;
//     });
//   });

//   describe("changeToCrafted", () => {

//     it("should revert when a non-crafting contract tries to change the status of a crafting card", async () => {
//       await craftingCard.createCraftingCard(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );
//       await craftingCard.connect(owner).changeCrafter(await owner.getAddress());
//       await expect(
//         craftingCard.connect(owner2).changeToCrafted(0)
//       ).to.be.revertedWith("You are not the Crafting Contract :)");
//     });

//     it("should revert when trying to change the status of a nonexistent crafting card", async () => {
//       await expect(
//         craftingCard.connect(craftingContract).changeToCrafted(1)
//       ).to.be.reverted;
//     });
//   });

//   describe("changeMinter", () => {
//     it("should revert when a non-owner tries to change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();

//       await expect(
//         craftingCard.connect(owner2).changeMinter(newMinter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a MinterContractChanged event when the minter contract address is changed", async () => {
//       const newMinter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(craftingCard.connect(owner).changeMinter(newMinter))
//         .to.emit(craftingCard, "MinterContractChanged")
//         .withArgs(oldOwner, newMinter);
//     });

//     it("should revert when changing the minter contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(
//         craftingCard.connect(owner).changeMinter(zeroAddress)
//       ).to.be.reverted;
//     });
//   });

//   describe("changeCrafter", () => {
//     it("should revert when a non-owner tries to change the crafting contract address", async () => {
//       const newCrafter = await owner2.getAddress();

//       await expect(
//         craftingCard.connect(owner2).changeCrafter(newCrafter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a CraftingContractChanged event when the crafting contract address is changed", async () => {
//       const newCrafter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(craftingCard.connect(owner).changeCrafter(newCrafter))
//         .to.emit(craftingCard, "CraftingContractChanged")
//         .withArgs(oldOwner, newCrafter);
//     });

//     it("should revert when changing the crafting contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(
//         craftingCard.connect(owner).changeCrafter(zeroAddress)
//       ).to.be.reverted;
//     });
//   });

//   describe("isCrafted", () => {

//     it("should return false if a crafting card is not crafted", async () => {
//       await craftingCard.createCraftingCard(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );

//       const isCraftedResult = await craftingCard.isCrafted(0);
//       expect(isCraftedResult).to.equal(false);
//     });

//     it("should revert when checking the status of a nonexistent crafting card", async () => {
//       await expect(craftingCard.isCrafted(0)).to.be.reverted;
//     });
//   });

//   describe("tokenURI", () => {
//     it("should return the correct URI for a given token ID", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await craftingCard.createCraftingCard(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );
//       const result = await craftingCard.tokenURI(0);
//       expect(result).to.equal(tokenURI);
//     });

//     it("should revert when querying URI for a nonexistent token", async () => {
//       await expect(craftingCard.tokenURI(0)).to.be.reverted;
//     });
//   });
// });
