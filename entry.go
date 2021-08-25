package slog

// Entry is a type to set log output information's
type Entry map[string]interface{}

// EntrySimple is a helper to create an entry with value only
// This has only effect on output format "text"
func EntrySimple(message string) Entry {
	return Entry{"": message}
}

// EntryMessage is a helper to create an entry with message field
func EntryMessage(message string) Entry {
	return Entry{"message": message}
}

// EntryEvent is a helper to create an entry with event field
func EntryEvent(event string) Entry {
	return Entry{"event": event}
}

// EntryError is a helper to create an entry with error field
func EntryError(err error) Entry {
	return Entry{"error": err.Error()}
}
