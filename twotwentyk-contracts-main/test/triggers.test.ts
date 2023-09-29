// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("Trigger", () => {
//   let trigger: Contract;
//   let owner: Signer;
//   let minterContract: Signer;
//   let craftingContract: Signer;
//   let owner2: Signer;

//   beforeEach(async () => {
//     [owner, minterContract, craftingContract, owner2] =
//       await ethers.getSigners();
//     const Trigger = await ethers.getContractFactory("Trigger");
//     trigger = await Trigger.connect(owner).deploy();
//     await trigger.deployed();
//     const addressOwner = await owner.getAddress();
//     await trigger.connect(owner).changeMinter(addressOwner);
//     await trigger.connect(owner).changeCrafter(addressOwner);
//   });

//   describe("createTrigger", () => {
//     it("should create a new card part token", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await expect(
//         await trigger.createTrigger(tokenURI, await owner.getAddress())
//       )
//         .to.emit(trigger, "Transfer")
//         .withArgs(ethers.constants.AddressZero, await owner.getAddress(), 0);

//       const tokenURIResult = await trigger.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });

//     it("should revert when creating a card part token with an empty token URI", async () => {
//       const emptyTokenURI = "";
//       await expect(trigger.createTrigger(emptyTokenURI)).to.be.reverted;
//     });
//   });

//   describe("changeToCrafted", () => {

//     it("should revert when a non-crafting contract tries to change the status of a card part", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await trigger.createTrigger(tokenURI, await owner.getAddress());
//       await trigger.connect(owner).changeCrafter(await owner.getAddress());
//       await expect(
//         trigger.connect(owner2).changeToCrafted(0)
//       ).to.be.revertedWith("You are not the Crafting Contract :)");
//     });

//     it("should revert when trying to change the status of a nonexistent card part", async () => {
//       await expect(trigger.connect(craftingContract).changeToCrafted(1)).to.be
//         .reverted;
//     });
//   });

//   describe("changeMinter", () => {
//     it("should revert when a non-owner tries to change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();
//       await expect(
//         trigger.connect(owner2).changeMinter(newMinter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a MinterContractChanged event when the minter contract address is changed", async () => {
//       const newMinter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(trigger.connect(owner).changeMinter(newMinter))
//         .to.emit(trigger, "MinterContractChanged")
//         .withArgs(oldOwner, newMinter);
//     });

//     it("should revert when changing the minter contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;
//       await expect(trigger.connect(owner).changeMinter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("changeCrafter", () => {
//     it("should revert when a non-owner tries to change the crafting contract address", async () => {
//       const newCrafter = await owner2.getAddress();
//       await expect(
//         trigger.connect(owner2).changeCrafter(newCrafter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a CraftingContractChanged event when the crafting contract address is changed", async () => {
//       const newCrafter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(trigger.connect(owner).changeCrafter(newCrafter))
//         .to.emit(trigger, "CraftingContractChanged")
//         .withArgs(oldOwner, newCrafter);
//     });

//     it("should revert when changing the crafting contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;
//       await expect(trigger.connect(owner).changeCrafter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("isCrafted", () => {

//     it("should return false if a card part is not crafted", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await trigger.createTrigger(tokenURI, await owner.getAddress());
//       const isCraftedResult = await trigger.isCrafted(0);
//       expect(isCraftedResult).to.equal(false);
//     });

//     it("should revert when checking the status of a nonexistent card part", async () => {
//       await expect(trigger.isCrafted(0)).to.be.reverted;
//     });
//   });

//   describe("tokenURI", () => {
//     it("should return the correct URI for a given token ID", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await trigger.createTrigger(tokenURI, await owner.getAddress());
//       const result = await trigger.tokenURI(0);
//       expect(result).to.equal(tokenURI);
//     });

//     it("should revert when querying URI for a nonexistent token", async () => {
//       await expect(trigger.tokenURI(0)).to.be.reverted;
//     });
//   });
// });
