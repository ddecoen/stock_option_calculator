# Build a Stock Option Calculator in Go

## Exercise Cost Calculation

The calculator will ask the user for:
- Number of options
- Strike price

It will then calculate the exercise cost (number of options × strike price).

## Future Value Calculation

The calculator will ask the user to enter:
- Future share value

It will then calculate the future value of the options (number of options × future share value).

## How to Run

### Prerequisites
- Go installed on your system (version 1.16 or later)

### Running the Calculator

1. Clone this repository:
   ```bash
   git clone https://github.com/ddecoen/stock_option_calculator.git
   cd stock_option_calculator
   ```

2. Run the calculator:
   ```bash
   go run calculator.go
   ```

3. Follow the prompts to enter:
   - Number of options
   - Strike price per option
   - Expected sales price per option

4. The calculator will display:
   - Gain per option
   - Total potential gain