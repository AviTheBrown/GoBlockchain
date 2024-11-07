# Go Blockchain Implementation 🚧 (Work In Progress)

A foundational blockchain implementation in Go focusing on core concepts and principles. This educational project is actively being developed to demonstrate and explore blockchain technology through hands-on code.

## Current Features 🎯
- Basic blockchain structure with proof-of-work mining
- SHA-256 hash generation and chain validation
- Configurable mining difficulty (leading zeros)
- Block generation time tracking
- Genesis block creation

## Planned Features 🔄
- Transactions and wallet implementation
- Merkle trees for efficient verification
- Distributed network capabilities
- P2P networking
- Consensus mechanisms
- Smart contracts
- Block persistence (database integration)
- Web interface for blockchain exploration

## Installation 🛠️

1. Clone the repository
```bash
git clone https://github.com/yourusername/go-blockchain.git
cd go-blockchain
```

2. Ensure Go is installed (version 1.16 or higher)
```bash
go version
```

3. Run the program
```bash
go run main.go
```

## Usage 💡

The program prompts for the number of leading zeros required for the proof-of-work:

```bash
Enter how many zeros the hash must start with:
> 4
```

Example output:
```
Genesis Block:
Id: 1
Timestamp: 1683924000000000000
Magic Number: 123456
Hash of the previous block:
0
Hash of the block:
0000a1b2c3d4e5f6...
Block was generating for 2 seconds

Block:
Id: 2
...
```

## How It Works 🔍

### Block Mining Process
1. User specifies required number of leading zeros
2. Program generates random magic numbers
3. Each magic number is used to create a potential block hash
4. When a hash with the required zeros is found, the block is created
5. Block is linked to previous block via hash chain
6. Generation time is recorded

### Validation
- Each block contains the previous block's hash
- Hash must start with specified number of zeros
- Genesis block has a previous hash of "0"

## Code Structure 📚

### Block Structure
```go
type Block struct {
    Id           int
    Magic_number int32
    Time_stamp   int64
    PrevHash     string
    Hash         string
    Gen_time     time.Duration
}
```

### Key Functions
```go
// Creates the genesis block
func CreateGenesisBlock(numZeros int) Block

// Creates a new block linked to the previous one
func CreateNewBlock(prevBlock Block, id int, numZeros int) Block

// Generates SHA-256 hash
func GeneratHash(id int, timestamp int64, magicNumber int32, preHash string, numZeros int) string

// Validates hash meets difficulty requirement
func validateHash(hash string, numZeros int) bool
```

## Educational Value 📖
This project is designed to help understand:
- Blockchain fundamentals
- Proof-of-work concepts
- Cryptographic hashing
- Chain validation
- Block generation and mining
- Timestamp and duration tracking

## Contributing 🤝
This is a work in progress, and contributions are welcome! Areas you can help with:
1. Implementing planned features
2. Improving documentation
3. Adding tests
4. Code optimization
5. Bug fixes

To contribute:
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## ⚠️ Disclaimer
This is an educational project designed for learning blockchain concepts. It's actively being developed and improved. Not intended for production use.

## License 📄
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact 📫
Your Name - [@yourusername](https://github.com/yourusername)

Project Link: [https://github.com/yourusername/go-blockchain](https://github.com/yourusername/go-blockchain)

---

Feel free to contribute, suggest features, or use for educational purposes! Star ⭐ the repository if you find it helpful!
