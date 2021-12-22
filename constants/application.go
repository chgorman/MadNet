package constants

const (
	// MaxTxVectorLength is the maximum size of input output vectors.
	// This prevents uint32 overflow.
	MaxTxVectorLength int = 128

	// MaxTxSize is the maximum size in bytes of a marshaled transaction;
	// 2.125 MiB. This is more than 100kB larger than the largest allowable
	// DataStore, so this should be more than sufficient.
	MaxTxSize int = 2228224
)

const (
	// MaxFeeSizeRatio is 2^123 as a float32.
	// Any transaction with a FeeSizeRatio above this value will
	// be rounded down to this value.
	MaxFeeSizeRatio float32 = 10633823966279326983230456482242756608
)

const (
	// DSPIMinDeposit is the minimum amount of deposit. This is calculated
	// assuming that no data is stored (datasize == 0) as well as storing
	// the data for 1 epoch.
	DSPIMinDeposit uint32 = BaseDatasizeConst

	// BaseDatasizeConst is the bytes added to the size of data (in bytes)
	// for the minimum cost.
	BaseDatasizeConst = 376
	// We discuss the rational now.
	//
	// At this point there are two possibilities for accounts,
	// Secp256k1 and BN256Eth. We have the following
	//
	//	Curve 			Size of RawData 	len(DataStore.MarshalBinary())
	//
	//	Secp256k1				32						272
	//	         				64						304
	//	         				96						336
	//
	//	Bn256Eth 				32						400
	//	         				64						432
	//	         				96						464
	//
	// From looking at this table, the base cost for Secp256k1 is 240 bytes
	// while the base cost for BN256Eth is 368 bytes. Many machines have
	// word size of 64 bits, so we add an additional 8 bytes to the baseline
	// to ensure all storage is needed to ensure proper alignment.
	// Thus, it makes sense to set 376 bytes as the base cost of DataStore.

	// MaxDataStoreSize is the largest size of RawData that we store in a
	// DataStore; 2 MiB (2^21)
	//
	MaxDataStoreSize uint32 = 2097152
	//
	// Do not change this value without ensuring BaseDepositEquation will
	// not overflow in uint64. This will not happen so long as
	//
	// 		MaxDataStoreSize + BaseDatasizeConst < 2^31
	//
	// This restriction should not cause problems.
)
