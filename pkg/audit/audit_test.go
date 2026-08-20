package audit

import "testing"

func TestAuditListMetadataSnapshot(t *testing.T) {
	c := New()
	c.Append("a", "act", "s", map[string]string{"token": "secret"})
	out := c.List()
	out[0].Metadata["token"] = "changed"
	if c.List()[0].Metadata["token"] != "secret" {
		t.Fatal("audit metadata escaped")
	}
}
