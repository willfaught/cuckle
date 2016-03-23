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

func (i Identifier) String() string {
	return fmt.Sprintf("%q", i)
}
