# Cosmos IBC Token Check

This is a simple script to check the balance of a token on a Cosmos IBC chain.

## Prepare
- Make sure you have the Go programming language installed on your machine.
- Make sure to have addresses.json file in the root directory with the following format:
```json
{
  "addresses": {
    "neutron-1": "neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c",
    "osmosis-1": "osmo1lzecpea0qxw5xae92xkm3vaddeszr278665crd",
    "phoenix-1": "terra1w7mtx2g478kkhs6pgynpcjpt6aw4930q34j36v",
    "stargaze-1": "stars1lzecpea0qxw5xae92xkm3vaddeszr278xas47w",
    "cosmoshub-4": "cosmos1lzecpea0qxw5xae92xkm3vaddeszr278jp8g4l",
    "stride-1": "stride1lzecpea0qxw5xae92xkm3vaddeszr2783285pn"
  }
}
```

## Building
```
go mod tidy
go build -o ibc-token-check
```

## Running
```
./ibc-token-check --config=addresses.json
```

## Output
```
neutron-1:
factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 113.000000, [neutron-1]
factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/lAsset, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/lAsset, 21.000000, [neutron-1]
ibc/C4A3E0BDA2A18D39FCB66C1D2945F6BF5A9714F0E5221D5E98976196B99F26E8, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 1.000000, [unknown]
ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 1900000.000000, [unknown]
ibc/DB12DE291358ACBF9FA7B9E710DF989AE17FD6638538BD4D70324A7C42A056CB, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 1.000000, [unknown]
osmosis-1:
ibc/544EC5AC9035F4E23397B9B877F0F17123F1F95B90A554BEBD5C0C16962835A7, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 10.000000, [unknown]
ibc/BE9E1A87A5A567F6C9E9A3655C0B204F82DF11CAC0D670461BC456BA97359E8D, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 1.000000, [unknown]
phoenix-1:
ibc/BD1DE0C000FC9ACF69AA38CBC766F81B54C038FF7A20ED80AABD36182FCB7FC8, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 2.000000, [unknown]
ibc/1716D72043436140079BDE06C8D2F76EEEA5397E661B635A9667285D2CC75F47, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 19.000000, [unknown]
stargaze-1:
cosmoshub-4:
ibc/3D459B55C08CE48681DB4E99101EE4BCF325A1FD02C21A85AB3BF62C9D8685BF, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 4.000000, [unknown]
stride-1:
ibc/7BC3C5F0FF1AB51D0F7F818F95CB118809C8D8D5B625E81EF0CCAC931AC359BD, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 2.000000, [unknown]
ibc/E9ED93B74C8A353A6433E9663B0781E96653B4C4718184682071CCBA08A2790D, factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 5.000000, [unknown]
TOTAL AMOUNTS:
factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset, 1900158.000000
factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/lAsset, 21.000000
```