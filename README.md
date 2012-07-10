grr is a naively simple file corruption generator.

It just skips a random number of bytes between 0 and ```j``` bytes of input and
then writes one random byte, looping until the end of the input file is reached.
Where ```j``` can be specified as an option and defaults to 32 which results 
roughly in 1 every 16 bytes corrupted.

It tries to conceal the introduced synthetic corruption by keeping files 
permissions, size and times unmodified.

It can also corrupt whole trees of files recursively when given directories as 
arguments.