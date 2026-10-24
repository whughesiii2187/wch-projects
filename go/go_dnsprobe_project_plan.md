# Go Project: `dnsprobe` — a concurrent DNS lookup CLI

**Why this project:** you've got the Go fundamentals (structs, DB queries, forms via `huh`) but haven't needed goroutines, channels, `context`, or the `net` package yet. This project builds all four, using a domain (DNS/DNSSEC) you already know cold from the ops side — so you're learning the language, not the subject matter, one thing at a time.

Six milestones, roughly in order of difficulty. Each one has a goal, a skeleton (signature only — you write the body), a checkpoint to verify against, and a hint if you get stuck. No full solutions here on purpose.

---

## M1 — Single lookup, single record type

**Goal:** given a domain name as a command-line argument, print its IP addresses.

**New concepts:** `os.Args` (or the `flag` package if you want named flags), the `net` package's lookup functions, Go's `(result, error)` return idiom applied to network calls instead of DB calls.

**Skeleton:**
```go
func lookupHost(domain string) ([]string, error) {
    // TODO
}

func main() {
    // read domain from os.Args[1]
    // call lookupHost, handle the error, print results
}
```

**Checkpoint:** `go run . google.com` should print a short list of IP addresses. `go run . not-a-real-domain-xyz123.com` should print a clean error message, not a panic or a stack trace.

**Hint if stuck:** look at `net.LookupHost` in the standard library docs. It returns exactly the shape you need — this milestone is mostly about wiring, not logic.

---

## M2 — Multiple record types

**Goal:** extend M1 to also look up MX, TXT, and NS records, and print a small report with a section per record type.

**New concepts:** a few more `net` package functions with slightly different return types (`[]*net.MX` instead of `[]string`), and a struct to hold a multi-part result cleanly instead of printing as you go.

**Skeleton:**
```go
type LookupResult struct {
    Domain string
    A      []string
    MX     []*net.MX
    TXT    []string
    NS     []*net.NS
}

func lookupAll(domain string) (*LookupResult, error) {
    // TODO — call the four net.Lookup* functions, populate the struct
    // decide: if one lookup fails (e.g. no MX records), should the whole thing fail?
}
```

**Checkpoint:** run against a domain you know has MX records (your own company's domain, or a big one like `microsoft.com`) and confirm all four sections populate. Run against a domain with no TXT records and confirm it doesn't crash — just prints an empty section.

**Design question to actually think about, not skip:** what should `lookupAll` do if the A record lookup succeeds but the MX lookup fails? Return the partial result with an error noted, or fail the whole thing? There's no single right answer — but you should make the choice deliberately, not by accident.

---

## M3 — Timeouts with `context`

**Goal:** make every lookup respect a timeout, so a slow or unresponsive DNS server can't hang your program forever.

**New concepts:** `context.WithTimeout`, and `net.Resolver` (a struct version of the package-level lookup functions that accepts a context). This is one of the most important idioms in Go — you'll use this exact pattern in almost any networked code you ever write, including the Palo Alto client later.

**Skeleton:**
```go
func lookupHostWithTimeout(domain string, timeout time.Duration) ([]string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    resolver := net.Resolver{}
    // TODO — use resolver.LookupHost(ctx, domain) or resolver.LookupIPAddr(ctx, domain)
}
```

**Checkpoint:** hard-code an artificially tiny timeout (like 1 nanosecond) temporarily and confirm you get a clean timeout error rather than a hang. Then set it back to something reasonable (2–5 seconds) and confirm normal lookups still work.

**Why this matters for your resume goal specifically:** "handles timeouts gracefully" is exactly the kind of thing that separates a tutorial project from something you can talk about in an interview. Don't skip this milestone even though it's tempting to.

---

## M4 — Concurrency: run all lookups in parallel

**Goal:** take M2's sequential lookups (A, then MX, then TXT, then NS, one after another) and run them all at the same time, then print how much faster it is.

**New concepts:** goroutines (`go someFunc()`), `sync.WaitGroup` to wait for them all to finish, and — importantly — **how to get results back safely**. Writing to a shared struct from multiple goroutines without protection is a real bug (a data race), not a style nitpick.

**Two ways to do this correctly — pick one and understand why it's safe:**

- **Channels:** each goroutine sends its result on a channel; main reads from the channel. No shared memory, no race possible.
- **WaitGroup + mutex:** each goroutine writes to a shared struct, but only while holding a `sync.Mutex`.

**Skeleton (channel version):**
```go
type lookupJob struct {
    recordType string
    result     interface{}
    err        error
}

func lookupAllConcurrent(domain string) (*LookupResult, error) {
    results := make(chan lookupJob, 4) // buffered so goroutines don't block on send
    var wg sync.WaitGroup

    wg.Add(4)
    go func() {
        defer wg.Done()
        // TODO — do the A lookup, send a lookupJob on results
    }()
    // TODO — same pattern for MX, TXT, NS

    go func() {
        wg.Wait()
        close(results)
    }()

    // TODO — range over results, assemble the final LookupResult
}
```

**Checkpoint:** time M2's sequential version and M4's concurrent version against the same domain (`time.Now()` before/after, print the difference). The concurrent version should be noticeably faster — if it isn't, something's wrong with your goroutine wiring, not with Go.

**Run with the race detector before you call this done:**
```bash
go run -race . yourdomain.com
```
If it reports a data race, you have a real bug — don't ignore it because the output "looked fine." This is the single most valuable Go-specific habit you can build in this whole project.

---

## M5 (stretch) — Real DNSSEC validation check

**Goal:** report whether a domain actually has DNSSEC signatures present — not just "does it resolve," but "is it signed."

**New concepts:** using a third-party library (`github.com/miekg/dns`), which works at a lower level than the standard library's `net` package — you construct and send raw DNS messages yourself.

This one I won't skeleton out in detail — it's a stretch goal specifically because figuring out the `miekg/dns` API from its docs and examples is itself the exercise. Rough shape: build a DNS query message asking for `RRSIG` records on the domain, send it to a resolver (e.g., `8.8.8.8:53`), and check whether any signatures came back.

**Why this one's worth doing if you have the time:** "wrote a Go tool that checks DNSSEC signature presence" is a specific, verifiable, resume-worthy claim — directly tied to the DNSSEC work you already do at Relativity, just demonstrated in code instead of in ops.

---

## M6 (stretch) — Multiple domains from a file

**Goal:** read a list of domains (one per line, from a file or stdin) and run the full concurrent lookup pipeline across all of them — not just one domain concurrently, but many domains concurrently.

**New concept:** the **worker pool pattern** — a fixed number of goroutines pulling work off a channel, rather than spawning one goroutine per domain unbounded (which breaks badly if someone hands you a file with 10,000 domains).

This is the pattern most real Go networking tools use at scale, and it's the same shape you'd reuse for the Palo Alto client if you ever wanted it to check multiple firewalls instead of one.

---

## After this: the Palo Alto client (preview only)

Once M1–M4 are solid, the Palo Alto API client is a short jump, not a new project from scratch. Rough shape, so you can decide if it's worth doing:

1. **Get an API key manually first** (one-time, via `curl` against your dev firewall's keygen endpoint using your own credentials) — the Go program never handles your password, just the resulting key.
2. **Make one read-only HTTP GET** to an operational command endpoint (`show system info` is the standard "hello world" of the PAN-OS API) using `net/http` and your M3 timeout pattern.
3. **Parse the XML response** into a Go struct using `encoding/xml` — same idea as parsing DNS results into a struct in M2, different package.
4. Print a small report.

Purely read-only, no config changes, safe against a non-prod firewall — and it reuses the context/timeout/error-handling muscle memory from this project directly. Happy to build this one out the same way once you're through the DNS tool.
