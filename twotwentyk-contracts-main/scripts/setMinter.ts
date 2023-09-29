import { ethers } from "ethers";
import * as dotenv from "dotenv";

dotenv.config();

import CollectionABI from "../artifacts/contracts/Collection/collection.sol/Collection.json";
import CardPackFactoryABI from "../artifacts/contracts/Factories/cardpackFactory.sol/CardPackFactory.json";
import CategoryFactoryABI from "../artifacts/contracts/Factories/categoryFactory.sol/CategoryFactory.json";
import CraftingCardFactoryABI from "../artifacts/contracts/Factories/CraftingCardFactory.sol/CraftingCardFactory.json";
import DayMonthFactoryABI from "../artifacts/contracts/Factories/daymonthFactory.sol/DayMonthFactory.json";
import YearFactoryABI from "../artifacts/contracts/Factories/yearFactory.sol/YearFactory.json";
import TriggerFactoryABI from "../artifacts/contracts/Factories/yearFactory.sol/YearFactory.json";

const main = async () => {
    
    const collectionAddr = "0x2AC89A2eAF323558bac2358b2B508DaD2891D330";
    const cardPackFactoryAddr = '0x75E9869181459ee29234E78a1100Ee213A1ADBca';
    const categoryFactoryAddr = '0xf474b7B4d946e559f156c7d5d02e8f5Ab994bDB6';    
    const yearFactoryAddr = '0x687Dcd759C323dBB76D53a9D5893F9D391b93358';
    const dayMonthFactoryAddr = '0xb040c97359377478D51Bf0b3C5D459FC46250E43';
    const craftingCardFactoryAddr = '0x3c4Ca73bFA4C4CCa9D67b7D79b23dBdc2703Fe5C';
    const triggerFactoryAddr = '0x78BD6D288237b0CEC3fF4Ee95Aa76A70F660A174';    

    const minterAddr = "0xf22679BBaf587B9c776D0A25a05e64B214f19CAd"

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    
    
    const cardPackFactory = new ethers.Contract(cardPackFactoryAddr, CardPackFactoryABI.abi, wallet);
    await cardPackFactory.changeAdmin(collectionAddr);

    const categoryFactory = new ethers.Contract(categoryFactoryAddr, CategoryFactoryABI.abi, wallet);
    await categoryFactory.changeAdmin(collectionAddr);
    
    const yearFactory = new ethers.Contract(yearFactoryAddr, YearFactoryABI.abi, wallet);
    await yearFactory.changeAdmin(collectionAddr);
    
    const daymonthFactory = new ethers.Contract(dayMonthFactoryAddr, DayMonthFactoryABI.abi, wallet);
    await daymonthFactory.changeAdmin(collectionAddr);

    const craftingCardFactory = new ethers.Contract(craftingCardFactoryAddr, CraftingCardFactoryABI.abi, wallet);
    await craftingCardFactory.changeAdmin(collectionAddr);

    const triggerFactory = new ethers.Contract(triggerFactoryAddr, TriggerFactoryABI.abi, wallet);
    await triggerFactory.changeAdmin(collectionAddr);

    const collection = new ethers.Contract(collectionAddr, CollectionABI.abi, wallet);
    await collection.changeMinter(minterAddr);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });