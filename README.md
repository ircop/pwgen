# pwgen
Friendly password generator in golang

This is a simple tool for generating semantic, easy-to-remember and hard-to-bruteforce passwords.

It contains it's own dictionary with simple words (which can be replaced in runtime), and generates passwords using simple semantic rules.

Example usage:

```go

p := pwgen.New(pwgen.Defaults())
for i := 0; i < 5; i++ {
		pw := p.Generate()
		fmt.Printf("password: '%s'\n", pw)
}
```

Result is:
```
password: '2man.bipbu'
password: 'wouNd1-tekba'
password: '1wound#nikSo'
password: 'test-poRna6'
password: 'smile@mitdo8'
```

You can pass your own dictionary:
```go
p := pwgen.NewWithDict(Defaults(), []string{"testword"})
for i := 0; i < 5; i++ {
		pw := p.Generate()
		fmt.Printf("password: '%s'\n", pw)
	}
```

result:
```
password: 'testWorD.1tupto'
password: '0testwOrd%retru'
password: 'TEstwoRd4-gorco'
password: 'tEstWord-7kutma'
password: '4testword.bekge'
```

And you can tune some parameters with config struct:
```
cfg := pwgen.Defaults()
cfg.Upper = true        // dont mix lower/uppercase
cfg.UpperChance = 3     // uint [0-9]: chance to make character uppercase, for each letter
cfg.Numbers = true      // to, or not to add random number to some part of password
cfg.Delims = []string{"-","^","_"}  // customize delimeters
```

