# Hello World

## Greet someone

### Inputs

| Name | Description | Default | Required |
|------|-------------|---------|----------|
| `greeting` | Greeting to use | `hello` | false |
| `who-to-greet` | Who to greet | `World` | true |


### Outputs

| Name | Description | Value |
|------|-------------|-------|
| `random-number` | Random number | `${{ steps.random-number-generator.outputs.random-number }}` |
| `test-action` | Just a test | `no value` |


## Steps

- Set Greeting
- Random Number Generator
- Set GitHub Path
- Run goodbye.sh