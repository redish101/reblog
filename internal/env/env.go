package env

import (
	"crypto/ed25519"
	"crypto/sha256"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Dev                      bool
	Port                     int
	DatabaseURL              string
	SecretKey                string
	OwnerEmail               string
	JWTPrivateKey            ed25519.PrivateKey
	JWTPublicKey             ed25519.PublicKey
	GitHubClientID           string
	GitHubSecret             string
	GitHubRedirectURL        string
	FrontendURL              string
	CopyrightContractAddress string
	CopyrightRPCURL          string
	IPFSAPIURL               string
)

func Init() {
	godotenv.Load()

	Dev = getAsBool("DEV", false)
	Port = getAsInt("PORT", 3000)
	DatabaseURL = getAsString("DATABASE_URL", "postgres://user:password@localhost:5432/reblog?sslmode=disable")
	SecretKey = getAsString("SECRET_KEY", "reblog")
	OwnerEmail = getAsString("OWNER_EMAIL", "i@redish101.top") // 不填就是我的（雾）
	GitHubClientID = getAsString("GITHUB_CLIENT_ID", "")
	GitHubSecret = getAsString("GITHUB_SECRET", "")
	GitHubRedirectURL = getAsString("GITHUB_REDIRECT_URL", "http://localhost:3000/api/v1/auth/github/callback")
	FrontendURL = getAsString("FRONTEND_URL", "http://localhost:3000")
	CopyrightContractAddress = getAsString("COPYRIGHT_CONTRACT_ADDRESS", "0x3bc13fB85EfB2b7C8Ab6D1350aCc3f9e91D8b234")
	CopyrightRPCURL = getAsString("COPYRIGHT_RPC_URL", "https://sepolia.infura.io")
	IPFSAPIURL = getAsString("IPFS_API_URL", "http://localhost:5001")

	// 从 SecretKey 生成 ed25519 密钥对
	generateKeyPair()
}

// generateKeyPair 从 SecretKey 生成 ed25519 密钥对
func generateKeyPair() {
	// 使用 SHA256 对 SecretKey 进行哈希，得到32字节的种子
	hash := sha256.Sum256([]byte(SecretKey))

	// 从种子生成 ed25519 私钥
	JWTPrivateKey = ed25519.NewKeyFromSeed(hash[:])

	// 从私钥获取公钥
	JWTPublicKey = JWTPrivateKey.Public().(ed25519.PublicKey)
}

func getAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func getAsBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func getAsString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
