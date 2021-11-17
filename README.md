# Viper-Jacket: viper config for Wire-Jacket
[![GoDoc][doc-img]][doc] [![Github release][release-img]][release] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

Jacket of spf13/viper: config for wire-jacket.
Simplified env-based go configuration package using viper.
- bang9211/wire-jacket : https://github.com/bang9211/wire-jacket
- spf13/viper : https://github.com/spf13/viper

# Installation
Install Wire-Jacket by running:
```
go get github.com/bang9211/viper-jacket
```
and ensuring that $GOPATH/bin is added to your $PATH.

# Example
Just call GetOrCreate().

Viper-Jacket uses app.conf by default, or you can specify config file with '--conf' flag explicitly.

It can get config value of the file or environment variable using get-or-else style getter.

```go
cfg := GetOrCreate()
defer cfg.Close()

cfg.GetBool("test_viper_config_bool_value", true)
cfg.GetString("test_viper_config_string_value", "default value")
cfg.GetInt("test_viper_config_int_value", 12345)
cfg.GetInt32("test_viper_config_int32_value", 12345)
cfg.GetInt64("test_viper_config_int64_value", 12345)
cfg.GetUint("test_viper_config_uint_value", 12345)
cfg.GetUint32("test_viper_config_uint32_value", 12345)
cfg.GetUint64("test_viper_config_uint64_value", 12345)
cfg.GetFloat64("test_viper_config_float64_value", 123.456)
cfg.GetTime("test_viper_config_time_value", time.Date(2021, 9, 15, 15, 31, 48, 123, time.UTC))
cfg.GetDuration("test_viper_config_duration_value", 12 * time.Second)
```

If the string values of GetString(), GetStringSlice(), GetStringSliceMap() contain environment variables, these will be expanded automatically.

For example $ENV_VAR=viperjacket
- /home/test/$ENV_VAR => /home/test/viperjacket
- /home/test/${ENV_VAR}/temp => /home/test/viperjacket/temp

Some complex Getter failed because viper doesn't support yet.
It can be checked in viperjacket_test.go.

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://pkg.go.dev/github.com/bang9211/viper-jacket

[release-img]: https://img.shields.io/github/release/bang9211/viper-jacket.svg
[release]: https://github.com/bang9211/viper-jacket/releases

[ci-img]: https://github.com/bang9211/viper-jacket/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/bang9211/viper-jacket/actions/workflows/go.yml

[cov-img]: https://codecov.io/gh/bang9211/viper-jacket/branch/main/graph/badge.svg
[cov]: https://codecov.io/gh/bang9211/viper-jacket/branch/main

[report-card-img]: https://goreportcard.com/badge/github.com/bang9211/viper-jacket
[report-card]: https://goreportcard.com/report/github.com/bang9211/viper-jacket

[release-policy]: https://golang.org/doc/devel/release.html#policy
