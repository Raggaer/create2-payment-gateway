const hre = require("hardhat");

async function main() {
  const [deployer, addr1, addr2] = await hre.ethers.getSigners();
  const gasPrice = await hre.ethers.provider.getGasPrice();

  const tx = await addr2.sendTransaction({
    to: "0x9828fc7bE2180eAc78E8969bf3367478F16BA153",
    gasLimit: 210000,
    gasPrice: gasPrice,
    value: hre.ethers.utils.parseEther("20"),
  });

  const receipt = await tx.wait();
  console.log(receipt);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
