package viperjacket

import (
	"flag"
	"math"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	. "github.com/stretchr/testify/assert"
)

var (
	testBoolVal                   = false
	defaultBoolVal                = true
	testStringVal                 = "test 998877"
	defaultStringVal              = "Hello World!"
	testIntVal            int     = 998877
	defaultIntVal         int     = 112233
	testInt32Val          int32   = 2147483647
	defaultInt32Val       int32   = 112233
	testInt64Val          int64   = 9223372036854775807
	defaultInt64Val       int64   = 112233
	testUintVal           uint    = 998877
	defaultUintVal        uint    = 112233
	testUint32Val         uint32  = 4294967295
	defaultUint32Val      uint32  = 112233
	testUint64Val         uint64  = 18446744073709551615
	defaultUint64Val      uint64  = 112233
	testFloat64Val        float64 = 998.877
	defaultFloat64Val     float64 = 112.233
	testTimeVal                   = time.Date(2021, 9, 15, 15, 31, 48, 123, time.UTC)
	defaultTimeVal                = time.Now()
	testDurationVal               = 9 * time.Minute
	defaultDurationVal            = 12 * time.Second
	testIntSliceVal               = []int{9, 9, 8, 8, 7, 7}
	defaultIntSliceVal            = []int{1, 1, 2, 2, 3, 3}
	testStringSliceVal            = []string{"test1", "test2", "test3"}
	defaultStringSliceVal         = []string{"Hello", "World!", "Nice Day!"}
	testStringMapVal              = map[string]interface{}{
		"test1": "testValue",
		"test2": 998.877,
		"test3": false,
	}
	defaultStringMapVal = map[string]interface{}{
		"hello": "World!",
		"Nice":  float64(112233),
		"Day":   true,
	}
	testStringMapStringVal = map[string]string{
		"test1": "testValue1",
		"test2": "testValue2",
		"test3": "testValue3",
	}
	defaultStringMapStringVal = map[string]string{
		"Hello":   "World!",
		"Nice To": "Meet You",
		"Me":      "Too",
	}
	testStringMapSliceVal = map[string][]string{
		"test1": {"testValue1-1", "testValue1-2", "testValue1-3"},
		"test2": {"testValue2-1", "testValue2-2", "testValue3-3"},
		"test3": {"testValue3-1", "testValue2-2", "testValue3-3"},
	}
	defaultStringMapSliceVal = map[string][]string{
		"test1": {"Hello", "World!"},
		"test2": {"Nice", "To"},
		"test3": {"Meet", "You"},
	}
	testStringWithEnvVal               = "white cow is better than black cow."
	testDefaultStringWithEnvVal        = "white cow is bigger than black cow."
	testDefaultStringWithNoExistEnvVal = "white cow is  than black cow."
)

func TestImplementConfig(t *testing.T) {
	Implements(t, (*Config)(nil), new(ViperJacket),
		"It must implements of interface config.Config")
}

func TestLoadDefault(t *testing.T) {
	cfg := GetOrCreate()
	defer cfg.Close()

	Equal(t, testBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, testStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, testIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, testInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, testInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, testUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, testUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, testUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, testFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, testTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, testDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	// Equal(t, testIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal)) // doesn't work
	Equal(t, testStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	// Equal(t, testStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))  // doesn't work
	// Equal(t, testStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))  // doesn't work
	// Equal(t, testStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))  // doesn't work
}

func TestLoadConfigFileNotFound(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/no_exist.conf")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, defaultBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, defaultStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, defaultIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, defaultInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, defaultInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, defaultUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, defaultUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, defaultUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, defaultFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, defaultTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, defaultDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	// Equal(t, defaultIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal)) // doesn't work
	Equal(t, defaultStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	// Equal(t, defaultStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))  // doesn't work
	// Equal(t, defaultStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))  // doesn't work
	// Equal(t, defaultStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))  // doesn't work
}

func TestLoadUnsupportedExt(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/unsupported.test")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, defaultBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, defaultStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, defaultIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, defaultInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, defaultInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, defaultUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, defaultUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, defaultUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, defaultFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, defaultTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, defaultDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	// Equal(t, defaultIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal)) // doesn't work
	Equal(t, defaultStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	// Equal(t, defaultStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))  // doesn't work
	// Equal(t, defaultStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))  // doesn't work
	// Equal(t, defaultStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))  // doesn't work
}

func TestLoadInvalidFormat(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/invalid_format.yaml")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, defaultBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, defaultStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, defaultIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, defaultInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, defaultInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, defaultUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, defaultUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, defaultUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, defaultFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, defaultTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, defaultDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	// Equal(t, defaultIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal)) // doesn't work
	Equal(t, defaultStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	// Equal(t, defaultStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))  // doesn't work
	// Equal(t, defaultStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))  // doesn't work
	// Equal(t, defaultStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))  // doesn't work
}

func TestLoadJSON(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/test.json")
	os.Setenv("VIPER_CONFIG_TEST1", "white cow")
	os.Setenv("VIPER_CONFIG_TEST2", "black cow")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, testBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, testStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, testIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, testInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	// Equal(t, testInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, testUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, testUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	// Equal(t, testUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, testFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, testTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, testDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	Equal(t, testIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal))
	Equal(t, testStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	Equal(t, testStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))
	Equal(t, testStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))
	Equal(t, testStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))
	Equal(t, testStringWithEnvVal, cfg.GetString("test_viper_config_string_with_env_value", defaultStringVal))
	os.Setenv("WHAT", "bigger")
	Equal(t, testDefaultStringWithEnvVal, cfg.GetString("no_exist", "white cow is $WHAT than black cow."))
	os.Unsetenv("WHAT")
	Equal(t, testDefaultStringWithNoExistEnvVal, cfg.GetString("no_exist", "white cow is $WHAT than black cow."))

}

func TestLoadYAML(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/test.yaml")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, testBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, testStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, testIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, testInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, testInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, testUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, testUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, testUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, testFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, testTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, testDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	Equal(t, testIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal))
	Equal(t, testStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	Equal(t, testStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))
	Equal(t, testStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))
	Equal(t, testStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))
}

func TestLoadTOML(t *testing.T) {
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/test.toml")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	Equal(t, testBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, testStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, testIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, testInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, testInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, testUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, testUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	// Equal(t, testUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, testFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, testTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, testDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	Equal(t, testIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal))
	Equal(t, testStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	Equal(t, testStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))
	Equal(t, testStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))
	Equal(t, testStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))
}

func TestGetBool(t *testing.T) {
	var GetBoolTests = []struct {
		title         string
		input_path    string
		input_default bool
		expected      bool
	}{
		{"Reading true value", "TEST_VIPER_CONFIG_BOOL_TRUE", false, true},
		{"Reading false value", "TEST_VIPER_CONFIG_BOOL_FALSE", true, false},
		{"Reading no exists value, set to default", "NO_EXIST_KEY", true, true},
		{"Reading no exists value, set to default", "NO_EXIST_KEY", false, false},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/bool_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetBoolTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetBool(test.input_path, test.input_default))
		})
	}
}

func TestGetString(t *testing.T) {
	var GetStringTests = []struct {
		title         string
		input_path    string
		input_default string
		expected      string
	}{
		{"Reading string value", "TEST_VIPER_CONFIG_STRING", "Hello, World!", "TEST text 1234 !@#$%^&*(),./ ???????????????"},
		{"Reading bool to string value", "TEST_VIPER_CONFIG_BOOL", "Hello, World!", "false"},
		{"Reading int to string value", "TEST_VIPER_CONFIG_INT", "Hello, World!", "998877"},
		{"Reading intslice to string value", "TEST_VIPER_CONFIG_INTSLICE", "Hello, World!", ""},
		{"Reading stringslice to string value", "TEST_VIPER_CONFIG_STRINGSLICE", "Hello, World!", ""},
		{"Reading stringmap to string value", "TEST_VIPER_CONFIG_STRINGMAP", "Hello, World!", ""},
		{"Reading no exists value", "NO_EXISTS_KEY", "Hello, World!", "Hello, World!"},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/string_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetStringTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetString(test.input_path, test.input_default))
		})
	}
}

func TestGetInt(t *testing.T) {
	var GetIntTests = []struct {
		title         string
		input_path    string
		input_default int
		expected      int
	}{
		{"Reading int value", "TEST_VIPER_CONFIG_INT", defaultIntVal, 998877},
		// {"Reading int max value", "TEST_VIPER_CONFIG_MAX_INT", defaultIntVal, math.MaxInt},
		{"Reading int min value", "TEST_VIPER_CONFIG_MIN_INT", defaultIntVal, math.MinInt},
		// {"Reading int over value", "TEST_VIPER_CONFIG_OVER_INT", defaultIntVal, math.MaxInt},
		{"Reading int under value", "TEST_VIPER_CONFIG_UNDER_INT", defaultIntVal, math.MinInt},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultIntVal, defaultIntVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/int_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetIntTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetInt(test.input_path, test.input_default))
		})
	}
}

func TestGetInt32(t *testing.T) {
	var GetInt32Tests = []struct {
		title         string
		input_path    string
		input_default int32
		expected      int32
	}{
		{"Reading int32 value", "TEST_VIPER_CONFIG_INT32", defaultInt32Val, 998877},
		{"Reading int32 max value", "TEST_VIPER_CONFIG_MAX_INT32", defaultInt32Val, math.MaxInt32},
		{"Reading int32 min value", "TEST_VIPER_CONFIG_MIN_INT32", defaultInt32Val, math.MinInt32},
		// {"Reading int32 over value", "TEST_VIPER_CONFIG_OVER_INT32", defaultInt32Val, math.MaxInt32},
		{"Reading int32 under value", "TEST_VIPER_CONFIG_UNDER_INT32", defaultInt32Val, math.MinInt32},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultInt32Val, defaultInt32Val},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/int32_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetInt32Tests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetInt32(test.input_path, test.input_default))
		})
	}
}

func TestGetInt64(t *testing.T) {
	var GetInt64Tests = []struct {
		title         string
		input_path    string
		input_default int64
		expected      int64
	}{
		{"Reading int64 value", "TEST_VIPER_CONFIG_INT64", defaultInt64Val, 998877},
		// {"Reading int64 max value", "TEST_VIPER_CONFIG_MAX_INT64", defaultInt64Val, math.MaxInt64},
		{"Reading int64 min value", "TEST_VIPER_CONFIG_MIN_INT64", defaultInt64Val, math.MinInt64},
		// {"Reading int64 over value", "TEST_VIPER_CONFIG_OVER_INT64", defaultInt64Val, math.MaxInt64},
		{"Reading int64 under value", "TEST_VIPER_CONFIG_UNDER_INT64", defaultInt64Val, math.MinInt64},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultInt64Val, defaultInt64Val},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/int64_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetInt64Tests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetInt64(test.input_path, test.input_default))
		})
	}
}

func TestGetUint(t *testing.T) {
	var GetUintTests = []struct {
		title         string
		input_path    string
		input_default uint
		expected      uint
	}{
		{"Reading uint value", "TEST_VIPER_CONFIG_UINT", defaultUintVal, 998877},
		// {"Reading uint max value", "TEST_VIPER_CONFIG_MAX_UINT", defaultUintVal, math.MaxUint},
		{"Reading uint min value", "TEST_VIPER_CONFIG_MIN_UINT", defaultUintVal, 0},
		// {"Reading uint over value", "TEST_VIPER_CONFIG_OVER_UINT", defaultUintVal, math.MaxUint},
		{"Reading uint under value", "TEST_VIPER_CONFIG_UNDER_UINT", defaultUintVal, 0},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultUintVal, defaultUintVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/uint_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetUintTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetUint(test.input_path, test.input_default))
		})
	}
}

func TestGetUint32(t *testing.T) {
	var GetUint32Tests = []struct {
		title         string
		input_path    string
		input_default uint32
		expected      uint32
	}{
		{"Reading uint32 value", "TEST_VIPER_CONFIG_UINT32", defaultUint32Val, 998877},
		{"Reading uint32 max value", "TEST_VIPER_CONFIG_MAX_UINT32", defaultUint32Val, math.MaxUint32},
		{"Reading uint32 min value", "TEST_VIPER_CONFIG_MIN_UINT32", defaultUint32Val, 0},
		// {"Reading uint32 over value", "TEST_VIPER_CONFIG_OVER_UINT32", defaultUint32Val, 0},
		// {"Reading uint32 under value", "TEST_VIPER_CONFIG_UNDER_UINT32", defaultUint32Val, 0},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultUint32Val, defaultUint32Val},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/uint32_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetUint32Tests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetUint32(test.input_path, test.input_default))
		})
	}
}

func TestGetUint64(t *testing.T) {
	var GetUint64Tests = []struct {
		title         string
		input_path    string
		input_default uint64
		expected      uint64
	}{
		{"Reading uint64 value", "TEST_VIPER_CONFIG_UINT64", defaultUint64Val, 998877},
		// {"Reading uint64 max value", "TEST_VIPER_CONFIG_MAX_UINT64", defaultUint64Val, math.MaxUint64},
		{"Reading uint64 min value", "TEST_VIPER_CONFIG_MIN_UINT64", defaultUint64Val, 0},
		// {"Reading uint64 over value", "TEST_VIPER_CONFIG_OVER_UINT64", defaultUint64Val, math.MaxUint64},
		{"Reading uint64 under value", "TEST_VIPER_CONFIG_UNDER_UINT64", defaultUint64Val, 0},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultUint64Val, defaultUint64Val},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/uint64_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetUint64Tests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetUint64(test.input_path, test.input_default))
		})
	}
}

func TestGetFloat64(t *testing.T) {
	var GetFloat64Tests = []struct {
		title         string
		input_path    string
		input_default float64
		expected      float64
	}{
		{"Reading float64 value", "TEST_VIPER_CONFIG_FLOAT64", 112.233, 998.877},
		{"Reading no exists value", "NO_EXISTS_KEY", 112.233, 112.233},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/float64_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetFloat64Tests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetFloat64(test.input_path, test.input_default))
		})
	}
}

func TestGetTime(t *testing.T) {
	var GetTimeTests = []struct {
		title         string
		input_path    string
		input_default time.Time
		expected      time.Time
	}{
		{"Reading time value", "TEST_VIPER_CONFIG_TIME", defaultTimeVal, time.Date(2021, 9, 15, 15, 31, 48, 123, time.UTC)},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultTimeVal, defaultTimeVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/time_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetTimeTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetTime(test.input_path, test.input_default))
		})
	}
}

func TestGetDuration(t *testing.T) {
	var GetDurationTests = []struct {
		title         string
		input_path    string
		input_default time.Duration
		expected      time.Duration
	}{
		{"Reading duration value", "TEST_VIPER_CONFIG_DURATION", defaultDurationVal, 9 * time.Minute},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultDurationVal, defaultDurationVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/duration_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetDurationTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetDuration(test.input_path, test.input_default))
		})
	}
}

func TestGetIntSlice(t *testing.T) {
	var GetIntSliceTests = []struct {
		title         string
		input_path    string
		input_default []int
		expected      []int
	}{
		{"Reading intslice value", "TEST_VIPER_CONFIG_INTSLICE", defaultIntSliceVal, []int{998877, 665544, 332211}},
		{"Reading empty intslice value", "TEST_VIPER_CONFIG_EMPTY_INTSLICE", defaultIntSliceVal, []int{}},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultIntSliceVal, defaultIntSliceVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/intslice_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetIntSliceTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetIntSlice(test.input_path, test.input_default))
		})
	}
}

func TestGetStringSlice(t *testing.T) {
	var GetStringSliceTests = []struct {
		title         string
		input_path    string
		input_default []string
		expected      []string
	}{
		{"Reading stringslice value", "TEST_VIPER_CONFIG_STRINGSLICE", defaultStringSliceVal, []string{"test1", "test2", "test3"}},
		{"Reading empty stringslice value", "TEST_VIPER_CONFIG_EMPTY_STRINGSLICE", defaultStringSliceVal, []string{}},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultStringSliceVal, defaultStringSliceVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/stringslice_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetStringSliceTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetStringSlice(test.input_path, test.input_default))
		})
	}
}

func TestGetStringMap(t *testing.T) {
	var GetStringMapTests = []struct {
		title         string
		input_path    string
		input_default map[string]interface{}
		expected      map[string]interface{}
	}{
		{"Reading stringmap value", "TEST_VIPER_CONFIG_STRINGMAP", defaultStringMapVal,
			map[string]interface{}{
				"test1": "testValue",
				"test2": 998.877,
				"test3": false,
			}},
		{"Reading empty stringmap value", "TEST_VIPER_CONFIG_EMPTY_STRINGMAP", defaultStringMapVal, map[string]interface{}{}},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultStringMapVal, defaultStringMapVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/stringmap_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetStringMapTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetStringMap(test.input_path, test.input_default))
		})
	}
}

func TestGetStringMapString(t *testing.T) {
	var GetStringMapTests = []struct {
		title         string
		input_path    string
		input_default map[string]string
		expected      map[string]string
	}{
		{"Reading stringmap value", "TEST_VIPER_CONFIG_STRINGMAPSTRING", defaultStringMapStringVal,
			map[string]string{
				"test1": "testValue1",
				"test2": "testValue2",
				"test3": "testValue3",
			}},
		{"Reading empty stringmap value", "TEST_VIPER_CONFIG_EMPTY_STRINGMAPSTRING", defaultStringMapStringVal, map[string]string{}},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultStringMapStringVal, defaultStringMapStringVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/stringmapstring_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetStringMapTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetStringMapString(test.input_path, test.input_default))
		})
	}
}

func TestGetStringMapSlice(t *testing.T) {
	var GetStringMapStringTests = []struct {
		title         string
		input_path    string
		input_default map[string][]string
		expected      map[string][]string
	}{
		{"Reading stringmap value", "TEST_VIPER_CONFIG_STRINGMAPSLICE", defaultStringMapSliceVal,
			map[string][]string{
				"test1": {"testValue1-1", "testValue1-2", "testValue1-3"},
				"test2": {"testValue2-1", "testValue2-2", "testValue3-3"},
				"test3": {"testValue3-1", "testValue2-2", "testValue3-3"},
			}},
		{"Reading empty stringmap value", "TEST_VIPER_CONFIG_EMPTY_STRINGMAPSLICE", defaultStringMapSliceVal, map[string][]string{}},
		{"Reading no exists value", "NO_EXISTS_KEY", defaultStringMapSliceVal, defaultStringMapSliceVal},
	}

	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "resources/stringmapslice_test.json")
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	defer cfg.Close()

	for _, test := range GetStringMapStringTests {
		t.Run(test.title, func(t *testing.T) {
			Equal(t, test.expected, cfg.GetStringMapSlice(test.input_path, test.input_default))
		})
	}
}

func TestClose(t *testing.T) {
	cfg := &ViperJacket{viper: viper.New(), flag: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
	cfg.init()
	NoError(t, cfg.Close(), "Failed to close ViperJacket")
}

var expandTests = []struct {
	in, out string
}{
	{"", ""},
	{"$*", "all the args"},
	{"$$", "PID"},
	{"${*}", "all the args"},
	{"$1", "ARGUMENT1"},
	{"${1}", "ARGUMENT1"},
	{"now is the time", "now is the time"},
	{"$HOME", "/usr/gopher"},
	{"$home_1", "/usr/foo"},
	{"${HOME}", "/usr/gopher"},
	{"${H}OME", "(Value of H)OME"},
	{"A$$$#$1$H$home_1*B", "APIDNARGSARGUMENT1(Value of H)/usr/foo*B"},
	{"start$+middle$^end$", "start$+middle$^end$"},
	{"mixed$|bag$$$", "mixed$|bagPID$"},
	{"$", "$"},
	{"$}", "$}"},
	{"${", ""},  // invalid syntax; eat up the characters
	{"${}", ""}, // invalid syntax; eat up the characters
}

func TestExpand(t *testing.T) {
	os.Setenv("*", "all the args")
	os.Setenv("#", "NARGS")
	os.Setenv("$", "PID")
	os.Setenv("1", "ARGUMENT1")
	os.Setenv("HOME", "/usr/gopher")
	os.Setenv("H", "(Value of H)")
	os.Setenv("home_1", "/usr/foo")
	os.Setenv("_", "underscore")
	for _, test := range expandTests {
		result, _ := expand(test.in)
		if result != test.out {
			t.Errorf("Expand(%q)=%q; expected %q", test.in, result, test.out)
		}
	}
}
