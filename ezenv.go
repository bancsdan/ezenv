package ezenv

import (
	"os"
	"strconv"
)

func GetBool(key string) (bool, error) {
	str := os.Getenv(key)
	return strconv.ParseBool(str)
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetUint(key string) (uint, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func GetUint8(key string) (uint8, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(val), nil
}

func GetUint16(key string) (uint16, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(val), nil
}

func GetUint32(key string) (uint32, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func GetUint64(key string) (uint64, error) {
	str := os.Getenv(key)
	return strconv.ParseUint(str, 10, 64)
}

func GetInt(key string) (int, error) {
	str := os.Getenv(key)
	return strconv.Atoi(str)
}

func GetInt8(key string) (int8, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(val), nil
}

func GetInt16(key string) (int16, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(val), nil
}

func GetInt32(key string) (int32, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func GetInt64(key string) (int64, error) {
	str := os.Getenv(key)
	return strconv.ParseInt(str, 10, 64)
}

func GetFloat32(key string) (float32, error) {
	str := os.Getenv(key)
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}

func GetFloat64(key string) (float64, error) {
	str := os.Getenv(key)
	return strconv.ParseFloat(str, 64)
}
