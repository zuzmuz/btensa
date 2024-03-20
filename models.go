package main


type Secret struct {
    Title string `json:"title"`
    Value string `json:"value"`
    Ishidden bool `json:"isHidder"`
}

type Capsule struct {
    Name string `json:"name"`
    Secrets []Secret `json:"secrets"`
}

type Vault struct {
    Capsules []Capsule `json:"capsule"`
}
