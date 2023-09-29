import { ethers } from "ethers";
import * as dotenv from "dotenv";

dotenv.config();

import CardPack from "../artifacts/contracts/CardPack/cardPack.sol/CardPack.json";
import Category from "../artifacts/contracts/Parts/category.sol/Category.json";
import CraftingCard from "../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json";
import DayMonth from "../artifacts/contracts/Parts/dayMonth.sol/DayMonth.json";
import Year from "../artifacts/contracts/Parts/year.sol/Year.json";
import Trigger from "../artifacts/contracts/Parts/triggers.sol/Trigger.json";
import CollectionCrafting from "../artifacts/contracts/Crafting/craftingIdenyity.sol/CollectionCrafting.json";
import PredictionCrafting from "../artifacts/contracts/Crafting/craftingPredictions.sol/PredictionCrafting.json";
import PackOpern from '../artifacts/contracts/OpenPack/packOpener.sol/PackOpener.json'

const main = async () => {
    const cardPackAddr = '0xf6D155F6D5F98A2A41595e2dd1D42989844B21B0';
    const categoryAddr = '0x9b036E2f59E2426C2901991E3d2A7901adc02b43';    
    const craftingCardAddr = '0x31397561EE61B6C533DbD03d8645ebcEd7fBF8B5';
    const craftingAddr = '0xeBbc8b54488e60bC141bc4cf4b22E310F297AFBB';
    const predictionAddrs = '0x8896A2BAc7490109572b9BCce834a8c590B1A67E';
    
    const dayMonthAddr = '0xBA105a9dcc4f508067AbE00eA8C38105688fd5ED';
    const yearAddr = '0x75f5628E11108EBF7E3Df47F42202d6CeEEfB99A';
    const triggerAddr = '0x58CD76087F1CA00503Bba64103E00424338CcA03';
    const PackOpernAddr = '0xe7871c5131B80216D676f0AE1fd7BAf00220d680';

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    

    const crafting = new ethers.Contract(craftingAddr, CollectionCrafting.abi, wallet);
    const prediction = new ethers.Contract(predictionAddrs, PredictionCrafting.abi, wallet);

    const packOpener = new ethers.Contract(PackOpernAddr, PackOpern.abi, wallet);

    const tokenURI = "https://example.com/token/1";

    const cardPack = new ethers.Contract(cardPackAddr, CardPack.abi, wallet);
    // Mint cardpack    
    // await cardPack.createStandardCard(tokenURI);

    //Mint category token
    await packOpener.openPack(
        2, // tokenId
        2, // dayMonthAmount
        3, // yearAmount
        4, // categoryAmount
        2, // craftingCardAmount
        1, // triggerAmount
        ["https://example.com/daymonth/1", "https://example.com/daymonth/2"],
        [
          "https://example.com/year/1",
          "https://example.com/year/2",
          "https://example.com/year/3",
        ],
        [
          "https://example.com/category/1",
          "https://example.com/category/2",
          "https://example.com/category/3",
          "https://example.com/category/4",
        ],
        ["https://example.com/craftingcard/1", "https://example.com/craftingcard/2"],
        ["https://example.com/trigger/1"]
      );
    
    // crafting identification
    // await crafting.craftCollection(0, 0, 0, 0, "https://example.com/identity/1");

    // // crafting prediction
    // await prediction.craftCollection(0, [0], 1, "https://example.com/prediction/1")
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });