const { expect } = require("chai");
const { ethers } = require("hardhat");

const utils = ethers.utils;

describe("ContractFactory", function () {
  function deriveAddress(addr, bytecode, salt) {
    return ethers.utils.getCreate2Address(
      addr,
      salt,
      utils.keccak256(bytecode)
    );
  }

  async function deployContract(factory, contract) {
    const params = utils.defaultAbiCoder.encode(["uint256"], [20]);
    const bc = factory.bytecode + params.slice(2);
    const salt = utils.formatBytes32String("Hello World");

    const tx = await contract.deploy(bc, salt);
    const receipt = await tx.wait();

    const event = receipt.events.find((e) => e.event === "Deployed");
    const expectedAddress = deriveAddress(contract.address, bc, salt);

    expect(event).to.not.be.undefined;
    expect(event.args[0]).to.equal(expectedAddress);

    return expectedAddress;
  }

  it("Should deploy", async function () {
    const [owner] = await ethers.getSigners();

    const factory = await ethers.getContractFactory("ContractFactory");
    const contract = await factory.deploy(owner.address);
  });

  it("Should create a contract", async function () {
    const [owner] = await ethers.getSigners();

    const sampleFactory = await ethers.getContractFactory("Sample");
    const factory = await ethers.getContractFactory("ContractFactory");
    const contract = await factory.deploy(owner.address);

    await deployContract(sampleFactory, contract);
  });

  it("Should be able to interact with the deployed contract", async function () {
    const [owner] = await ethers.getSigners();

    const sampleFactory = await ethers.getContractFactory("Sample");
    const factory = await ethers.getContractFactory("ContractFactory");
    const contract = await factory.deploy(owner.address);

    const addr = await deployContract(sampleFactory, contract);

    const sample = await ethers.getContractAt("Sample", addr);
    expect(await sample.get()).to.equal(20);

    await sample.set(100);
    expect(await sample.get()).to.equal(100);
  });
});
