package p1797designauthmanager

type token struct {
	id        string
	expiry    int
	expiryIdx int
}

type AuthenticationManager struct {
	ttl    int
	expiry []*token
	start  int
	tokens map[string]*token
}

func Constructor(timeToLive int) AuthenticationManager {
	m := AuthenticationManager{
		ttl:    timeToLive,
		expiry: make([]*token, 0, 2001),
		tokens: make(map[string]*token),
	}
	return m
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	n := len(this.expiry)
	token := &token{
		id:        tokenId,
		expiry:    currentTime + this.ttl,
		expiryIdx: n,
	}
	this.tokens[tokenId] = token
	this.expiry = append(this.expiry, token)
}

func (this *AuthenticationManager) cleanExpired(currentTime int) {
	newStart := this.start
	for i, token := range this.expiry[this.start:] {
		if token != nil && token.expiry > currentTime {
			break
		}
		if token != nil && token.expiry <= currentTime {
			delete(this.tokens, token.id)
			this.expiry[this.start+i] = nil
		}
		newStart++
	}
	this.start = newStart
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	this.cleanExpired(currentTime)
	t, exists := this.tokens[tokenId]
	if !exists {
		return
	}
	this.expiry[t.expiryIdx] = nil
	t.expiry = currentTime + this.ttl
	t.expiryIdx = len(this.expiry)
	this.expiry = append(this.expiry, t)
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	this.cleanExpired(currentTime)
	return len(this.tokens)
}
