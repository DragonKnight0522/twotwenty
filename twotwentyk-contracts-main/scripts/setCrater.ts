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

const main = async () => {
    
    const cardPackAddr = '0xf6D155F6D5F98A2A41595e2dd1D42989844B21B0';
    const categoryAddr = '0x9b036E2f59E2426C2901991E3d2A7901adc02b43';    
    const craftingCardAddr = '0x31397561EE61B6C533DbD03d8645ebcEd7fBF8B5';
    const craftingAddr = '0xeBbc8b54488e60bC141bc4cf4b22E310F297AFBB';
    const predictionAddrs = '0x8896A2BAc7490109572b9BCce834a8c590B1A67E';
    const packOpenerAddr = '0xe7871c5131B80216D676f0AE1fd7BAf00220d680';
    const dayMonthAddr = '0xBA105a9dcc4f508067AbE00eA8C38105688fd5ED';
    const yearAddr = '0x75f5628E11108EBF7E3Df47F42202d6CeEEfB99A';
    const triggerAddr = '0x58CD76087F1CA00503Bba64103E00424338CcA03';    

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    
    
    const cardPack = new ethers.Contract(cardPackAddr, CardPack.abi, wallet);
    await cardPack.changeMinter(packOpenerAddr);

    // Set PackOpener address as the minter and crafter for Category
    
    const category = new ethers.Contract(categoryAddr, Category.abi, wallet);
    await category.changeMinter(packOpenerAddr);
    await category.changeCrafter(craftingAddr);
  
    // Set PackOpener address as the minter and crafter for CraftingCard

    const craftingCard = new ethers.Contract(craftingCardAddr, CraftingCard.abi, wallet);
    await craftingCard.changeMinter(packOpenerAddr);
    await craftingCard.changeCrafter(craftingAddr);
    await craftingCard.changePredictor(predictionAddrs);
  
    // Set PackOpener address as the minter and crafter for DayMonth

    const dayMonth = new ethers.Contract(dayMonthAddr, DayMonth.abi, wallet);
    await dayMonth.changeMinter(packOpenerAddr);
    await dayMonth.changeCrafter(craftingAddr);
  
    // Set PackOpener address as the minter and crafter for Year

    const year = new ethers.Contract(yearAddr, Year.abi, wallet);
    await year.changeMinter(packOpenerAddr);
    await year.changeCrafter(craftingAddr);
  
    const trigger = new ethers.Contract(triggerAddr, Trigger.abi, wallet)
    await trigger.changeMinter(packOpenerAddr);
    await trigger.changeCrafter(predictionAddrs);
  
    const crafting = new ethers.Contract(craftingAddr, CollectionCrafting.abi, wallet)
    await crafting.changeCrafter(predictionAddrs);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });