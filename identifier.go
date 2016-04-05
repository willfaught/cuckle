package cuckle

import "fmt"

type Identifier string

const (
	FuncAvg             Identifier = "avg"
	FuncDateOf          Identifier = "dateof"
	FuncFromJSON        Identifier = "fromjson"
	FuncMax             Identifier = "max"
	FuncMaxTimeuuid     Identifier = "maxtimeuuid"
	FuncMin             Identifier = "min"
	FuncMinTimeuuid     Identifier = "mintimeuuid"
	FuncNow             Identifier = "now"
	FuncSum             Identifier = "sum"
	FuncToDate          Identifier = "todate"
	FuncToJSON          Identifier = "tojson"
	FuncToken           Identifier = "token"
	FuncToTimestamp     Identifier = "totimestamp"
	FuncToUnixTimestamp Identifier = "tounixtimestamp"
	FuncUnixTimestampOf Identifier = "unixtimestampof"
	FuncUUID            Identifier = "uuid"
)

const (
	FuncBlobToAscii     Identifier = "blobtoascii"
	FuncBlobToBigint    Identifier = "blobtobigint"
	FuncBlobToBoolean   Identifier = "blobtoboolean"
	FuncBlobToCounter   Identifier = "blobtocounter"
	FuncBlobToDate      Identifier = "blobtodate"
	FuncBlobToDecimal   Identifier = "blobtodecimal"
	FuncBlobToDouble    Identifier = "blobtodouble"
	FuncBlobToFloat     Identifier = "blobtofloat"
	FuncBlobToInet      Identifier = "blobtoinet"
	FuncBlobToInt       Identifier = "blobtoint"
	FuncBlobToSmallint  Identifier = "blobtosmallint"
	FuncBlobToText      Identifier = "blobtotext"
	FuncBlobToTime      Identifier = "blobtotime"
	FuncBlobToTimestamp Identifier = "blobtotimestamp"
	FuncBlobToTimeuuid  Identifier = "blobtotimeuuid"
	FuncBlobToTinyint   Identifier = "blobtotinyint"
	FuncBlobToUuid      Identifier = "blobtouuid"
	FuncBlobToVarchar   Identifier = "blobtovarchar"
	FuncBlobToVarint    Identifier = "blobtovarint"
)

const (
	FuncAsciiToBlob     Identifier = "asciitoblob"
	FuncBigintToBlob    Identifier = "biginttoblob"
	FuncBooleanToBlob   Identifier = "booleantoblob"
	FuncCounterToBlob   Identifier = "countertoblob"
	FuncDateToBlob      Identifier = "datetoblob"
	FuncDecimalToBlob   Identifier = "decimaltoblob"
	FuncDoubleToBlob    Identifier = "doubletoblob"
	FuncFloatToBlob     Identifier = "floattoblob"
	FuncInetToBlob      Identifier = "inettoblob"
	FuncIntToBlob       Identifier = "inttoblob"
	FuncSmallintToBlob  Identifier = "smallinttoblob"
	FuncTextToBlob      Identifier = "texttoblob"
	FuncTimeToBlob      Identifier = "timetoblob"
	FuncTimestampToBlob Identifier = "timestamptoblob"
	FuncTimeuuidToBlob  Identifier = "timeuuidtoblob"
	FuncTinyintToBlob   Identifier = "tinyinttoblob"
	FuncUuidToBlob      Identifier = "uuidtoblob"
	FuncVarcharToBlob   Identifier = "varchartoblob"
	FuncVarintToBlob    Identifier = "varinttoblob"
)

const (
	KeyspaceDurableWrites Identifier = "durable_writes"
	KeyspaceReplication   Identifier = "replication"
)

const (
	TableBaseTimeSeconds              Identifier = "base_time_seconds"
	TableBloomFilterFPChance          Identifier = "bloom_filter_fp_chance"
	TableBucketHigh                   Identifier = "bucket_high"
	TableBucketLow                    Identifier = "bucket_low"
	TableCRCCheckChance               Identifier = "crc_check_chance"
	TableCaching                      Identifier = "caching"
	TableChunkLengthInKB              Identifier = "chunk_length_in_kb"
	TableClass                        Identifier = "class"
	TableComment                      Identifier = "comment"
	TableCompaction                   Identifier = "compaction"
	TableCompression                  Identifier = "compression"
	TableDCLocalReadRepairChance      Identifier = "dclocal_read_repair_chance"
	TableDefaultTimeToLive            Identifier = "default_time_to_live"
	TableEnabled                      Identifier = "enabled"
	TableGCGraceSeconds               Identifier = "gc_grace_seconds"
	TableKeys                         Identifier = "keys"
	TableMaxSSTableAgeDays            Identifier = "max_sstable_age_days"
	TableMaxThreshold                 Identifier = "max_threshold"
	TableMinSSTableSize               Identifier = "min_sstable_size"
	TableMinThreshold                 Identifier = "min_threshold"
	TableReadRepairChance             Identifier = "read_repair_chance"
	TableRowsPerPartition             Identifier = "rows_per_partition"
	TableSSTableSizeInMB              Identifier = "sstable_size_in_mb"
	TableTimestampResolution          Identifier = "timestamp_resolution"
	TableTombstoneCompactionInternal  Identifier = "tombstone_compaction_interval"
	TableTombstoneThreshold           Identifier = "tombstone_threshold"
	TableUncheckedTombstoneCompaction Identifier = "unchecked_tombstone_compaction"
)

func (i Identifier) String() string {
	return fmt.Sprintf("%q", string(i))
}
