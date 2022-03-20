# exitanalyser

The analyzer that parses AST and finds explicit os.Exit calls.

# Usage:

1. import "github.com/DrGermanius/exitanalyser"
2. Put analyser into multichecker `multichecker.Main(exitanalyser.Analyzer)`
3. Run multichecker
