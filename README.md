grr is a naively simple file corruption generator

It just skips a random number of bytes between 0 and ```j``` of input and then writes one random byte, looping until the end of the input file is reached.

```j``` can be specified as an option and defaults to 32.

