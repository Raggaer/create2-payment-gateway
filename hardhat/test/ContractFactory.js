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

  function deriveDrainAddress(factory, contract, owner) {
    const params = utils.defaultAbiCoder.encode(["address payable"], [owner]);
    const bc = factory.bytecode + params.slice(2); // Remove the 0x from the params
    const salt = utils.formatBytes32String("payment-test");

    return [bc, salt, deriveAddress(contract.address, bc, salt)];
  }

  // Deploys the Drain contract on the calculated address
  async function deployDrainContract(factory, contract, owner) {
    const [bc, salt, expectedAddress] = deriveDrainAddress(
      factory,
      contract,
      owner
    );

    const tx = await contract.deploy(bc, salt);
    const receipt = await tx.wait();

    const event = receipt.events.find((e) => e.event === "Deployed");

    expect(event).to.not.be.undefined;
    expect(event.args[0]).to.equal(expectedAddress);

    return expectedAddress;
  }

  // Deploys the Sample testing contract on the calculated address
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

  it("Should be able to deploy the Drain contract", async function () {
    const [owner, addr1, addr2] = await ethers.getSigners();

    const drainFactory = await ethers.getContractFactory("Drain");
    const factory = await ethers.getContractFactory("ContractFactory");
    const contract = await factory.deploy(owner.address);

    const addr = await deployDrainContract(
      drainFactory,
      contract,
      addr2.address // Where the funds will be sent to
    );

    const drain = await ethers.getContractAt("Drain", addr);
  });

  it("Should be able to drain correctly after being deployed", async function () {
    const [owner, addr1, addr2] = await ethers.getSigners();

    const drainFactory = await ethers.getContractFactory("Drain");
    const factory = await ethers.getContractFactory("ContractFactory");
    const contract = await factory.deploy(owner.address);

    // Derive the address of the contract first so we can send funds to it
    [bc, salt, expectedAddress] = deriveDrainAddress(
      drainFactory,
      contract,
      addr2.address
    );

    // Send ETH to the derived address
    await addr1.sendTransaction({
      to: expectedAddress,
      value: ethers.utils.parseEther("22"),
    });

    const addr = await deployDrainContract(
      drainFactory,
      contract,
      addr2.address // Where the funds will be sent to
    );

    expect(addr).to.equal(expectedAddress);

    const drain = await ethers.getContractAt("Drain", addr);

    // Now the funds are drained by the constructor and sent to addr2
    expect(await ethers.provider.getBalance(addr2.address)).to.equal(
      ethers.utils.parseEther("10022")
    );

    // Send ETH to the derived address
    await addr1.sendTransaction({
      to: expectedAddress,
      value: ethers.utils.parseEther("22"),
    });
    expect(await ethers.provider.getBalance(addr2.address)).to.equal(
      ethers.utils.parseEther("10044")
    );
  });
});
