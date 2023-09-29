// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("Category", () => {
//   let category: Contract;
//   let owner: Signer;
//   let minterContract: Signer;
//   let craftingContract: Signer;
//   let owner2: Signer;

//   beforeEach(async () => {
//     [owner, minterContract, craftingContract, owner2] =
//       await ethers.getSigners();
//     const Category = await ethers.getContractFactory("Category");
//     category = await Category.connect(owner).deploy();
//     await category.deployed();
//     const addressOwner = await owner.getAddress();
//     await category.connect(owner).changeMinter(addressOwner);
//     await category.connect(owner).changeCrafter(addressOwner);
//   });

//   describe("createCategory", () => {
//     it("should create a new category token", async () => {
//       const tokenURI = "https://example.com/token/1";
//       const tokenCounter = await category.tokenCounter();

//       await expect(category.createCategory("https://example.com/token/1", await owner.getAddress()))
//         .to.emit(category, "CardPartCreated")
//         .withArgs(tokenCounter+1, tokenURI);

//       const tokenURIResult = await category.tokenURI(tokenCounter);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });

//     it("should revert when creating a category with an empty token URI", async () => {
//       const emptyTokenURI = "";

//       await expect(category.createCategory(emptyTokenURI)).to.be.reverted;
//     });
//   });

//   describe("changeToCrafted", () => {

//     it("should revert when a non-crafting contract tries to change the status of a category", async () => {
//       const tokenId = 0;
//       await category.createCategory(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );
//       await category.connect(owner).changeCrafter(await owner.getAddress());
//       await expect(
//         category.connect(owner2).changeToCrafted(tokenId)
//       ).to.be.revertedWith("You are not the Crafting Contract :)");
//     });

//     it("should revert when trying to change the status of a nonexistent category", async () => {
//       await expect(category.connect(craftingContract).changeToCrafted(1)).to.be
//         .reverted;
//     });
//   });

//   describe("changeMinter", () => {
//     it("should revert when a non-owner tries to change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();

//       await expect(
//         category.connect(owner2).changeMinter(newMinter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a MinterContractChanged event when the minter contract address is changed", async () => {
//       const newMinter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(category.connect(owner).changeMinter(newMinter))
//         .to.emit(category, "MinterContractChanged")
//         .withArgs(oldOwner, newMinter);
//     });
//     it("should revert when changing the minter contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(category.connect(owner).changeMinter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("changeCrafter", () => {
//     it("should revert when a non-owner tries to change the crafting contract address", async () => {
//       const newCrafter = await owner2.getAddress();
//       await expect(
//         category.connect(owner2).changeCrafter(newCrafter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a CraftingContractChanged event when the crafting contract address is changed", async () => {
//       const newCrafter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(category.connect(owner).changeCrafter(newCrafter))
//         .to.emit(category, "CraftingContractChanged")
//         .withArgs(oldOwner, newCrafter);
//     });

//     it("should revert when changing the crafting contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(category.connect(owner).changeCrafter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("isCrafted", () => {

//     it("should return false if a category is not crafted", async () => {
//       const tokenId = 0;
//       await category.createCategory(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );

//       const isCraftedResult = await category.isCrafted(tokenId);
//       expect(isCraftedResult).to.equal(false);
//     });

//     it("should revert when checking the status of a nonexistent category", async () => {
//       await expect(category.isCrafted(0)).to.be.reverted;
//     });
//   });

//   describe("tokenURI", () => {
//     it("should return the correct URI for a given token ID", async () => {
//       const tokenId = 0;
//       const tokenURI = "https://example.com/token/1";
//       await category.createCategory(
//         "https://example.com/token/1",
//         await owner.getAddress()
//       );
//       const result = await category.tokenURI(tokenId);
//       expect(result).to.equal(tokenURI);
//     });

//     it("should revert when querying URI for a nonexistent token", async () => {
//       await expect(category.tokenURI(0)).to.be.reverted;
//     });
//   });
// });