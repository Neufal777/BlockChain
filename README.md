### Simple Blockchain in Go

## Usage

```golang
bc := NewBlockchain()

bc.AddBlock("0.6 BTC")
bc.AddBlock("0.3 BTC")
bc.AddBlock("0.1 BTC")
bc.AddBlock("0.7 BTC")

for _, block := range bc.blocks {

	fmt.Printf("Previous block hash: %x\n", block.PrevBlockHash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Actual block hash: %x\n\n", block.Hash)
}

```

## Output
```
Previous block hash: 
Data: Genesis Block
Actual block hash: ff43c00c7033390f0c4256a97897fb42bff3d327366589ac931ccd180c78060f

Previous block hash: ff43c00c7033390f0c4256a97897fb42bff3d327366589ac931ccd180c78060f
Data: 0.6 BTC
Actual block hash: 41b11f02acabaf048229c67f678dda12d87472599bf7d6f773a36820c75712d8

Previous block hash: 41b11f02acabaf048229c67f678dda12d87472599bf7d6f773a36820c75712d8
Data: 0.3 BTC
Actual block hash: a52ade72f26406a9383fe9eae1521e0d6b325abe05d45363cb3f2e0209121845

Previous block hash: a52ade72f26406a9383fe9eae1521e0d6b325abe05d45363cb3f2e0209121845
Data: 0.1 BTC
Actual block hash: 404998f34e638924473acacdc2cf5ce005598d758c8337f02dcfed7d138c54dd

Previous block hash: 404998f34e638924473acacdc2cf5ce005598d758c8337f02dcfed7d138c54dd
Data: 0.7 BTC
Actual block hash: ebd2c3f644b7f09296deb460efa5905b11d522fd8d97ea4dc2a09e41f0715bad

```
