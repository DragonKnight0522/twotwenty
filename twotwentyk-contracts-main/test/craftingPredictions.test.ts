// import { ethers } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// describe("CraftingPrediction", () => {
//   let packOpener: Contract;
//   let cardPack: Contract;
//   let dayMonth: Contract;
//   let year: Contract;
//   let category: Contract;
//   let craftingCard: Contract;
//   let crafting: Contract;
//   let owner: Signer;
//   let packOwner: Signer;
//   let trigger: Contract;
//   let predictor: Contract;

//   beforeEach(async () => {
//     [owner, packOwner] = await ethers.getSigners();

//     const CardPack = await ethers.getContractFactory("CardPack");
//     cardPack = await CardPack.connect(owner).deploy(
//       100, // standardAmount
//       50, // premiumAmount
//       20, // eliteAmount
//       1000 // totalLimit
//     );
//     await cardPack.deployed();

//     const DayMonth = await ethers.getContractFactory("DayMonth");
//     dayMonth = await DayMonth.connect(owner).deploy();
//     await dayMonth.deployed();

//     const Year = await ethers.getContractFactory("Year");
//     year = await Year.connect(owner).deploy();
//     await year.deployed();

//     const Category = await ethers.getContractFactory("Category");
//     category = await Category.connect(owner).deploy();
//     await category.deployed();

//     const CraftingCard = await ethers.getContractFactory("CraftingCard");
//     craftingCard = await CraftingCard.connect(owner).deploy();
//     await craftingCard.deployed();

//     const Trigger = await ethers.getContractFactory("Trigger");
//     trigger = await Trigger.connect(owner).deploy();
//     await trigger.deployed();

//     const PackOpener = await ethers.getContractFactory("PackOpener");
//     packOpener = await PackOpener.connect(owner).deploy(
//       cardPack.address, // cardPackContract
//       dayMonth.address, // dayMonthContract
//       year.address, // yearContract
//       category.address, // categoryContract
//       craftingCard.address, // craftingCardContract
//       trigger.address // triggerContract
//     );
//     await packOpener.deployed();

//     const CraftingIdentity = await ethers.getContractFactory(
//       "CollectionCrafting"
//     );
//     crafting = await CraftingIdentity.connect(owner).deploy(
//       dayMonth.address, // dayMonthContract
//       year.address, // yearContract
//       category.address, // categoryContract
//       craftingCard.address // craftingCardContract
//     );
//     await crafting.deployed();
//     const Predictor = await ethers.getContractFactory("PredictionCrafting");
//     predictor = await Predictor.connect(owner).deploy(
//       crafting.address, // dayMonthContract
//       trigger.address, // yearContract
//       craftingCard.address // craftingCardContract
//     );
//     await predictor.deployed();

//     await cardPack.connect(owner).changeMinter(packOpener.address);
//     await dayMonth.connect(owner).changeMinter(packOpener.address);
//     await year.connect(owner).changeMinter(packOpener.address);
//     await category.connect(owner).changeMinter(packOpener.address);
//     await craftingCard.connect(owner).changeMinter(packOpener.address);
//     await trigger.connect(owner).changeMinter(packOpener.address);
//     await trigger.connect(owner).changeCrafter(predictor.address);
//     await dayMonth.connect(owner).changeCrafter(crafting.address);
//     await year.connect(owner).changeCrafter(crafting.address);
//     await category.connect(owner).changeCrafter(crafting.address);
//     await craftingCard.connect(owner).changeCrafter(crafting.address);
//     await craftingCard.connect(owner).changePredictor(predictor.address);
//     await crafting.connect(owner).changeCrafter(predictor.address);
//     for (let i = 0; i < 100; i++) {
//       await cardPack
//         .connect(owner)
//         .createStandardCard("https://example.com/cardpack/1");
//       await packOpener.connect(owner).openPack(
//         i, // tokenId
//         2, // dayMonthAmount
//         3, // yearAmount
//         4, // categoryAmount
//         1, // craftingCardAmount
//         1, // triggerAmount
//         ["https://example.com/daymonth/1", "https://example.com/daymonth/2"],
//         [
//           "https://example.com/year/1",
//           "https://example.com/year/2",
//           "https://example.com/year/3",
//         ],
//         [
//           "https://example.com/category/1",
//           "https://example.com/category/2",
//           "https://example.com/category/3",
//           "https://example.com/category/4",
//         ],
//         ["https://example.com/craftingcard/1"],
//         ["https://example.com/trigger/1"]
//       );
//     }
//   });

//   describe("craftCollection", () => {
//     it("should allow a user with the right cards to craft an identity", async () => {
//       for (let i = 0; i < 10; i++) {
//         await crafting
//           .connect(owner)
//           .craftCollection(i, i, i, i, "https://example.com/cardpack/1");
//       }
//       const identityOwner = await crafting.ownerOf(0);
//       for (let i = 0; i < 10; i++) {
//         await predictor
//           .connect(owner)
//           .craftCollection(i, [i], i + 10, "https://example.com/cardpack/1");
//       }
//       expect(identityOwner).to.equal(await owner.getAddress());
//     });

//     it("should not allow a user without the right cards to craft an identity", async () => {
//       for (let i = 0; i < 10; i++) {
//         await crafting
//           .connect(owner)
//           .craftCollection(i, i, i, i, "https://example.com/cardpack/1");
//       }
//       const identityOwner = await crafting.ownerOf(0);
//       for (let i = 0; i < 10; i++) {
//         await predictor
//           .connect(owner)
//           .craftCollection(i, [i], i + 10, "https://example.com/cardpack/1");
//       }
//       expect(identityOwner).to.equal(await owner.getAddress());

//       await expect(
//         predictor.connect(owner).craftCollection(111, [111], 11, "")
//       ).to.be.revertedWith("ERC721: invalid token ID" || "Not enough cards");
//     });

//     it("should not allow a user to craft an identity with the same crafting card", async () => {
//             for (let i = 0; i < 10; i++) {
//               await crafting
//                 .connect(owner)
//                 .craftCollection(i, i, i, i, "https://example.com/cardpack/1");
//             }
//             const identityOwner = await crafting.ownerOf(0);
//             for (let i = 0; i < 10; i++) {
//               await predictor
//                 .connect(owner)
//                 .craftCollection(
//                   i,
//                   [i],
//                   i + 10,
//                   "https://example.com/cardpack/1"
//                 );
//             }
//             expect(identityOwner).to.equal(await owner.getAddress());

//       await expect(
//         predictor.connect(owner).craftCollection(1, [1], 1, "")
//       ).to.be.revertedWith(
//         "ERC721: invalid token ID" ||
//           "Day/Month already used" ||
//           "Duplicate cards not allowed"
//       );
//     });

//     it("should burn the cards used to craft an identity", async () => {
//       for (let i = 0; i < 10; i++) {
//         await crafting
//           .connect(owner)
//           .craftCollection(i, i, i, i, "https://example.com/cardpack/1");
//       }

//       await expect(craftingCard.ownerOf(1)).to.be.reverted;
//     });
//   });
// });
