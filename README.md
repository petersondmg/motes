### Motes

Motes is a digital communication protocol based on **M**usical N**otes**

**Working:**

- Each byte is converted to a base 5 number
- Each digit represents a musical note, starting from note A 440z (1) to E (5).
- Zero (0) is F. G is a byte separator.
- Doubled notes are not allowed sequentially. In this case the repetition is replaced by G#.
- Each byte is converted to a musical note (Mote) then a G is appended indicating the end of the byte.

**Example:**

Working example using a tone generator

```sh
go run example.go motes.go "hello world"
```

**Pourpose:**

Just experimentation.


