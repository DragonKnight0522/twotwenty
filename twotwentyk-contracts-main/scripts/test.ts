import { ethers } from "ethers";
import * as dotenv from "dotenv";
import CollectionABI from "../artifacts/contracts/Collection/collection.sol/Collection.json";
import packABI from "../artifacts/contracts/CardPack/cardPack.sol/CardPack.json"
import CategoryABI from "../artifacts/contracts/Parts/category.sol/Category.json"
import YearABI from "../artifacts/contracts/Parts/year.sol/Year.json";
import DayMonthABI from "../artifacts/contracts/Parts/dayMonth.sol/DayMonth.json"
import CraftingCardABI from "../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json";
import TriggerABI from "../artifacts/contracts/Parts/triggers.sol/Trigger.json";
import packOpenerABI from "../artifacts/contracts/OpenPack/packOpener.sol/PackOpener.json"

dotenv.config()

const main = async() => {
    const collectionAddr = '0x2AC89A2eAF323558bac2358b2B508DaD2891D330';    
    const packAddr = "0x8BAf45B2D89ff1A15AE854C612b6EB9F56941929"
    const categoryAddr = "0x9E7cC11958A4578Bc193Cf2D36743A3a90A2D925"
    const yearAddr = "0x0E714bcAC55aD0270d34B092e20f078059e05A35"
    const dayMAddr = "0xCcd8e99d3DdBDC5f83FC551e79A2CE01b022baE3"
    const craftingCAddr = "0x8e2bda190094ba5ee22cb9be48a052fedad266c3"
    const triggerAddr = "0x1255ed868f6310e91b278c7a419364a25b64ca69"
    const packOpenerAddr = "0x5D6A8216590cEEeFd149f3CB26ad096B8516bB00"

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    

    const collection = new ethers.Contract(collectionAddr, CollectionABI.abi, wallet);    

    await collection.createCollection("https://example.com/collection/0", 300, 200, 100, 600)

    const pack = new ethers.Contract(packAddr, packABI.abi, wallet)
    const cateogry = new ethers.Contract(categoryAddr, CategoryABI.abi, provider)
    const year = new ethers.Contract(yearAddr, YearABI.abi, provider)
    const daymonth = new ethers.Contract(dayMAddr, DayMonthABI.abi, provider)
    const craftingCard = new ethers.Contract(craftingCAddr, CraftingCardABI.abi, provider)
    const triggerCard = new ethers.Contract(triggerAddr, TriggerABI.abi, provider)
    const packOpener = new ethers.Contract(packOpenerAddr, packOpenerABI.abi, wallet)

    console.log('minter category', await cateogry.name())
    console.log('minter year', await year.name())
    console.log('minter daymonth', await daymonth.name())
    console.log('minter crafting card', await craftingCard.name())
    console.log('minter trigger', await triggerCard.name())
    console.log('token owner', await pack.isOpened(1))
    console.log('token owner', await pack.ownerOf(0))
    // console.log('token owner', await packOpener.adminWallet())
    // await packOpener.openPack(
    //     1,
    //     packAddr,
    //     categoryAddr,
    //     yearAddr,
    //     dayMAddr,
    //     triggerAddr,
    //     craftingCAddr,
        
    //     ["https://example.com/category/1"],
    //     ["https://example.com/category/1"],
    //     ["https://example.com/category/1"],
    //     ["https://example.com/category/1", "https://example.com/category/1"],
    //     ["https://example.com/category/1"],
    // )
       
}

main().then()