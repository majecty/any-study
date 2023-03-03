import { ethers } from "ethers";

async function main() {
    // 10000000000000000000000000000000000n
    // const a = ethers.parseUnits("10000000000000000", 18);
    // console.log(a);

    const b = ethers.formatEther("100000000000000000");
    console.log(b);
    const c = ethers.parseEther("0.01");
    console.log(c);

}

main().catch(console.error);