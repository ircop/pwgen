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
password: 'Bread3.gottu8'
password: 'cOmplEte4#sedbu2'
password: 'print.ropde'
password: 'Sign#moRmi'
password: 'tasTe@Cegka'
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
password: 'Testword-teske'
password: 'TestWord#sedbu2'
password: 'testwoRd.decbE'
password: 'tesTword%gosDe'
password: 'teStWOrd.nInda8'
```

And you can tune some parameters with config struct:
```
cfg := pwgen.Defaults()
cfg.Upper = true        // dont mix lower/uppercase
cfg.UpperChance = 3     // uint [0-9]: chance to make character uppercase, for each letter
cfg.Numbers = true      // to, or not to add random numbers
cfg.NumChance = 5       // chance to add number, uint [0-9]
cfg.Delims = []string{"-","^","_"}  // customize delimeters
```

