Go language is used for the project and the source code is contained is src folder

use "go run main.go -input ./datasets/[dataset].txt -algo [kruskal|boruvaka]" command to execute the code.
THREAD_COUNT can be customized using .env file or by using THREAD_COUNT=[no of threads] in the above "go run" command.

prims algorithm is implemented in cpp (./src/prims.cpp)
use "g++ prims.cpp -fopenmp && ./a.out" to run the code
