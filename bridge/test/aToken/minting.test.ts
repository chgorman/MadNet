import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { expect, factoryCallAnyFixture, Fixture, getFixture } from "../setup";
import { getState, init, state } from "./setup";

describe("Testing AToken", async () => {
  let user: SignerWithAddress;
  let admin: SignerWithAddress;
  let expectedState: state;
  let currentState: state;
  const amount = 1000;
  let fixture: Fixture;

  beforeEach(async function () {
    fixture = await getFixture();
    [admin, user] = await ethers.getSigners();
    await init(fixture);
    expectedState = await getState(fixture);
  });

  describe("Testing minting operation", async () => {
    describe("Methods with onlyFactory modifier", async () => {
      describe("Methods with onlyATokenMinter modifier", async () => {
        it("Should not mint when called by external address not identified as minter", async function () {
          await expect(
            fixture.aToken.externalMint(user.address, amount)
          ).to.be.revertedWith("2013");

          await expect(
            fixture.aToken.connect(admin).externalMint(user.address, amount)
          ).to.be.revertedWith("2013");
        });
      });

      describe("Business methods with onlyFactory modifier", async () => {
        it("Should mint when called by external identified as minter impersonating factory", async function () {
          await factoryCallAnyFixture(fixture, "aTokenMinter", "mint", [
            user.address,
            amount,
          ]);
          expectedState.Balances.aToken.user += amount;
          currentState = await getState(fixture);
          expect(currentState).to.be.deep.eq(expectedState);
        });

        it("Should not mint when called by external identified as minter not impersonating factory", async function () {
          await expect(
            fixture.aTokenMinter.mint(user.address, amount)
          ).to.be.rejectedWith("2000");
        });
      });
    });
  });

  describe("Testing staking operations", async () => {
    describe("Test token accumulator functions", async () => {
      it("Ensure accumulatorEth is 0", async function () {
        // currentState = await getState(fixture);
        let stakingNFTIFace = await ethers.getContractFactory("StakingNFTMock");
        let stakingNFT = await stakingNFTIFace.deploy();
        let epoch = 1;
        let additionalTokens = 1000;
        let rewardEra = 1000;
        let txResponse = await stakingNFT.epochRewardMock(
          epoch,
          additionalTokens,
          rewardEra
        );
        let value = await txResponse.toBigInt();
        await expect(value).to.be.equal(1000);
        // expect(txResponse).to.be.revertedWith // assert with certain error codes
      });
    });
  });
});
