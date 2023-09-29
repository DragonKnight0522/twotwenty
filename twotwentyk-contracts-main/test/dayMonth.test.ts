// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("DayMonth", () => {
//   let dayMonth: Contract;
//   let owner: Signer;
//   let minterContract: Signer;
//   let craftingContract: Signer;
//   let owner2: Signer;

//   beforeEach(async () => {
//     [owner, minterContract, craftingContract, owner2] =
//       await ethers.getSigners();
//     const DayMonth = await ethers.getContractFactory("DayMonth");
//     dayMonth = await DayMonth.connect(owner).deploy();
//     await dayMonth.deployed();
//     const addressOwner = await owner.getAddress();
//     await dayMonth.connect(owner).changeMinter(addressOwner);
//     await dayMonth.connect(owner).changeCrafter(addressOwner);
//   });

//   describe("createDayMonth", () => {
//     it("should create a new day-month token", async () => {
//       const tokenURI = "https://example.com/token/1";

//       await expect(await dayMonth.createDayMonth("https://example.com/token/1",await owner.getAddress()))
//         .to.emit(dayMonth, "Transfer")
//         .withArgs(ethers.constants.AddressZero, await owner.getAddress(), 0);

//       const tokenURIResult = await dayMonth.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });

//     it("should revert when creating a day-month token with an empty token URI", async () => {
//       const emptyTokenURI = "";

//       await expect(dayMonth.createDayMonth(emptyTokenURI)).to.be.reverted;
//     });
//   });

//   describe("changeToCrafted", () => {

//     it("should revert when a non-crafting contract tries to change the status of a day-month token", async () => {
//       await dayMonth.createDayMonth("https://example.com/token/1",await owner.getAddress());
//       await dayMonth.connect(owner).changeCrafter(await owner.getAddress());
//       await expect(
//         dayMonth.connect(owner2).changeToCrafted(0)
//       ).to.be.revertedWith("You are not the Crafting Contract :)");
//     });

//     it("should revert when trying to change the status of a nonexistent day-month token", async () => {
//       await expect(dayMonth.connect(craftingContract).changeToCrafted(1)).to.be
//         .reverted;
//     });
//   });

//   describe("changeMinter", () => {
//     it("should revert when a non-owner tries to change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();

//       await expect(
//         dayMonth.connect(owner2).changeMinter(newMinter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a MinterContractChanged event when the minter contract address is changed", async () => {
//       const newMinter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(dayMonth.connect(owner).changeMinter(newMinter))
//         .to.emit(dayMonth, "MinterContractChanged")
//         .withArgs(oldOwner, newMinter);
//     });

//     it("should revert when changing the minter contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(dayMonth.connect(owner).changeMinter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("changeCrafter", () => {
//     it("should revert when a non-owner tries to change the crafting contract address", async () => {
//       const newCrafter = await owner2.getAddress();

//       await expect(
//         dayMonth.connect(owner2).changeCrafter(newCrafter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a CraftingContractChanged event when the crafting contract address is changed", async () => {
//       const newCrafter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(dayMonth.connect(owner).changeCrafter(newCrafter))
//         .to.emit(dayMonth, "CraftingContractChanged")
//         .withArgs(oldOwner, newCrafter);
//     });

//     it("should revert when changing the crafting contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(dayMonth.connect(owner).changeCrafter(zeroAddress)).to.be
//         .reverted;
//     });
//   });

//   describe("isCrafted", () => {

//     it("should return false if a day-month token is not crafted", async () => {
//       await dayMonth.createDayMonth("https://example.com/token/1",await owner.getAddress());

//       const isCraftedResult = await dayMonth.isCrafted(0);
//       expect(isCraftedResult).to.equal(false);
//     });

//     it("should revert when checking the status of a nonexistent day-month token", async () => {
//       await expect(dayMonth.isCrafted(0)).to.be.reverted;
//     });
//   });

//   describe("tokenURI", () => {
//     it("should return the correct URI for a given token ID", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await dayMonth.createDayMonth("https://example.com/token/1",await owner.getAddress());

//       const result = await dayMonth.tokenURI(0);
//       expect(result).to.equal(tokenURI);
//     });

//     it("should revert when querying URI for a nonexistent token", async () => {
//       await expect(dayMonth.tokenURI(0)).to.be.reverted;
//     });
//   });
// });
