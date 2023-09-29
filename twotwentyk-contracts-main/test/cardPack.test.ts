// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("CardPack", () => {
//   let cardPack: Contract;
//   let owner: Signer;
//   let minterContract: Signer;
//   let owner2: Signer;

//   beforeEach(async () => {
//     [owner, minterContract, owner2] = await ethers.getSigners();
//     const CardPack = await ethers.getContractFactory("CardPack");
//     cardPack = await CardPack.connect(owner).deploy(
//       100, // standardAmount
//       50, // premiumAmount
//       20, // eliteAmount
//       1000 // totalLimit
//     );
//     await cardPack.deployed();
//     const addressOwner = await owner.getAddress();
//     await cardPack.connect(owner).changeMinter(addressOwner);
//   });

//   describe("createStandardCard", () => {
//     it("should create a new standard card pack token", async () => {
//       const tokenURI = "https://example.com/token/1";

//       await expect(await cardPack.createStandardCard(tokenURI))
//         .to.emit(cardPack, "CardPackCreated")
//         .withArgs(0, 0, tokenURI, "Standard");

//       const tokenURIResult = await cardPack.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });
//   });

//   describe("createPremiumCard", () => {
//     it("should create a new premium card pack token", async () => {
//       const tokenURI = "https://example.com/token/1";

//       await expect(await cardPack.createPremiumCard(tokenURI))
//         .to.emit(cardPack, "CardPackCreated")
//         .withArgs(0, 0, tokenURI, "Premium");

//       const tokenURIResult = await cardPack.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });
//   });

//   describe("createEliteCard", () => {
//     it("should create a new elite card pack token", async () => {
//       const tokenURI = "https://example.com/token/1";

//       await expect(await cardPack.createEliteCard(tokenURI))
//         .to.emit(cardPack, "CardPackCreated")
//         .withArgs(0, 0, tokenURI, "Elite");

//       const tokenURIResult = await cardPack.tokenURI(0);
//       expect(tokenURIResult).to.equal(tokenURI);
//     });
//   });

//   describe("changeToOpened", () => {

//     it("should revert when a non-minter contract tries to change the status of a card pack", async () => {
//       await cardPack.createStandardCard("https://example.com/token/1");
//       await cardPack.connect(owner).changeMinter(await owner.getAddress());
//       await expect(
//         cardPack.connect(owner2).changeToOpened(0)
//       ).to.be.revertedWith("Card Pack: You are not the Minting Contract :)");
//     });

//     it("should revert when trying to change the status of a nonexistent card pack", async () => {
//       await expect(
//         cardPack.connect(owner).changeToOpened(1)
//       ).to.be.revertedWith("CardPack: Pack does not exist");
//     });
//   });

//   describe("isOpened", () => {

//     it("should return false if a card pack has not been opened", async () => {
//       await cardPack.createStandardCard("https://example.com/token/1");

//       const isOpenedResult = await cardPack.isOpened(0);
//       expect(isOpenedResult).to.equal(false);
//     });

//     it("should revert when checking the status of a nonexistent card pack", async () => {
//       await expect(cardPack.isOpened(0)).to.be.revertedWith(
//         "CardPack: Pack does not exist"
//       );
//     });
//   });

//   describe("tokenURI", () => {
//     it("should return the correct URI for a given token ID", async () => {
//       const tokenURI = "https://example.com/token/1";
//       await cardPack.createStandardCard(tokenURI);

//       const result = await cardPack.tokenURI(0);
//       expect(result).to.equal(tokenURI);
//     });

//     it("should revert when querying URI for a nonexistent token", async () => {
//       await expect(cardPack.tokenURI(0)).to.be.revertedWith(
//         "URI query for nonexistent token"
//       );
//     });
//   });

//   describe("changeMinter", () => {
//     it("should change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();

//       await cardPack.connect(owner).changeMinter(newMinter);

//       const minterAddress = await cardPack.minterContract();
//       expect(minterAddress).to.equal(newMinter);
//     });

//     it("should revert when a non-owner tries to change the minter contract address", async () => {
//       const newMinter = await owner2.getAddress();

//       await expect(
//         cardPack.connect(owner2).changeMinter(newMinter)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });

//     it("should emit a MinterContractChanged event when the minter contract address is changed", async () => {
//       const newMinter = await owner2.getAddress();
//       const oldOwner = await owner.getAddress();
//       await expect(cardPack.connect(owner).changeMinter(newMinter))
//         .to.emit(cardPack, "MinterContractChanged")
//         .withArgs(oldOwner, newMinter);
//     });

//     it("should revert when changing the minter contract address to the zero address", async () => {
//       const zeroAddress = ethers.constants.AddressZero;

//       await expect(
//         cardPack.connect(owner).changeMinter(zeroAddress)
//       ).to.be.revertedWith("CardPack: No zero address");
//     });
//   });

//   describe("changeToTotalLimit", () => {
//     it("should change the total limit of card packs", async () => {
//       const newLimit = 2000;

//       await cardPack.connect(owner).changeToTotalLimit(newLimit);

//       const totalLimit = await cardPack.totalLimit();
//       expect(totalLimit).to.equal(newLimit);
//     });

//     it("should revert when a non-owner tries to change the total limit", async () => {
//       const newLimit = 2000;

//       await expect(
//         cardPack.connect(owner2).changeToTotalLimit(newLimit)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });
//   });

//   describe("changeToStandardLimit", () => {
//     it("should change the standard card pack limit", async () => {
//       const newLimit = 150;

//       await cardPack.connect(owner).changeToStandardLimit(newLimit);

//       const standardLimit = await cardPack.standardLimit();
//       expect(standardLimit).to.equal(newLimit);
//     });

//     it("should revert when a non-owner tries to change the standard card pack limit", async () => {
//       const newLimit = 150;

//       await expect(
//         cardPack.connect(owner2).changeToStandardLimit(newLimit)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });
//   });

//   describe("changeToPremiumLimit", () => {
//     it("should change the premium card pack limit", async () => {
//       const newLimit = 75;

//       await cardPack.connect(owner).changeToPremiumLimit(newLimit);

//       const premiumLimit = await cardPack.premiumLimit();
//       expect(premiumLimit).to.equal(newLimit);
//     });

//     it("should revert when a non-owner tries to change the premium card pack limit", async () => {
//       const newLimit = 75;

//       await expect(
//         cardPack.connect(owner2).changeToPremiumLimit(newLimit)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });
//   });

//   describe("changeToEliteLimit", () => {
//     it("should change the elite card pack limit", async () => {
//       const newLimit = 30;

//       await cardPack.connect(owner).changeToEliteLimit(newLimit);

//       const eliteLimit = await cardPack.eliteLimit();
//       expect(eliteLimit).to.equal(newLimit);
//     });

//     it("should revert when a non-owner tries to change the elite card pack limit", async () => {
//       const newLimit = 30;

//       await expect(
//         cardPack.connect(owner2).changeToEliteLimit(newLimit)
//       ).to.be.revertedWith("Ownable: caller is not the owner");
//     });
//   });
// });

