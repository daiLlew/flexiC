# flexiC
Command line client for adding durations - useful for calculating your flexi time.


## Getting started
Requires Go 1.13 or higher.

### Install
:warning: FlexiC uses Go Modules so it should be cloned to a location outside your GOPATH:
```
git clone git@github.com:daiLlew/flexiC.git
cd flexiC
go install github.com/daiLlew/flexiC
```

### Run
Once installed simply run:
```
flexiC
```
You should be presented with:
```
[flexiC] Input start end times using the format "hhMM hhMM".
[flexiC] Enter return after each pair to submit (repeat as necessary).
[flexiC] Enter "d" to complete and display the total duration.
>
```
Enter times in pairs using the format `hhMM hhMM` for example:
```
> 0800 0932
```
Hitting enter will submit the times - repeat as many times as necessary. Once you have entered all times enter `d` to complete and display the total duration.
```
[flexiC] Input start end times using the format "hhMM hhMM".
[flexiC] Enter return after each pair to submit (repeat as necessary).
[flexiC] Enter "d" to complete and display the total duration.
> 0800 1230
> 1300 1547
> d
>
[flexiC] Total: 07h:17m
```