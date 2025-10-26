# Hello World

## Greet someone

### Inputs

| Name         | Description  | Default | Required |
| ------------ | ------------ | ------- | -------- |
| who-to-greet | Who to greet | World   | true     |

### Outputs

| Name          | Description   | Value                                                      |
| ------------- | ------------- | ---------------------------------------------------------- |
| random-number | Random number | ${{ steps.random-number-generator.outputs.random-number }} |

## Steps

- Set Greeting
- Random Number Generator
- Set GitHub Path
- Run goodbye.sh
