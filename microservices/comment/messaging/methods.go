package messaging

func calculateBatchSize(count uint32, batchSize uint32) uint32 {
	if count == 0 {
		return 0
	}
	if remainder := count % batchSize; remainder != 0 {
		return remainder
	}
	return batchSize
}
