// --- internal/pengguna/domain/entity.go ---
package domain

import "time"

type RefreshToken struct {
	ID         string `gorm:"primaryKey;type:uuid"`
	PenggunaID string `gorm:"index"`
	Token      string `gorm:"uniqueIndex"`
	ExpiresAt  time.Time
	CreatedAt  time.Time
}
