# col selects a column from a csv from input

    printf 'a,b\nc,d\nf,g' | go run col.go 1

It outputs

    b
    d
    g

# TODO: use some formula inline

   printf 'item,quantity,price\nbeer,6,2.5\ncoffee,1,10.5' | go run col.go "$1 * $2"

# TODO: build bin