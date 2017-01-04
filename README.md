# Salary
[![CircleCI](https://circleci.com/gh/aalvesjr/salary.svg?style=svg)](https://circleci.com/gh/aalvesjr/salary)

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr/salary
git clone https://github.com/aalvesjr/salary.git
```

## Using

Add the model `salary` at your project


```
import(
  salary "github.com/aalvesjr/salary"
)
```

And the function `func NewSalary(value, discount float32) Salary` will be available to create a struct `Salary`

```
s := salary.NewSalary(5000.00, 107.32)
```

With this `s` will be a struct `Salary` with the following attributes:

```
type Salary struct {
	Gross             float32
	Net               float32
	INSSRate          float32
	INSSBase          float32
	INSS              float32
	IRBase            float32
	IRRate            float32
	IRDiscount        float32
	IRWithoutDiscount float32
	IR                float32
	Discounts         float32
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request
