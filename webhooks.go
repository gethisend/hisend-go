package hisend

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// VerifyWebhook verifies the HMAC-SHA256 signature of an incoming Hisend webhook.
//
// payload   – the raw request body bytes exactly as received.
// signature – the value of the "X-Hisend-Signature" request header.
// secret    – your Webhook Signing Secret (starts with "whsec_").
//
// Returns true if the signature is valid, false otherwise.
//
// Example:
//
//	func HisendHandler(w http.ResponseWriter, r *http.Request) {
//	    body, _ := io.ReadAll(r.Body)
//	    sig     := r.Header.Get("X-Hisend-Signature")
//	    secret  := os.Getenv("HISEND_WEBHOOK_SECRET")
//
//	    if !hisend.VerifyWebhook(body, sig, secret) {
//	        http.Error(w, "Invalid signature", http.StatusUnauthorized)
//	        return
//	    }
//	    // ... handle event
//	}
func VerifyWebhook(payload []byte, signature, secret string) bool {
	if len(payload) == 0 || signature == "" || secret == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expected := hex.EncodeToString(mac.Sum(nil))

	// Use hmac.Equal for constant-time comparison
	return hmac.Equal([]byte(signature), []byte(expected))
}
