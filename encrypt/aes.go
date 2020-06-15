package encrypt

import "bytes"

// 高级加密标准 (Advanced Encryption Standard, AES)

//16, 24, 32 位字符串的话， 分别对应AES-128, AES-192, AES-256加密方法
// key 不能泄漏
var PwdKey = []byte("DIS**#KKKDJJSKDI")

// PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个， 然后合并成新的字节切片🔙
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEcrypt(origData) {

}
