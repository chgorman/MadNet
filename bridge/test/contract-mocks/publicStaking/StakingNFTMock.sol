// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.11;

import "contracts/libraries/StakingNFT/StakingNFT.sol";

contract StakingNFTMock is StakingNFT {
    constructor() StakingNFT() {}

    function depositMock(
        uint256 shares_,
        uint256 delta_,
        Accumulator memory state_
    ) public pure returns (Accumulator memory) {
        return StakingNFT._deposit(shares_, delta_, state_);
    }

    // _slushSkim flushes value from the slush into the accumulator;
    // if there are no currently staked positions, all value is stored in the slush
    function slushSkimMock(
        uint256 shares_,
        uint256 accumulator_,
        uint256 slush_
    ) public pure returns (uint256, uint256) {
        return StakingNFT._slushSkim(shares_, accumulator_, slush_);
    }

    function epochRewardMock(
        uint32 epoch_,
        uint256 additionalNewTokens_,
        uint256 rewardEra_
    ) public pure returns (uint256) {
        return StakingNFT._epochReward(epoch_, additionalNewTokens_, rewardEra_);
    }

    function updateAccumulatorForMintingMock(
        uint32 epoch_,
        uint256 shares_,
        Accumulator memory state_,
        uint256 reserveToken_,
        uint256 additionalNewTokens_,
        uint256 rewardEra_
    ) public pure returns (Accumulator memory, uint256) {
        return
            StakingNFT._updateAccumulatorForMinting(
                epoch_,
                shares_,
                state_,
                reserveToken_,
                additionalNewTokens_,
                rewardEra_
            );
    }
}
