# Reading Colour Values

Divide each channel by `clear` to get a lighting-independent ratio:

```
redRatio   = red   / clear
greenRatio = green / clear
blueRatio  = blue  / clear
```

A ratio near 1.0 on a single channel means that colour dominates the scene.
