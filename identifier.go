package cuckle

import "fmt"

// Identifier is a double-quoted identifier.
type Identifier string

// General functions.
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

// Functions that convert from blob to other types.
const (
	FuncBlobToASCII     Identifier = "blobtoascii"
	FuncBlobToBigInt    Identifier = "blobtobigint"
	FuncBlobToBoolean   Identifier = "blobtoboolean"
	FuncBlobToCounter   Identifier = "blobtocounter"
	FuncBlobToDate      Identifier = "blobtodate"
	FuncBlobToDecimal   Identifier = "blobtodecimal"
	FuncBlobToDouble    Identifier = "blobtodouble"
	FuncBlobToFloat     Identifier = "blobtofloat"
	FuncBlobToInet      Identifier = "blobtoinet"
	FuncBlobToInt       Identifier = "blobtoint"
	FuncBlobToSmallInt  Identifier = "blobtosmallint"
	FuncBlobToText      Identifier = "blobtotext"
	FuncBlobToTime      Identifier = "blobtotime"
	FuncBlobToTimeUUID  Identifier = "blobtotimeuuid"
	FuncBlobToTimestamp Identifier = "blobtotimestamp"
	FuncBlobToTinyInt   Identifier = "blobtotinyint"
	FuncBlobToUUID      Identifier = "blobtouuid"
	FuncBlobToVarChar   Identifier = "blobtovarchar"
	FuncBlobToVarInt    Identifier = "blobtovarint"
)

// Functions that convert from other types to blob.
const (
	FuncASCIIToBlob     Identifier = "asciitoblob"
	FuncBigIntToBlob    Identifier = "biginttoblob"
	FuncBooleanToBlob   Identifier = "booleantoblob"
	FuncCounterToBlob   Identifier = "countertoblob"
	FuncDateToBlob      Identifier = "datetoblob"
	FuncDecimalToBlob   Identifier = "decimaltoblob"
	FuncDoubleToBlob    Identifier = "doubletoblob"
	FuncFloatToBlob     Identifier = "floattoblob"
	FuncInetToBlob      Identifier = "inettoblob"
	FuncIntToBlob       Identifier = "inttoblob"
	FuncSmallIntToBlob  Identifier = "smallinttoblob"
	FuncTextToBlob      Identifier = "texttoblob"
	FuncTimeToBlob      Identifier = "timetoblob"
	FuncTimeUUIDToBlob  Identifier = "timeuuidtoblob"
	FuncTimestampToBlob Identifier = "timestamptoblob"
	FuncTinyIntToBlob   Identifier = "tinyinttoblob"
	FuncUUIDToBlob      Identifier = "uuidtoblob"
	FuncVarCharToBlob   Identifier = "varchartoblob"
	FuncVarIntToBlob    Identifier = "varinttoblob"
)

// Keyspace properties.
const (
	KeyspaceDurableWrites Identifier = "durable_writes"
	KeyspaceReplication   Identifier = "replication"
)

// Table properties.
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
