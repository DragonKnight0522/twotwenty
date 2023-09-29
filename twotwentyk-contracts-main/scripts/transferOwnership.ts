import { ethers } from "ethers";

import CardPack from "../artifacts/contracts/CardPack/cardPack.sol/CardPack.json";
import Category from "../artifacts/contracts/Parts/category.sol/Category.json";
import CraftingCard from "../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json";
import DayMonth from "../artifacts/contracts/Parts/dayMonth.sol/DayMonth.json";
import Year from "../artifacts/contracts/Parts/year.sol/Year.json";
import Trigger from "../artifacts/contracts/Parts/triggers.sol/Trigger.json";
import CollectionCrafting from "../artifacts/contracts/Crafting/craftingIdenyity.sol/CollectionCrafting.json";
import ProdictionCrafting from "../artifacts/contracts/Crafting/craftingPredictions.sol/PredictionCrafting.json";
import PackOper from '../artifacts/contracts/OpenPack/packOpener.sol/PackOpener.json'

import * as dotenv from "dotenv";

dotenv.config();

const main = async() => {
    const newOwner = '0xf22679BBaf587B9c776D0A25a05e64B214f19CAd';

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    

    const cardPackAddr = '0xf6D155F6D5F98A2A41595e2dd1D42989844B21B0';
    const categoryAddr = '0x9b036E2f59E2426C2901991E3d2A7901adc02b43';    
    const craftingCardAddr = '0x31397561EE61B6C533DbD03d8645ebcEd7fBF8B5';
    const craftingAddr = '0xeBbc8b54488e60bC141bc4cf4b22E310F297AFBB';
    const predictionAddrs = '0x8896A2BAc7490109572b9BCce834a8c590B1A67E';
    const packOpenerAddr = '0xe7871c5131B80216D676f0AE1fd7BAf00220d680';
    const dayMonthAddr = '0xBA105a9dcc4f508067AbE00eA8C38105688fd5ED';
    const yearAddr = '0x75f5628E11108EBF7E3Df47F42202d6CeEEfB99A';
    const triggerAddr = '0x58CD76087F1CA00503Bba64103E00424338CcA03';    

    const cardPack = new ethers.Contract(cardPackAddr, CardPack.abi, wallet);

    const year = new ethers.Contract(yearAddr, Year.abi, wallet);
    const category = new ethers.Contract(categoryAddr, Category.abi, wallet);
    const dayMonth = new ethers.Contract(dayMonthAddr, DayMonth.abi, wallet);
    const trigger = new ethers.Contract(triggerAddr, Trigger.abi, wallet)

    const crafting = new ethers.Contract(craftingAddr, CollectionCrafting.abi, wallet);    
    const prediction = new ethers.Contract(predictionAddrs, ProdictionCrafting.abi, wallet);

    const craftingCard = new ethers.Contract(craftingCardAddr, CraftingCard.abi, wallet);
    const packOpener = new ethers.Contract(packOpenerAddr, PackOper.abi, wallet)


    await cardPack.transferOwnership(newOwner);
    await category.transferOwnership(newOwner);
    await dayMonth.transferOwnership(newOwner);
    await year.transferOwnership(newOwner);
    await trigger.transferOwnership(newOwner);
    await crafting.transferOwnership(newOwner);
    await prediction.transferOwnership(newOwner);
    await craftingCard.transferOwnership(newOwner);
    await packOpener.transferOwnership(newOwner);
}

main().then()