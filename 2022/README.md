## Advent of Code 2022

I'm using this year as an excuse to learn Go. So if (for some reason?) you're
following along, don't expect anything but amateurish Go hacking. Hopefully
it'll get better as we go along, but the problems will also get harder, so
we'll see.

### Running

I won't bother checking in any puzzle inputs. My personal workflow is to
generally right the sample input to a file called `demo` and the real input
to `input`. Then I run things like `go run . <demo` and `go run . <input`.
Call it a personal tic, but I tend to always write these to read from stdin
like that, and unless I have a good reason not to I always try to process
the input as a stream and only make a single pass, rather than reading all
lines into memory first.
