# GOSTR

gostr provides a collection of string helper functions ported to [GO (golang)](http://golang.org/) from PHP Laravel Framework

## Install
```bash
go get github.com/danyelkeddah/gostr
```

## Usage

```go
// Return the remainder of a string after the first occurrence of a given value.
gostr.After("Lorem ipsum dolor sit amet", "dolor") //  sit amet

// Return the remainder of a string after the last occurrence of a given value.
gostr.AfterLast("Lorem ipsum dolor sit dolor amet", "dolor") // amet

// Get the portion of a string before the first occurrence of a given value.
gostr.Before("Lorem ipsum dolor sit amet", "dolor") // Lorem ipsum

// Get the portion of a string before the last occurrence of a given value.
gostr.BeforeLast("Lorem ipsum dolor sit dolor amet", "dolor") // Lorem ipsum dolor sit

// Get the portion of a string between two given values.
gostr.Between("Lorem ipsum dolor sit dolor amet", "dolor", "dolor") //  sit

// Determine if a given string contains any of the the given substrings.
gostr.Contains("Lorem ipsum dolor sit dolor amet", []string{"amet", "foo"}) // true

// Determine if a given string contains all given substrings.
gostr.ContainsAll("Lorem ipsum dolor sit dolor amet", []string{"amet", "foo"}) // false

// Determine if a given string ends with any of the the given substrings.
gostr.EndsWith("Lorem ipsum dolor sit dolor amet", []string{"amet", "foo"}) // true

// Determine if a given string is a valid UUID.
gostr.IsUUID("6ba7b814-9dad-11d1-80b4-00c04fd430c8") // true

// Return the length of the given string.
gostr.Length("abc") // 3

// Convert the given string to lower-case.
gostr.Lower("ABC") // abc

// Limit the number of words in a string.
gostr.Words("Lorem ipsum dolor sit dolor amet", 3, "...") // Lorem ipsum dolor...

// Pad the right side of a string with another.
gostr.PadRight("Lorem", 10, "*") // Lorem*****

// Pad the left side of a string with another.
gostr.PadLeft("Lorem", 10, "*") // *****Lorem

// Generate a more truly "random" alpha-numeric string.
gostr.Random(16) // hZHYDOhf8GqQQjnx

// Repeat the given string.
gostr.Repeat("Lorem", 5); // LoremLoremLoremLoremLorem

// Convert the given string to upper-case.
gostr.Upper("abc"); // ABC
```

## Credits

-   [Danyel Alkeddah](https://github.com/danyelkeddah)
-   [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
