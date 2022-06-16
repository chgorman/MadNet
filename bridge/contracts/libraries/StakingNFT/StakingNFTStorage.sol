// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.11;

abstract contract StakingNFTStorage {
    // Position describes a staked position
    struct Position {
        // number of aTokens when including weight
        uint224 weightedShares;
        // lockedStakingPosition specifies if the staked position is locked;
        // this is only used for locked staking positions.
        bool lockedStakingPosition;
        // number of aToken
        uint224 shares;
        // block number after which the position may be burned.
        // prevents double spend of voting weight
        uint32 freeAfter;
        // block number after which the position may be collected or burned.
        uint256 withdrawFreeAfter;
        // the last value of the ethState accumulator this account performed a
        // withdraw at
        uint256 accumulatorEth;
        // the last value of the tokenState accumulator this account performed a
        // withdraw at
        uint256 accumulatorToken;
    }

    // Accumulator is a struct that allows values to be collected such that the
    // remainders of floor division may be cleaned up
    struct Accumulator {
        // accumulator is a sum of all changes always increasing
        uint256 accumulator;
        // slush stores division remainders until they may be distributed evenly
        uint256 slush;
    }

    // _MAX_MINT_LOCK describes the maximum interval a Position may be locked
    // during a call to mintTo
    uint256 internal constant _MAX_MINT_LOCK = 1051200;
    // 10**18
    uint256 internal constant _ACCUMULATOR_SCALE_FACTOR = 1000000000000000000;
    // constants for the cb state
    bool internal constant _CIRCUIT_BREAKER_OPENED = true;
    bool internal constant _CIRCUIT_BREAKER_CLOSED = false;

    // monotonically increasing counter
    uint256 internal _counter;

    // cb is the circuit breaker
    // cb is a set only object
    bool internal _circuitBreaker;

    // _shares stores total amount of AToken staked in contract
    uint256 internal _sharesEth;

    // _shares stores total amount of AToken staked in contract
    uint256 internal _sharesToken;

    // _tokenState tracks distribution of AToken that originate from slashing
    // events
    Accumulator internal _tokenState;

    // _ethState tracks the distribution of Eth that originate from the sale of
    // BTokens
    Accumulator internal _ethState;

    // _positions tracks all staked positions based on tokenID
    mapping(uint256 => Position) internal _positions;

    // state to keep track of the amount of Eth deposited and collected from the
    // contract
    uint256 internal _reserveEth;

    // state to keep track of the amount of ATokens deposited and collected
    // from the contract
    uint256 internal _reserveToken;

    /*
    // denominator used when computing weighted stake
    uint256 internal constant _LOCKING_TIER_DENOMINATOR = 1000000;

    // Tier 0; required locking for one block; no AToken rewards
    // uint256 internal constant _LOCKING_TIER_0 = 1;
    // uint256 internal constant _LOCKING_TIER_0_NUMERATOR = 0;

    // 1 week;     84 epochs
    uint256 internal constant _LOCKING_TIER_1 = (_MAX_MINT_LOCK * 70) / 1825;
    uint256 internal constant _LOCKING_TIER_1_NUMERATOR = 1000001;

    // 1 month;   367 epochs
    uint256 internal constant _LOCKING_TIER_2 = _MAX_MINT_LOCK / 6;
    uint256 internal constant _LOCKING_TIER_2_NUMERATOR = 1010000;

    // 3 months; 1100 epochs
    uint256 internal constant _LOCKING_TIER_3 = _MAX_MINT_LOCK / 2;
    uint256 internal constant _LOCKING_TIER_3_NUMERATOR = 1100000;

    // 6 months; 2200 epochs
    uint256 internal constant _LOCKING_TIER_4 = _MAX_MINT_LOCK;
    uint256 internal constant _LOCKING_TIER_4_NUMERATOR = 2000000;
    */

    // _REWARD_ERA specifies the number of epochs per reward era,
    // which determines the specific decay rate of additional ATokens which
    // are minted each snapshot (epoch).
    uint32 internal constant _REWARD_ERA = 2200;

    // _ADDITIONAL_ATOKENS specifies the total number of additional ATokens
    // which will be minted.
    //
    // TODO: this value must be determined
    uint256 internal constant _ADDITIONAL_ATOKENS = 220000000000000000000000000;
}
