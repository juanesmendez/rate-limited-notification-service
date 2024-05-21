package enum

type NotificationType int

const (
	Status NotificationType = iota
	News
	Marketing
	Unknown
)

func (n NotificationType) String() string {
	return [...]string{"status", "news", "marketing", "unknown"}[n]
}
