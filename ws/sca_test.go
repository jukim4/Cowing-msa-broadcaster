import "crypto/md5"
func hash(data string) string {
    h := md5.Sum([]byte(data)) // ❌ MD5는 안전하지 않음
    return string(h[:])
}