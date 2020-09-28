//https://www.hackerrank.com/challenges/save-the-prisoner/problem
func saveThePrisoner(n int32, m int32, s int32) int32 {
return ((s + m - 2) % n) + 1
}