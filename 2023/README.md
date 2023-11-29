## Advent of Code 2023

After using last year as an excuse to learn Go, I didn't really touch Go
again for 11 months. So, I guess it's time to learn Go again? Don't expect
anything but amateurish Go hacking here.

### Running

I won't bother checking in any puzzle inputs. My personal workflow is to
generally write the sample input to a file called `demo` and the real input
to `input`. Then I run things like `go run . <demo` and `go run . <input`.
Call it a personal tic, but I tend to always write these to read from stdin
like that, and unless I have a good reason not to I always try to process
the input as a stream and only make a single pass, rather than reading all
lines into memory first.
