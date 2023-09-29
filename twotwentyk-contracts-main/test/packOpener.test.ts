// import { ethers, upgrades } from "hardhat";
// import { Contract, Signer } from "ethers";
// import { expect } from "chai";

// import CategoryABI from '../artifacts/contracts/Parts/category.sol/Category.json'
// import YearABI from '../artifacts/contracts/Parts/year.sol/Year.json'
// import DayMonthABI from '../artifacts/contracts/Parts/dayMonth.sol/DayMonth.json'
// import craftingABI from '../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json'
// import TriggerABI from '../artifacts/contracts/Parts/triggers.sol/Trigger.json'

// const ifaceCat = new ethers.utils.Interface(["event CategoryDeployed(string name, string symbol, address category, address owner, uint)"]);
// const ifaceYr = new ethers.utils.Interface(["event YearDeployed(string name, string symbol, address year, address owner, uint)"]);
// const ifaceDM = new ethers.utils.Interface(["event DayMonthDeployed(string name, string symbol, address daymonth, address owner, uint)"]);
// const ifaceCrf = new ethers.utils.Interface(["event CraftingCardDeployed(string name, string symbol, address crafting, address owner, uint)"]);
// const ifaceTrg = new ethers.utils.Interface(["event TriggerDeployed(string name, string symbol, address trigger, address owner, uint)"]);

// describe("PackOpener", () => {
//   let packOpener: Contract;
//   let cardPack: Contract;
//   let factory: Contract;
//   let dayMonth: Contract;
//   let year: Contract;
//   let category: Contract;
//   let craftingCard: Contract;
//   let trigger: Contract;
//   let owner: Signer;
//   let attacker: Signer;

//   beforeEach(async () => {
//     [owner, attacker] = await ethers.getSigners();

//     const CardPack = await ethers.getContractFactory("CardPack");
//     cardPack = await upgrades.deployProxy(CardPack, [  
//         100, // standardAmount
//         50, // premiumAmount
//         20, // eliteAmount
//         1000 // totalLimit
//     ]);
//     await cardPack.deployed();

//     const Factory = await ethers.getContractFactory("CategoryFactory");
//     const _factory = await Factory.connect(owner).deploy();
//     factory = await _factory.deployed();

    
//     const txCat = await factory.deployCategory("Category", "CAT1", await owner.getAddress())
//     const receiptCat = await ethers.provider.getTransactionReceipt(txCat.hash);            
//     const eventCat = ifaceCat.parseLog(receiptCat.logs[2]);
//     category = new ethers.Contract(eventCat.args.category, CategoryABI.abi, owner)

//     const txYr = await factory.deployYear("Year", "YEAR1", await owner.getAddress())
//     const receiptYr = await ethers.provider.getTransactionReceipt(txYr.hash);        
//     const eventYr = ifaceYr.parseLog(receiptYr.logs[2]);
//     year = new ethers.Contract(eventYr.args.year, YearABI.abi, owner)

//     const txDM = await factory.deployDayMonth("DayMonth", "DM1", await owner.getAddress())
//     const receiptDM = await ethers.provider.getTransactionReceipt(txDM.hash);        
//     const eventDM = ifaceDM.parseLog(receiptDM.logs[2]);
//     dayMonth = new ethers.Contract(eventDM.args.daymonth, DayMonthABI.abi, owner)

//     const txCrf = await factory.deployDraftingCard("Crafting", "CRF1", await owner.getAddress())
//     const receiptCrf = await ethers.provider.getTransactionReceipt(txCrf.hash);        
//     const eventCrf = ifaceCrf.parseLog(receiptCrf.logs[2]);
//     craftingCard = new ethers.Contract(eventCrf.args.crafting, craftingABI.abi, owner)

//     const txTrg = await factory.deployTrigger("Year", "YEAR1", await owner.getAddress())
//     const receiptTrg = await ethers.provider.getTransactionReceipt(txTrg.hash);        
//     const eventTrg = ifaceTrg.parseLog(receiptTrg.logs[2]);
//     trigger = new ethers.Contract(eventTrg.args.trigger, TriggerABI.abi, owner)

//     const PackOpener = await ethers.getContractFactory("PackOpener");
//     packOpener = await upgrades.deployProxy(PackOpener, [cardPack.address])
//     await packOpener.deployed();   
    
    
//     await cardPack.connect(owner).changeMinter(packOpener.address);
//     await dayMonth.connect(owner).changeMinter(packOpener.address);
//     await year.connect(owner).changeMinter(packOpener.address);
//     await category.connect(owner).changeMinter(packOpener.address);
//     await craftingCard.connect(owner).changeMinter(packOpener.address);
//     await trigger.connect(owner).changeMinter(packOpener.address);
//     await cardPack
//       .connect(owner)
//       .createStandardCard("https://example.com/cardpack/1");
//   });

//   describe("openPack", () => {
//     it("should create new cards of each type when a pack is opened", async () => {
//       const packOwnerId = await cardPack.ownerOf(0);
//       const dayMonthAmountBefore = await dayMonth.balanceOf(packOwnerId);
//       const yearAmountBefore = await year.balanceOf(packOwnerId);
//       const categoryAmountBefore = await category.balanceOf(packOwnerId);
//       const craftingCardAmountBefore = await craftingCard.balanceOf(
//         packOwnerId
//       );
//       const triggerAmountBefore = await trigger.balanceOf(packOwnerId);

//       await packOpener.connect(owner).openPack(
//         0, // tokenId
//         dayMonth.address,
//         year.address,
//         category.address,
//         craftingCard.address,
//         trigger.address,        
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

//       const dayMonthAmountAfter = await dayMonth.balanceOf(packOwnerId);
//       const yearAmountAfter = await year.balanceOf(packOwnerId);
//       const categoryAmountAfter = await category.balanceOf(packOwnerId);
//       const craftingCardAmountAfter = await craftingCard.balanceOf(packOwnerId);
//       const triggerAmountAfter = await trigger.balanceOf(packOwnerId);

//       expect(dayMonthAmountAfter).to.equal(dayMonthAmountBefore.toNumber() + 2);
//       expect(yearAmountAfter).to.equal(yearAmountBefore.toNumber() + 3);
//       expect(categoryAmountAfter).to.equal(categoryAmountBefore.toNumber() + 4);
//       expect(craftingCardAmountAfter).to.equal(
//         craftingCardAmountBefore.toNumber() + 1
//       );
//       expect(triggerAmountAfter).to.equal(triggerAmountBefore.toNumber() + 1);
//     });

//     it("should only allow the pack owner to open a pack", async () => {
//     //   await expect(
//     //     packOpener.connect(attacker).openPack(
//     //       0, // tokenId
//     //       dayMonth.address,
//     //       year.address,
//     //       category.address,
//     //       craftingCard.address,
//     //       trigger.address,          
//     //       ["https://example.com/daymonth/1", "https://example.com/daymonth/2"],
//     //       [
//     //         "https://example.com/year/1",
//     //         "https://example.com/year/2",
//     //         "https://example.com/year/3",
//     //       ],
//     //       [
//     //         "https://example.com/category/1",
//     //         "https://example.com/category/2",
//     //         "https://example.com/category/3",
//     //         "https://example.com/category/4",
//     //       ],
//     //       ["https://example.com/craftingcard/1"],
//     //       ["https://example.com/trigger/1"]
//     //     )
//     //   ).to.be.revertedWith("You must be the admin wallet");
//     });

//     it("should only allow unopened packs to be opened", async () => {
//     //   await packOpener.connect(owner).openPack(
//     //     0, // tokenId
//     //     dayMonth.address,
//     //     year.address,
//     //     category.address,
//     //     craftingCard.address,
//     //     trigger.address,        
//     //     ["https://example.com/daymonth/1", "https://example.com/daymonth/2"],
//     //     [
//     //       "https://example.com/year/1",
//     //       "https://example.com/year/2",
//     //       "https://example.com/year/3",
//     //     ],
//     //     [
//     //       "https://example.com/category/1",
//     //       "https://example.com/category/2",
//     //       "https://example.com/category/3",
//     //       "https://example.com/category/4",
//     //     ],
//     //     ["https://example.com/craftingcard/1"],
//     //     ["https://example.com/trigger/1"]
//     //   );

//     //   await expect(
//     //     packOpener.connect(owner).openPack(
//     //       0, // tokenId
//     //       dayMonth.address,
//     //       year.address,
//     //       category.address,
//     //       craftingCard.address,
//     //       trigger.address,          
//     //       ["https://example.com/daymonth/1", "https://example.com/daymonth/2"],
//     //       [
//     //         "https://example.com/year/1",
//     //         "https://example.com/year/2",
//     //         "https://example.com/year/3",
//     //       ],
//     //       [
//     //         "https://example.com/category/1",
//     //         "https://example.com/category/2",
//     //         "https://example.com/category/3",
//     //         "https://example.com/category/4",
//     //       ],
//     //       ["https://example.com/craftingcard/1"],
//     //       ["https://example.com/trigger/1"]
//     //     )
//     //   ).to.be.revertedWith("CardPack: Pack does not exist"||"Pack already opened");
//     });
//   });
// });
