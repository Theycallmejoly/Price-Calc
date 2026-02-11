
---

```md
# Tax Included Prices Job

A simple Go tool for calculating prices with tax included.  
This project reads prices from a text file, applies a tax rate, and displays the results.

It is useful for practicing file processing, package structure, and basic financial calculations in Go.

---

## Features

- Load prices from a text file (`prices.txt`)
- Calculate prices with a defined tax rate
- Print results as a map (original price → price with tax)
- Modular project structure with separate packages

---

## Project Structure

```

.
├── conversion/      # Data conversion utilities (e.g. string to float)
├── fileManager/     # File management (reading and writing data)
├── prices/          # Core logic for price calculation
├── prices.txt       # Input file with prices
├── main.go          # Program entry point
└── README.md

````

---

## Main Class

### TaxIncludedPricesJob

This class manages the main operations of the project.

### Methods

- `LoadData()`  
  Loads prices from the file and converts them to floating-point numbers.

- `Process()`  
  Calculates tax-included prices and prints the results.

---

## Installation and Run

### 1. Clone the repository

```bash
git clone https://github.com/username/repo-name.git
````

### 2. Go to the project directory

```bash
cd repo-name
```

### 3. Prepare the `prices.txt` file

Example:

```
100.0
200.0
300.0
```

### 4. Run the project

```bash
go run main.go
```

---

## Usage

You can define your own tax rate in `main.go`:

```go
package main

import (
	"example.com/prices"
	"example.com/fileManager"
)

func main() {
	taxRate := 0.15 // 15% tax rate
	job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
	job.Process()
}
```

---

## Sample Output

```
Original Price: 100.0 -> Price with Tax: 115.0
Original Price: 200.0 -> Price with Tax: 230.0
Original Price: 300.0 -> Price with Tax: 345.0
```

---

## Notes

Possible improvements:

* Save results to a file
* Allow user input for tax rate
* Better error handling
* Add unit tests

```

---

اگر بخوای، می‌تونم نسخه **خیلی حرفه‌ای‌تر** (مثل پروژه‌های واقعی GitHub) یا **مینیمال‌تر** هم برات بسازم 🔥
```
