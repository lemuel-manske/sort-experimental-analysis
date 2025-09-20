rm -rf results
mkdir results

for sortAlgorithm in bubblesort mergesort quicksort
do

  for size in 1000 10000 100000
  do
    resultFile="./results/${sortAlgorithm}_${size}.txt"

    touch "${resultFile}"

    echo "Starting :: ${sortAlgorithm} with size ${size}"

    {
      for i in $(seq 1 10); do
        go test ./${sortAlgorithm} -bench=. -args -size=${size}
      done
    } >> "${resultFile}"

    echo "Finishing :: ${sortAlgorithm} with size ${size}"
  done
done

