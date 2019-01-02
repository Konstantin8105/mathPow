# mathPow

optimization math.Pow

```
name                                   time/op
/SIMPLE/_____pow/x:3.14e+00/y:7.0-4    425ns ± 1%
/SQRT2+/_____pow/x:3.14e+00/y:0.5-4    196ns ± 0%
/SQRT2-/_____pow/x:3.14e+00/y:-0.5-4   242ns ± 1%
/SQRT3/_____pow/x:3.14e+00/y:-0.3-4    719ns ± 1%
/SQRT8+/_____pow/x:3.14e+00/y:0.1-4    680ns ± 0%
/SQRT8-/_____pow/x:3.14e+00/y:-0.1-4   721ns ± 1%
/SQRT32+/_____pow/x:3.14e+00/y:0.0-4   679ns ± 1%
/SQRT32-/_____pow/x:3.14e+00/y:-0.0-4  721ns ± 1%
/SQRT64+/_____pow/x:3.14e+00/y:0.0-4   682ns ± 1%
/SQRT64-/_____pow/x:3.14e+00/y:-0.0-4  720ns ± 1%
/BASE/_____pow/x:1.23e-05/y:-2.5-4     770ns ± 0%
/PI/_____pow/x:3.14e+00/y:126.0-4      498ns ± 1%
/PI/_____pow/x:3.14e+00/y:125.0-4      496ns ± 1%
/PI/_____pow/x:3.14e+00/y:2.0-4        414ns ± 0%
/PI/_____pow/x:3.14e+00/y:-2.0-4       456ns ± 1%
/PI/_____pow/x:3.14e+00/y:-125.0-4     536ns ± 1%
/PI/_____pow/x:3.14e+00/y:-126.0-4     537ns ± 1%
/SMALL/_____pow/x:1.23e-05/y:126.0-4   507ns ± 1%
/SMALL/_____pow/x:1.23e-05/y:125.0-4   506ns ± 1%
/SMALL/_____pow/x:1.23e-05/y:2.0-4     416ns ± 1%
/SMALL/_____pow/x:1.23e-05/y:-2.0-4    454ns ± 1%
/SMALL/_____pow/x:1.23e-05/y:-125.0-4  533ns ± 0%
/SMALL/_____pow/x:1.23e-05/y:-126.0-4  533ns ± 1%
/SIMPLE/_pow_win/x:3.14e+00/y:7.0-4    202ns ± 0%
/SQRT2+/_pow_win/x:3.14e+00/y:0.5-4    278ns ± 1%
/SQRT2-/_pow_win/x:3.14e+00/y:-0.5-4   310ns ± 1%
/SQRT3/_pow_win/x:3.14e+00/y:-0.3-4    787ns ± 1%
/SQRT8+/_pow_win/x:3.14e+00/y:0.1-4    386ns ± 1%
/SQRT8-/_pow_win/x:3.14e+00/y:-0.1-4   419ns ± 1%
/SQRT32+/_pow_win/x:3.14e+00/y:0.0-4   496ns ± 1%
/SQRT32-/_pow_win/x:3.14e+00/y:-0.0-4  525ns ± 1%
/SQRT64+/_pow_win/x:3.14e+00/y:0.0-4   549ns ± 1%
/SQRT64-/_pow_win/x:3.14e+00/y:-0.0-4  579ns ± 1%
/BASE/_pow_win/x:1.23e-05/y:-2.5-4     842ns ± 0%
/PI/_pow_win/x:3.14e+00/y:126.0-4      293ns ± 0%
/PI/_pow_win/x:3.14e+00/y:125.0-4      285ns ± 1%
/PI/_pow_win/x:3.14e+00/y:2.0-4        177ns ± 1%
/PI/_pow_win/x:3.14e+00/y:-2.0-4       216ns ± 1%
/PI/_pow_win/x:3.14e+00/y:-125.0-4     326ns ± 1%
/PI/_pow_win/x:3.14e+00/y:-126.0-4     332ns ± 1%
/SMALL/_pow_win/x:1.23e-05/y:126.0-4   686ns ± 1%
/SMALL/_pow_win/x:1.23e-05/y:125.0-4   679ns ± 1%
/SMALL/_pow_win/x:1.23e-05/y:2.0-4     179ns ± 0%
/SMALL/_pow_win/x:1.23e-05/y:-2.0-4    216ns ± 0%
/SMALL/_pow_win/x:1.23e-05/y:-125.0-4  691ns ± 0%
/SMALL/_pow_win/x:1.23e-05/y:-126.0-4  698ns ± 1%
/SIMPLE/math.Pow/x:3.14e+00/y:7.0-4    375ns ± 0%
/SQRT2+/math.Pow/x:3.14e+00/y:0.5-4    144ns ± 0%
/SQRT2-/math.Pow/x:3.14e+00/y:-0.5-4   189ns ± 1%
/SQRT3/math.Pow/x:3.14e+00/y:-0.3-4    663ns ± 1%
/SQRT8+/math.Pow/x:3.14e+00/y:0.1-4    624ns ± 0%
/SQRT8-/math.Pow/x:3.14e+00/y:-0.1-4   666ns ± 1%
/SQRT32+/math.Pow/x:3.14e+00/y:0.0-4   626ns ± 1%
/SQRT32-/math.Pow/x:3.14e+00/y:-0.0-4  665ns ± 0%
/SQRT64+/math.Pow/x:3.14e+00/y:0.0-4   623ns ± 1%
/SQRT64-/math.Pow/x:3.14e+00/y:-0.0-4  663ns ± 1%
/BASE/math.Pow/x:1.23e-05/y:-2.5-4     704ns ± 1%
/PI/math.Pow/x:3.14e+00/y:126.0-4      464ns ± 1%
/PI/math.Pow/x:3.14e+00/y:125.0-4      448ns ± 0%
```
