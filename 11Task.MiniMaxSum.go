//https://www.hackerrank.com/challenges/mini-max-sum/problem

func miniMaxSum(arr []int32) {
var minTmp, maxTmp, sum int32 = arr[0], arr[0], arr[0]

for i := 1; i < len(arr); i++ {
sum += arr[i]
if arr[i] < minTmp {
minTmp = arr[i]
}
if arr[i] > maxTmp {
maxTmp = arr[i]
}
}

min := sum - maxTmp
max := sum - minTmp

fmt.Println(min, max)

}
